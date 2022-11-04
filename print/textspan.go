package print

type TextSpan struct {
	// the file this span is in
	File string

	// the starting and ending indexes of this text span
	StartIndex int
	EndIndex   int

	// starting position
	StartLine   int
	StartColumn int

	// ending position
	EndLine   int
	EndColumn int
}

func (span1 TextSpan) SpanBetween(span2 TextSpan) TextSpan {
	// null checks
	if span1.File == "" && span2.File == "" {
		return TextSpan{}
	} else if span1.File == "" {
		return span2
	} else if span2.File == "" {
		return span1
	}

	// we good
	var first TextSpan
	var last TextSpan

	if span1.StartIndex < span2.StartIndex {
		first = span1
	} else {
		first = span2
	}

	if span1.EndIndex > span2.EndIndex {
		last = span1
	} else {
		last = span2
	}

	var union TextSpan
	union.File = span1.File

	union.StartIndex = first.StartIndex
	union.EndIndex = last.EndIndex

	union.StartLine = first.StartLine
	union.EndLine = last.EndLine

	union.StartColumn = first.StartColumn
	union.EndColumn = last.EndColumn

	return union
}
