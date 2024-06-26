package json

import (
	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/fs"
	"github.com/jsightapi/jsight-schema-core/internal/ds"
	"github.com/jsightapi/jsight-schema-core/kit"
	"github.com/jsightapi/jsight-schema-core/lexeme"
)

type (
	stepFunc func(*scanner, byte) state
)

// state values are returned by the state transition functions assigned to
// scanner.state and the method scanner.eof.
// They give details about the current state of the scan that callers might be
// interested to know about.
// It is okay to ignore the return value of any particular call to scanner.state.
type state uint8

const (
	// scanContinue indicates an uninteresting byte, so we can keep scanning forward.
	scanContinue state = iota // uninteresting byte

	// scanBeginObject indicates beginning of an object.
	scanBeginObject

	// scanBeginArray indicates beginning of an array.
	scanBeginArray

	// scanBeginLiteral indicates beginning of any value outside an array or object.
	scanBeginLiteral
)

// scanner represents a scanner is a JSON scanning state machine.
// Callers call scan.reset() and then pass bytes in one at a time by calling
// scan.step(&scan, c) for each byte.
// The return value, referred to as an opcode, tells the caller about significant
// parsing events like beginning and ending literals, objects, and arrays, so that
// the caller can follow along if it wishes.
// The return value scanEnd indicates that a single top-level JSON value has been
// completed, *before* the byte that just got passed in.  (The indication must be
// delayed in order to recognize the end of numbers: is 123 a whole value or the
// beginning of 12345e+6?).
type scanner struct {
	// step is a func to be called to execute the next transition.
	// Also tried using an integer constant and a single func with a switch, but
	// using the func directly was 10% faster on a 64-bit Mac Mini, and it's nicer
	// to read.
	step stepFunc

	// returnToStep a stack of step functions, to preserve the sequence of steps
	// (and return to them) in some cases.
	returnToStep *ds.Stack[stepFunc]

	// file a structure containing JSON data.
	file *fs.File

	// data a JSON content.
	data bytes.Bytes

	// stack a stack of found lexical event. The stack is needed for the scanner
	// to take into account the nesting of JSON or SCHEME elements.
	stack *ds.Stack[lexeme.LexEvent]

	// finds a list of found types of lexical event for the current step. Several
	// lexical events can be found in one step (example: ArrayItemBegin and LiteralBegin).
	finds []lexeme.LexEventType

	// index scanned byte index.
	index bytes.Index

	// dataSize a size of JSON data in bytes. Count once for optimization.
	dataSize bytes.Index

	// unfinishedLiteral a sign that a literal has been started but not completed.
	unfinishedLiteral bool

	// allowTrailingNonSpaceCharacters allows to have non-empty characters at the
	// end of the JSON.
	allowTrailingNonSpaceCharacters bool
}

func newScanner(file *fs.File) *scanner {
	return &scanner{
		step:         stateFoundRootValue,
		file:         file,
		data:         file.Content(),
		dataSize:     file.Content().LenIndex(),
		returnToStep: &ds.Stack[stepFunc]{},
		stack:        &ds.Stack[lexeme.LexEvent]{},
		finds:        make([]lexeme.LexEventType, 0, 3),
	}
}

func (s *scanner) Length() uint {
	var length uint
	for {
		lex, ok := s.Next()
		if !ok {
			break
		}

		if lex.Type() == lexeme.EndTop {
			// Found character after the end of the schema and spaces. Ex: char
			// "s" in "{} some text".
			length = uint(lex.End()) - 1
			break
		}
		length = uint(lex.End()) + 1
	}
	for {
		if length == 0 {
			break
		}
		c := s.data.Byte(length - 1)
		if bytes.IsBlank(c) {
			length--
		} else {
			break
		}
	}
	return length
}

