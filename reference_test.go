package brave

import (
	"fmt"
	"testing"
)

func TestFindAllReferences(t *testing.T) {
	input := `
	彼得前書5章5-9節
	讀經: 詩篇116篇
2014年 01月 16日 - David C. McCasland

MP3 下載
《灵命日粮》网络广播是由李芳主持
读经: 歌罗西书1章1-12节， 4章12节

因为父喜欢叫一切的丰盛在祂里面居住。 —歌罗西书1章19节 

全年读经: 创世记39-40章 马太福音11章 
`
	fmt.Println(FindAllReferences(input))
}
