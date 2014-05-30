package brave

import (
	"unicode/utf8"
)

type Reference struct {
	Book         Book
	ChapterVerse ChapterVerse
	Annotation   Annotation
	Text         string
}

func parseReference(bp *bookParser, input string, pos int, prevBook *Book) *Reference {
	if prevBook != nil {
		cv := parseChapterVerse(input, pos)
		if cv != nil {
			return &Reference{
				Book:         *prevBook,
				ChapterVerse: *cv,
				Annotation: Annotation{
					Begin: pos,
					End:   cv.Annotation.End,
				},
				Text: input[pos:cv.Annotation.End],
			}
		}
	}
	b := bp.parseFrom(input, pos)
	if b != nil {
		cv := parseChapterVerse(input, b.Annotation.End)
		if cv != nil {
			return &Reference{
				Book:         *b,
				ChapterVerse: *cv,
				Annotation: Annotation{
					Begin: pos,
					End:   cv.Annotation.End,
				},
				Text: input[pos:cv.Annotation.End],
			}
		}
	}
	return nil
}

var chineseBookParser = newChineseBookParser()

func FindAllReferences(input string) []Reference {
	result := []Reference{}
	charSinceLastRef := 0
	prevBook := (*Book)(nil)
	for i := 0; i < len(input); {
		r := parseReference(chineseBookParser, input, i, prevBook)
		if r != nil {
			result = append(result, *r)
			i = r.Annotation.End
			charSinceLastRef = 0
			prevBook = &(r.Book)
		} else {
			if input[i] < utf8.RuneSelf {
				i++
			} else {
				_, w := utf8.DecodeRuneInString(input[i:])
				i += w
			}
			charSinceLastRef++
			if prevBook != nil && charSinceLastRef > 10 {
				prevBook = nil
			}
		}
	}
	return result
}
