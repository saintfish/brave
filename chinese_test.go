package brave

import (
	"github.com/saintfish/bible.go/bible"
	"reflect"
	"testing"
)

func TestChineseParser(t *testing.T) {
	type testDataEntry struct {
		input string
		ref   bible.RefRangeList
	}
	var testData = []testDataEntry{
		testDataEntry{"希伯來書4章14-16節", bible.SingleRangeRef(bible.Hebrews, 4, 14, 16)},
		testDataEntry{"詩篇73篇", bible.SingleRangeRef(bible.Psalm, 73, bible.ChapterBegin, bible.ChapterEnd)},
		testDataEntry{"詩篇73篇", bible.SingleRangeRef(bible.Psalm, 73, bible.ChapterBegin, bible.ChapterEnd)},
	}
	for _, entry := range testData {
		match, offset, value := chineseParser.parse(entry.input)
		if !match || offset != len(entry.input) {
			t.Error(entry.input, match, offset)
		}
		if !reflect.DeepEqual(value, entry.ref) {
			t.Error(value, entry.ref)
		}
	}
}
