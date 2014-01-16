package brave

import (
	"fmt"
	"testing"
)

func TestParseChapterVerse(t *testing.T) {
	fmt.Println(parseChapterVerse("  第1章 12-15, 16, 19-26 节", 0))
}
