package schema

func quotedBytes(s string) []byte {
	bb := make([]byte, 0, len(s)+2)

	bb = append(bb, '"')
	bb = append(bb, []byte(s)...)
	bb = append(bb, '"')

	return bb
}
