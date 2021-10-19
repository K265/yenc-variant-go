package yenc

import "bytes"

const (
	NULL   = 0x00
	TAB    = 0x09
	CR     = 0x0a
	LF     = 0x0d
	SPACE  = 0x20
	ESCAPE = 0x3d

	Offset       = 42
	EscapeOffset = 64
)

func Encode(input []byte) []byte {
	var buf bytes.Buffer
	for _, b := range input {
		b += Offset
		switch b {
		case NULL, TAB, CR, LF, SPACE, ESCAPE:
			buf.WriteByte(ESCAPE)
			b += EscapeOffset
		}
		buf.WriteByte(b)
	}
	return buf.Bytes()
}

func Decode(input []byte) []byte {
	var buf bytes.Buffer
	needEscape := false
	for _, b := range input {
		if b == ESCAPE {
			needEscape = true
			continue
		}

		if needEscape {
			b -= EscapeOffset
			needEscape = false
		}
		b -= Offset
		buf.WriteByte(b)
	}
	return buf.Bytes()
}