// Next reads JSON byte by byte.
// Panic if an invalid JSON structure is found.
// Stops if it detects lexical events.
// Returns pointer to found lexeme event, or nil if you have complete JSON reading.
func (s *scanner) Next() (lexeme.LexEvent, bool) {
	if len(s.finds) != 0 {
		return s.processingFoundLexeme(s.shiftFound()), true
	}

	for s.index < s.dataSize {
		c := s.data.Byte(s.index)
		s.index++

		s.step(s, c)

		if len(s.finds) != 0 {
			return s.processingFoundLexeme(s.shiftFound()), true
		}
	}

	if s.stack.Len() != 0 {
		s.index++
		switch s.stack.Peek().Type() { //nolint:exhaustive // We handle all cases.
		case lexeme.LiteralBegin:
			if s.unfinishedLiteral {
				break
			}
			return s.processingFoundLexeme(lexeme.LiteralEnd), true
		case lexeme.InlineAnnotationBegin:
			return s.processingFoundLexeme(lexeme.InlineAnnotationEnd), true
		case lexeme.InlineAnnotationTextBegin:
			return s.processingFoundLexeme(lexeme.InlineAnnotationTextEnd), true
		}
		err := kit.NewJSchemaError(s.file, errs.ErrUnexpectedEOF.F())
		err.SetIndex(s.dataSize - 1)
		panic(err)
	}

	return lexeme.LexEvent{}, false
}

func (s *scanner) found(lexType lexeme.LexEventType) {
	s.finds = append(s.finds, lexType)
}

func (s *scanner) shiftFound() lexeme.LexEventType {
	length := len(s.finds)
	if length == 0 {
		panic(errs.ErrEmptySetOfLexicalEvents.F())
	}
	lexType := s.finds[0]
	copy(s.finds[0:], s.finds[1:])
	s.finds = s.finds[:length-1]
	return lexType
}

func (s *scanner) processingFoundLexeme(lexType lexeme.LexEventType) lexeme.LexEvent {
	i := s.index - 1
	if lexType == lexeme.NewLine || lexType == lexeme.EndTop {
		return lexeme.NewLexEvent(lexType, i, i, s.file)
	}

	if lexType.IsOpening() {
		var lex lexeme.LexEvent
		if lexType == lexeme.InlineAnnotationBegin || lexType == lexeme.MultiLineAnnotationBegin {
			lex = lexeme.NewLexEvent(lexType, i-1, i, s.file) // `//` or `/*`
		} else {
			// `{`, `[`, `"` or literal first character (ex: `1` in `123`).
			lex = lexeme.NewLexEvent(lexType, i, i, s.file)
		}
		s.stack.Push(lex)
		return lex
	}

	return s.processFoundLexemeClosingTag(lexType, i)
}

func (s *scanner) processFoundLexemeClosingTag(lexType lexeme.LexEventType, i bytes.Index) lexeme.LexEvent {
	pair := s.stack.Pop()
	pairType := pair.Type()
	if isNonScalarPair(pairType, lexType) {
		return lexeme.NewLexEvent(lexType, pair.Begin(), i, s.file)
	}

	if isScalarPair(pairType, lexType) {
		return lexeme.NewLexEvent(lexType, pair.Begin(), i-1, s.file)
	}
	panic(errs.ErrIncorrectEndingOfTheLexicalEvent.F())
}

func isNonScalarPair(pairType, lexType lexeme.LexEventType) bool {
	return (pairType == lexeme.ObjectBegin && lexType == lexeme.ObjectEnd) ||
		(pairType == lexeme.ArrayBegin && lexType == lexeme.ArrayEnd)
}

func isScalarPair(pairType, lexType lexeme.LexEventType) bool {
	return (pairType == lexeme.LiteralBegin && lexType == lexeme.LiteralEnd) ||
		(pairType == lexeme.ArrayItemBegin && lexType == lexeme.ArrayItemEnd) ||
		(pairType == lexeme.ObjectKeyBegin && lexType == lexeme.ObjectKeyEnd) ||
		(pairType == lexeme.ObjectValueBegin && lexType == lexeme.ObjectValueEnd)
}

func stateFoundRootValue(s *scanner, c byte) state {
	r := stateBeginValue(s, c)
	switch r { //nolint:exhaustive // It's okay.
	case scanBeginObject:
		s.found(lexeme.ObjectBegin)

	case scanBeginArray:
		s.found(lexeme.ArrayBegin)

	case scanBeginLiteral:
		s.found(lexeme.LiteralBegin)
	}
	return r
}

func stateFoundObjectKeyBeginOrEmpty(s *scanner, c byte) state {
	if bytes.IsBlank(c) {
		return scanContinue
	}

	return stateBeginKeyOrEmpty(s, c)
}

func stateFoundObjectKeyBegin(s *scanner, c byte) state {
	if bytes.IsBlank(c) {
		return scanContinue
	}

	r := stateBeginString(s, c)
	s.found(lexeme.ObjectKeyBegin)
	return r
}

