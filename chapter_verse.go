package brave

import (
	"bytes"
	"fmt"
	p "github.com/saintfish/parser.go"
	"strconv"
)

type Range struct {
	Begin, End byte
}

type Number struct {
	Number byte
	Ranges []Range
}

func (this Number) String() string {
	if len(this.Ranges) == 0 {
		return fmt.Sprintf("%d", this.Number)
	}
	b := &bytes.Buffer{}
	for i, r := range this.Ranges {
		if i != 0 {
			b.WriteString(",")
		}
		if r.Begin == r.End {
			b.WriteString(fmt.Sprintf("%d", r.Begin))
		} else {
			b.WriteString(fmt.Sprintf("%d-%d", r.Begin, r.End))
		}
	}
	return b.String()
}

type Annotation struct {
	Begin, End int
}

type ChapterVerse struct {
	Chapter    Number
	Verse      *Number
	Annotation Annotation
}

var (
	spaces = p.Regexp("[ \\t\u3000]*")
	number = p.HandleRegexp(
		func(b *p.Buffer, r p.Run) p.Value {
			n, err := strconv.ParseUint(b.Run(r), 10, 8)
			if err != nil {
				return byte(0)
			}
			return byte(n)
		},
		"\\d+")
	range_ = p.HandleAlter(
		func(b *p.Buffer, r p.Run, v p.Value, index int) p.Value {
			if index == 0 {
				values := v.([]p.Value)
				return Range{Begin: values[0].(byte), End: values[4].(byte)}
			}
			return Range{Begin: v.(byte), End: v.(byte)}
		},
		p.Cat(number, spaces, p.Regexp("[-—~～]"), spaces, number),
		number)
	rangeList = p.HandleCat(
		func(b *p.Buffer, r p.Run, values []p.Value) p.Value {
			ranges := []Range{}
			ranges = append(ranges, values[0].(Range))
			for _, v := range values[1].([]p.Value) {
				ranges = append(ranges, v.([]p.Value)[3].(Range))
			}
			if len(ranges) == 1 && ranges[0].Begin == ranges[0].End {
				return Number{Number: ranges[0].Begin}
			}
			return Number{Ranges: ranges}
		},
		range_, p.Repeat(p.Cat(spaces, p.Rune(",，"), spaces, range_)))
	singleNumber = p.HandleCat(
		func(b *p.Buffer, r p.Run, values []p.Value) p.Value {
			return Number{Number: values[0].(byte)}
		},
		number)
	chapterVerse = p.HandleAlter(
		func(b *p.Buffer, r p.Run, value p.Value, index int) p.Value {
			annotation := Annotation{
				Begin: r.Start,
				End:   r.End,
			}
			if index == 0 {
				chapter := value.([]p.Value)[2].(Number)
				verse := value.([]p.Value)[8].(Number)
				return ChapterVerse{
					Chapter:    chapter,
					Verse:      &verse,
					Annotation: annotation,
				}
			} else if index == 1 {
				return ChapterVerse{
					Chapter:    value.([]p.Value)[2].(Number),
					Annotation: annotation,
				}
			}
			chapter := value.([]p.Value)[0].(Number)
			verse := value.([]p.Value)[4].(Number)
			return ChapterVerse{
				Chapter:    chapter,
				Verse:      &verse,
				Annotation: annotation,
			}
		},
		p.Cat(
			p.Option(p.Literal("第")), spaces, singleNumber, spaces, p.Rune("章篇"), spaces,
			p.Option(p.Literal("第")), spaces, rangeList, spaces, p.Rune("节節")),
		p.Cat(p.Option(p.Literal("第")), spaces, rangeList, spaces, p.Rune("章篇")),
		p.Cat(singleNumber, spaces, p.Rune(":："), spaces, rangeList))
	spaceChapterVerse = p.HandleCat(
		func(b *p.Buffer, r p.Run, values []p.Value) p.Value {
			return values[1]
		},
		spaces, chapterVerse, spaces)
)

func parseChapterVerse(input string, offset int) *ChapterVerse {
	v, err := p.ParseString(input, offset, spaceChapterVerse)
	if err != nil {
		return nil
	}
	cv := v.(ChapterVerse)
	return &cv
}
