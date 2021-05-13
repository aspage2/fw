package fw

// Convert will convert certain ascii runes into their fullwidth version.
//  * Range U+FF01-U+FF5E maps directly to ascii 0x21-0x7E ('!' - '~').
//  * Ascii <space> is encoded by U+3000, a full-width whitespace character.
func Convert(r rune) rune {
	if r == ' ' {
		return rune(0x3000)
	}
	if r >= 0x21 && r <= 0x7E {
		return 0xFF01 + (r - 0x21)
	}
	return r
}