func stateFoundObjectValueBegin(s *scanner, c byte) state {
	r := stateBeginValue(s, c)
	switch r { //nolint:exhaustive // It's okay.
	case scanBeginLiteral:
		s.found(lexeme.ObjectValueBegin)
		s.found(lexeme.LiteralBegin)

	case scanBeginObject:
		s.found(lexeme.ObjectValueBegin)
		s.found(lexeme.ObjectBegin)

	case scanBeginArray:
		s.found(lexeme.ObjectValueBegin)
		s.found(lexeme.ArrayBegin)
	}
	return r
}

func stateFoundArrayItemBeginOrEmpty(s *scanner, c byte) state {
	r := stateBeginArrayItemOrEmpty(s, c)
	switch r { //nolint:exhaustive // It's okay.
	case scanBeginLiteral:
		s.found(lexeme.ArrayItemBegin)
		s.found(lexeme.LiteralBegin)

	case scanBeginObject:
		s.found(lexeme.ArrayItemBegin)
		s.found(lexeme.ObjectBegin)

	case scanBeginArray:
		s.found(lexeme.ArrayItemBegin)
		s.found(lexeme.ArrayBegin)
	}
	return r
}

func stateFoundArrayItemBegin(s *scanner, c byte) state {
	r := stateBeginValue(s, c)
	switch r { //nolint:exhaustive // It's okay.
	case scanBeginLiteral:
		s.found(lexeme.ArrayItemBegin)
		s.found(lexeme.LiteralBegin)

	case scanBeginObject:
		s.found(lexeme.ArrayItemBegin)
		s.found(lexeme.ObjectBegin)

	case scanBeginArray:
		s.found(lexeme.ArrayItemBegin)
		s.found(lexeme.ArrayBegin)
	}
	return r
}

func stateBeginValue(s *scanner, c byte) state {
	if bytes.IsBlank(c) {
		return scanContinue
	}
	switch c {
	case '{':
		s.step = stateFoundObjectKeyBeginOrEmpty
		return scanBeginObject
	case '[':
		s.step = stateFoundArrayItemBeginOrEmpty
		return scanBeginArray
	case '"':
		s.step = stateInString
		s.unfinishedLiteral = true
		return scanBeginLiteral
	case '-':
		s.step = stateNeg
		s.unfinishedLiteral = true
		return scanBeginLiteral
	case '0': // beginning of 0.123
		s.step = state0
		return scanBeginLiteral
	case 't': // beginning of true
		s.step = stateT
		s.unfinishedLiteral = true
		return scanBeginLiteral
	case 'f': // beginning of false
		s.step = stateF
		s.unfinishedLiteral = true
		return scanBeginLiteral
	case 'n': // beginning of null
		s.step = stateN
		s.unfinishedLiteral = true
		return scanBeginLiteral
	}
	if '1' <= c && c <= '9' { // beginning of 1234.5
		s.step = state1
		return scanBeginLiteral
	}
	panic(s.newJSchemaErrorAtCharacter("— JSON value expected (number, string, boolean, object, array, or null)"))
}

// After reading `[`.
func stateBeginArrayItemOrEmpty(s *scanner, c byte) state {
	if c == ']' {
		return stateFoundArrayEnd(s)
	}
	return stateBeginValue(s, c)
}

// After reading `{`.
func stateBeginKeyOrEmpty(s *scanner, c byte) state {
	if c == '}' {
		return stateFoundObjectEnd(s)
	}
	s.found(lexeme.ObjectKeyBegin)
	return stateBeginString(s, c)
}

// After reading `{"key": value,`.
func stateBeginString(s *scanner, c byte) state {
	if c == '"' {
		s.step = stateInString
		return scanBeginLiteral
	}
	panic(s.newJSchemaErrorAtCharacter("— string literal expected (starting with the quotation mark `\"`)"))
}

func stateEndValue(s *scanner, c byte) state {
	length := s.stack.Len()

	if length == 0 { // json ex `{} `
		s.step = stateEndTop
		return s.step(s, c)
	}

	t := s.stack.Peek().Type()

	if t == lexeme.LiteralBegin {
		s.found(lexeme.LiteralEnd)

		if length == 1 { // json ex `123 `
			s.step = stateEndTop
			return s.step(s, c)
		}

		t = s.stack.Get(length - 2).Type()
	}

	switch t { //nolint:exhaustive // We will throw a panic in over cases.
	case lexeme.ObjectKeyBegin:
		s.found(lexeme.ObjectKeyEnd)
		s.step = stateAfterObjectKey
		return s.step(s, c)
	case lexeme.ObjectValueBegin:
		s.found(lexeme.ObjectValueEnd)
		s.step = stateAfterObjectValue
		return s.step(s, c)
	case lexeme.ArrayItemBegin:
		s.found(lexeme.ArrayItemEnd)
		s.step = stateAfterArrayItem
		return s.step(s, c)
	}
	panic(s.newJSchemaErrorAtCharacter("at the end of value"))
}

