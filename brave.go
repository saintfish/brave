package brave

import (
	"github.com/saintfish/bible.go/bible"
)

func ParseChinesePrefix(input string) (bool, int, bible.RefRangeList) {
	match, offset, value := chineseParser.parse(input)
	if match && value != nil {
		return match, offset, value.(bible.RefRangeList)
	}
	return match, offset, nil
}

func ParseChineseFull(input string) (bool, bible.RefRangeList) {
	match, offset, value := ParseChinesePrefix(input)
	match = match && offset == len(input)
	return match, value
}
