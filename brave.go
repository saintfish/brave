package brave

import (
	"github.com/saintfish/bible.go/bible"
)

func ParseChinesePrefix(input string) (bool, int, bible.RefRangeList) {
	match, offset, value := chineseParser.parse(input)
	return match, offset, value.(bible.RefRangeList)
}

func ParseChineseFull(input string) (bool, bible.RefRangeList) {
	match, offset, value := ParseChinesePrefix(input)
	match = match && offset == len(input)
	return match, value
}