func stateAfterObjectKey(s *scanner, c byte) state {
	if bytes.IsBlank(c) {
		return scanContinue
	}

	if c == ':' {
		s.step = stateFoundObjectValueBegin
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("after object key"))
}

func stateAfterObjectValue(s *scanner, c byte) state {
	if bytes.IsBlank(c) {
		return scanContinue
	}
	if c == ',' {
		s.step = stateFoundObjectKeyBegin
		return scanContinue
	}
	if c == '}' {
		return stateFoundObjectEnd(s)
	}
	panic(s.newJSchemaErrorAtCharacter("after the object property, should be \",\" or \"}\""))
}

func stateAfterArrayItem(s *scanner, c byte) state {
	if bytes.IsBlank(c) {
		return scanContinue
	}
	if c == ',' {
		s.step = stateFoundArrayItemBegin
		return scanContinue
	}
	if c == ']' {
		return stateFoundArrayEnd(s)
	}
	panic(s.newJSchemaErrorAtCharacter("after array item"))
}

func stateFoundObjectEnd(s *scanner) state {
	s.found(lexeme.ObjectEnd)
	s.step = stateEndValue
	return scanContinue
}

func stateFoundArrayEnd(s *scanner) state {
	s.found(lexeme.ArrayEnd)
	if s.stack.Len() == 0 {
		s.step = stateEndTop
	} else {
		s.step = stateEndValue
	}
	return scanContinue
}

// stateEndTop is the state after finishing the top-level value, such as after
// reading `{}` or `[1,2,3]`.
// Only space characters should be seen now.
func stateEndTop(s *scanner, c byte) state {
	if !bytes.IsBlank(c) {
		if !s.allowTrailingNonSpaceCharacters {
			panic(s.newJSchemaErrorAtCharacter("non-space byte after top-level value"))
		}
		s.found(lexeme.EndTop)
	}
	return scanContinue
}

// After reading `"`.
func stateInString(s *scanner, c byte) state {
	switch c {
	case '"':
		s.step = stateEndValue
		s.unfinishedLiteral = false
		return scanContinue
	case '\\':
		s.step = stateInStringEsc
		return scanContinue
	}
	if c < 0x20 {
		panic(s.newJSchemaErrorAtCharacter("in string literal"))
	}
	return scanContinue
}

