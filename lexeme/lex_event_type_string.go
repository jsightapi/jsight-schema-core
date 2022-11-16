// Code generated by "stringer -type LexEventType -linecomment -output /home/add/go/src/jsightapi/jsight-schema-core/internal/lexeme/lex_event_type_string.go"; DO NOT EDIT.

package lexeme

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LiteralBegin-0]
	_ = x[LiteralEnd-1]
	_ = x[ObjectBegin-2]
	_ = x[ObjectEnd-3]
	_ = x[ObjectKeyBegin-4]
	_ = x[ObjectKeyEnd-5]
	_ = x[ObjectValueBegin-6]
	_ = x[ObjectValueEnd-7]
	_ = x[ArrayBegin-8]
	_ = x[ArrayEnd-9]
	_ = x[ArrayItemBegin-10]
	_ = x[ArrayItemEnd-11]
	_ = x[InlineAnnotationBegin-12]
	_ = x[InlineAnnotationEnd-13]
	_ = x[InlineAnnotationTextBegin-14]
	_ = x[InlineAnnotationTextEnd-15]
	_ = x[MultiLineAnnotationBegin-16]
	_ = x[MultiLineAnnotationEnd-17]
	_ = x[MultiLineAnnotationTextBegin-18]
	_ = x[MultiLineAnnotationTextEnd-19]
	_ = x[NewLine-20]
	_ = x[TypesShortcutBegin-21]
	_ = x[TypesShortcutEnd-22]
	_ = x[KeyShortcutBegin-23]
	_ = x[KeyShortcutEnd-24]
	_ = x[MixedValueBegin-25]
	_ = x[MixedValueEnd-26]
	_ = x[EndTop-27]
}

const _LexEventType_name = "literal-beginliteral-endobject-beginobject-endkey-beginkey-endvalue-beginvalue-endarray-beginarray-enditem-beginitem-endinline-annotation-begininline-annotation-endinline-annotation-text-begininline-annotation-text-endmulti-line-annotation-beginmulti-line-annotation-endmulti-line-annotation-text-beginmulti-line-annotation-text-endnew-linetypes-shortcut-begintypes-shortcut-endkey-shortcut-beginkey-shortcut-endmixed-value-beginmixed-value-endend-top"

var _LexEventType_index = [...]uint16{0, 13, 24, 36, 46, 55, 62, 73, 82, 93, 102, 112, 120, 143, 164, 192, 218, 245, 270, 302, 332, 340, 360, 378, 396, 412, 429, 444, 451}

func (e LexEventType) String() string {
	if e >= LexEventType(len(_LexEventType_index)-1) {
		panic("Unknown lexical event type")
	}
	return _LexEventType_name[_LexEventType_index[e]:_LexEventType_index[e+1]]
}