// After reading `"\` during a quoted string.
func stateInStringEsc(s *scanner, c byte) state {
	switch c {
	case 'b', 'f', 'n', 'r', 't', '\\', '/', '"':
		s.step = stateInString
		return scanContinue
	case 'u':
		s.returnToStep.Push(stateInString)
		s.step = stateInStringEscU
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in string escape code"))
}

// After reading `"\u` during a quoted string.
func stateInStringEscU(s *scanner, c byte) state {
	if bytes.IsHexDigit(c) {
		s.step = stateInStringEscU1
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in \\u hexadecimal character escape"))
}

// After reading `"\u1` during a quoted string.
func stateInStringEscU1(s *scanner, c byte) state {
	if bytes.IsHexDigit(c) {
		s.step = stateInStringEscU12
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in \\u hexadecimal character escape"))
}

// After reading `"\u12` during a quoted string.
func stateInStringEscU12(s *scanner, c byte) state {
	if bytes.IsHexDigit(c) {
		s.step = stateInStringEscU123
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in \\u hexadecimal character escape"))
}

// After reading `"\u123` during a quoted string.
func stateInStringEscU123(s *scanner, c byte) state {
	if bytes.IsHexDigit(c) {
		s.step = s.returnToStep.Pop() // = stateInString for JSON, = stateInAnnotationObjectKey for AnnotationObject
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in \\u hexadecimal character escape"))
}

// After reading `-` during a number.
func stateNeg(s *scanner, c byte) state {
	if c == '0' {
		s.step = state0
		s.unfinishedLiteral = false
		return scanContinue
	}
	if '1' <= c && c <= '9' {
		s.step = state1
		s.unfinishedLiteral = false
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in numeric literal"))
}

// After reading a non-zero integer during a number, such as after reading `1` or
// `100` but not `0`.
func state1(s *scanner, c byte) state {
	if bytes.IsDigit(c) {
		s.step = state1
		return scanContinue
	}
	return state0(s, c)
}

// After reading `0` during a number.
func state0(s *scanner, c byte) state {
	if c == '.' {
		s.step = stateDot
		return scanContinue
	}
	if c == 'e' || c == 'E' {
		s.step = stateE
		return scanContinue
	}
	return stateEndValue(s, c)
}

// After reading the integer and decimal point in a number, such as after reading
// `1.`.
func stateDot(s *scanner, c byte) state {
	if bytes.IsDigit(c) {
		s.step = stateDot0
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("after decimal point in numeric literal"))
}

// After reading the integer, decimal point, and subsequent digits of a number,
// such as after reading `3.14`.
func stateDot0(s *scanner, c byte) state {
	if bytes.IsDigit(c) {
		return scanContinue
	}
	if c == 'e' || c == 'E' {
		s.step = stateE
		return scanContinue
	}
	return stateEndValue(s, c)
}

// After reading the mantissa and e in a number, such as after reading `314e` or
// `0.314e`.
func stateE(s *scanner, c byte) state {
	if c == '+' || c == '-' {
		s.step = stateESign
		return scanContinue
	}
	return stateESign(s, c)
}

// After reading the mantissa, e, and sign in a number, such as after reading
// `314e-` or `0.314e+`.
func stateESign(s *scanner, c byte) state {
	if bytes.IsDigit(c) {
		s.step = stateE0
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in exponent of numeric literal"))
}

// After reading the mantissa, e, optional sign, and at least one digit of the
// exponent in a number, such as after reading `314e-2` or `0.314e+1` or `3.14e0`.
func stateE0(s *scanner, c byte) state {
	if bytes.IsDigit(c) {
		return scanContinue
	}
	return stateEndValue(s, c)
}

// After reading `t`.
func stateT(s *scanner, c byte) state {
	if c == 'r' {
		s.step = stateTr
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal true (expecting 'r')"))
}

// After reading `tr`.
func stateTr(s *scanner, c byte) state {
	if c == 'u' {
		s.step = stateTru
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal true (expecting 'u')"))
}

// After reading `tru`.
func stateTru(s *scanner, c byte) state {
	if c == 'e' {
		s.step = stateEndValue
		s.unfinishedLiteral = false
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal true (expecting 'e')"))
}

// After reading `f`.
func stateF(s *scanner, c byte) state {
	if c == 'a' {
		s.step = stateFa
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal false (expecting 'a')"))
}

// After reading `fa`.
func stateFa(s *scanner, c byte) state {
	if c == 'l' {
		s.step = stateFal
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal false (expecting 'l')"))
}

// After reading `fal`.
func stateFal(s *scanner, c byte) state {
	if c == 's' {
		s.step = stateFals
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal false (expecting 's')"))
}

// After reading `fals`.
func stateFals(s *scanner, c byte) state {
	if c == 'e' {
		s.step = stateEndValue
		s.unfinishedLiteral = false
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal false (expecting 'e')"))
}

// After reading `n`.
func stateN(s *scanner, c byte) state {
	if c == 'u' {
		s.step = stateNu
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal null (expecting 'u')"))
}

// After reading `nu`.
func stateNu(s *scanner, c byte) state {
	if c == 'l' {
		s.step = stateNul
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal null (expecting 'l')"))
}

// After reading `nul`.
func stateNul(s *scanner, c byte) state {
	if c == 'l' {
		s.step = stateEndValue
		s.unfinishedLiteral = false
		return scanContinue
	}
	panic(s.newJSchemaErrorAtCharacter("in literal null (expecting 'l')"))
}

func (s *scanner) newJSchemaErrorAtCharacter(context string) kit.JSchemaError {
	// Make runes (utf8 symbols) from current index to last of slice s.data.
	// Get first rune. Then make string with format ' symbol '
	r := s.data.SubLow(s.index - 1).DecodeRune()
	e := errs.ErrInvalidCharacter.F(string(r), context)
	err := kit.NewJSchemaError(s.file, e)
	err.SetIndex(s.index - 1)
	return err
}
