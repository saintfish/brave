package brave

import (
	"fmt"
	"github.com/saintfish/bible.go/bible"
	"strconv"
)

var chinesePatterns = []struct {
	pattern string
	ref     bible.RefRangeList
}{
	{
		"希伯來書 12 章 1-5 節",
		bible.SingleRangeRef(bible.Hebrews, 12, 1, 5),
	},
	{
		"提摩太後書 3 章 1-5 , 10-17 節",
		bible.RefRangeList{
			bible.SingleRangeRef(bible.Timothy2, 3, 1, 5)[0],
			bible.SingleRangeRef(bible.Timothy2, 3, 10, 17)[0],
		},
	},
	{
		"羅馬書 8 章 14-17 節 , 24-26 節",
		bible.RefRangeList{
			bible.SingleRangeRef(bible.Romans, 8, 14, 17)[0],
			bible.SingleRangeRef(bible.Romans, 8, 24, 26)[0],
		},
	},
	{
		"尼希米記 6 章 1-9 , 15 節",
		bible.RefRangeList{
			bible.SingleRangeRef(bible.Nehemiah, 6, 1, 9)[0],
			bible.SingleRangeRef(bible.Nehemiah, 6, 15, 15)[0],
		},
	},
	{
		"诗篇 23 篇",
		bible.SingleRangeRef(bible.Psalm, 23, bible.ChapterBegin, bible.ChapterEnd),
	},
	{
		"撒母耳記上 18 章 1-4 , 23 章 15-19 節",
		bible.RefRangeList{
			bible.SingleRangeRef(bible.Samuel1, 18, 1, 4)[0],
			bible.SingleRangeRef(bible.Samuel1, 23, 15, 19)[0],
		},
	},
	{
		"民數記 13 章 25 節 - 14 章 19 節",
		bible.RefRangeList{
			bible.SingleRangeRef(bible.Numbers, 13, 25, 25)[0],
			bible.SingleRangeRef(bible.Numbers, 14, 19, 19)[0],
		},
	},
}

var chineseBookNames = map[string]bible.BookID{
	"创世记":     bible.Genesis,
	"創世記":     bible.Genesis,
	"出埃及記":    bible.Exodus,
	"出埃及记":    bible.Exodus,
	"利未記":     bible.Leviticus,
	"利未记":     bible.Leviticus,
	"民数记":     bible.Numbers,
	"民數記":     bible.Numbers,
	"申命記":     bible.Deuteronomy,
	"申命记":     bible.Deuteronomy,
	"約書亞記":    bible.Joshua,
	"约书亚记":    bible.Joshua,
	"士师记":     bible.Judges,
	"士師記":     bible.Judges,
	"路得記":     bible.Ruth,
	"路得记":     bible.Ruth,
	"撒母耳記上":   bible.Samuel1,
	"撒母耳记上":   bible.Samuel1,
	"撒母耳記下":   bible.Samuel2,
	"撒母耳记下":   bible.Samuel2,
	"列王紀上":    bible.Kings1,
	"列王纪上":    bible.Kings1,
	"列王紀下":    bible.Kings2,
	"列王纪下":    bible.Kings2,
	"历代志上":    bible.Chronicles1,
	"歷代志上":    bible.Chronicles1,
	"历代志下":    bible.Chronicles2,
	"歷代志下":    bible.Chronicles2,
	"以斯拉記":    bible.Ezra,
	"以斯拉记":    bible.Ezra,
	"尼希米記":    bible.Nehemiah,
	"尼希米记":    bible.Nehemiah,
	"以斯帖記":    bible.Esther,
	"以斯帖记":    bible.Esther,
	"約伯記":     bible.Job,
	"约伯记":     bible.Job,
	"詩篇":      bible.Psalm,
	"诗篇":      bible.Psalm,
	"箴言":      bible.Proverbs,
	"传道书":     bible.Ecclesiastes,
	"傳道書":     bible.Ecclesiastes,
	"雅歌":      bible.SongOfSongs,
	"以賽亞書":    bible.Isaiah,
	"以赛亚书":    bible.Isaiah,
	"耶利米书":    bible.Jeremiah,
	"耶利米書":    bible.Jeremiah,
	"耶利米哀歌":   bible.Lamentations,
	"以西結書":    bible.Ezekiel,
	"以西结书":    bible.Ezekiel,
	"但以理书":    bible.Daniel,
	"但以理書":    bible.Daniel,
	"何西阿书":    bible.Hosea,
	"何西阿書":    bible.Hosea,
	"約珥書":     bible.Joel,
	"约珥书":     bible.Joel,
	"阿摩司书":    bible.Amos,
	"阿摩司書":    bible.Amos,
	"俄巴底亚书":   bible.Obadiah,
	"俄巴底亞書":   bible.Obadiah,
	"約拿書":     bible.Jonah,
	"约拿书":     bible.Jonah,
	"弥迦书":     bible.Micah,
	"彌迦書":     bible.Micah,
	"那鴻書":     bible.Nahum,
	"那鸿书":     bible.Nahum,
	"哈巴谷书":    bible.Habakkuk,
	"哈巴谷書":    bible.Habakkuk,
	"西番雅书":    bible.Zephaniah,
	"西番雅書":    bible.Zephaniah,
	"哈該書":     bible.Haggai,
	"哈该书":     bible.Haggai,
	"撒迦利亚书":   bible.Zechariah,
	"撒迦利亞書":   bible.Zechariah,
	"玛拉基书":    bible.Malachi,
	"瑪拉基書":    bible.Malachi,
	"馬太福音":    bible.Matthew,
	"马太福音":    bible.Matthew,
	"馬可福音":    bible.Mark,
	"马可福音":    bible.Mark,
	"路加福音":    bible.Luke,
	"約翰福音":    bible.John,
	"约翰福音":    bible.John,
	"使徒行传":    bible.Acts,
	"使徒行傳":    bible.Acts,
	"罗马书":     bible.Romans,
	"羅馬書":     bible.Romans,
	"哥林多前书":   bible.Corinthians1,
	"哥林多前書":   bible.Corinthians1,
	"哥林多后书":   bible.Corinthians2,
	"哥林多後書":   bible.Corinthians2,
	"加拉太书":    bible.Galatians,
	"加拉太書":    bible.Galatians,
	"以弗所书":    bible.Ephesians,
	"以弗所書":    bible.Ephesians,
	"腓立比書":    bible.Philippians,
	"腓利比書":    bible.Philippians,
	"腓立比书":    bible.Philippians,
	"歌罗西书":    bible.Colossians,
	"歌羅西書":    bible.Colossians,
	"帖撒罗尼迦前书": bible.Thessalonians1,
	"帖撒羅尼迦前書": bible.Thessalonians1,
	"帖撒罗尼迦后书": bible.Thessalonians2,
	"帖撒羅尼迦後書": bible.Thessalonians2,
	"提摩太前书":   bible.Timothy1,
	"提摩太前書":   bible.Timothy1,
	"提摩太后书":   bible.Timothy2,
	"提摩太後書":   bible.Timothy2,
	"提多书":     bible.Titus,
	"提多書":     bible.Titus,
	"腓利門書":    bible.Philemon,
	"腓利门书":    bible.Philemon,
	"希伯來書":    bible.Hebrews,
	"希伯来书":    bible.Hebrews,
	"雅各书":     bible.James,
	"雅各書":     bible.James,
	"彼得前书":    bible.Peter1,
	"彼得前書":    bible.Peter1,
	"彼得后书":    bible.Peter2,
	"彼得後書":    bible.Peter2,
	"約翰壹書":    bible.John1,
	"约翰一书":    bible.John1,
	"約翰貳書":    bible.John2,
	"约翰二书":    bible.John2,
	"約翰參書":    bible.John3,
	"约翰三书":    bible.John3,
	"犹大书":     bible.Jude,
	"猶大書":     bible.Jude,
	"启示录":     bible.Revelation,
	"啟示錄":     bible.Revelation,
}

var chineseBookAbbrs = map[string]bible.BookID{
	"创":  bible.Genesis,
	"創":  bible.Genesis,
	"出":  bible.Exodus,
	"利":  bible.Leviticus,
	"民":  bible.Numbers,
	"申":  bible.Deuteronomy,
	"书":  bible.Joshua,
	"書":  bible.Joshua,
	"士":  bible.Judges,
	"得":  bible.Ruth,
	"撒上": bible.Samuel1,
	"撒下": bible.Samuel2,
	"王上": bible.Kings1,
	"王下": bible.Kings2,
	"代上": bible.Chronicles1,
	"代下": bible.Chronicles2,
	"拉":  bible.Ezra,
	"尼":  bible.Nehemiah,
	"斯":  bible.Esther,
	"伯":  bible.Job,
	"詩":  bible.Psalm,
	"诗":  bible.Psalm,
	"箴":  bible.Proverbs,
	"传":  bible.Ecclesiastes,
	"傳":  bible.Ecclesiastes,
	"歌":  bible.SongOfSongs,
	"賽":  bible.Isaiah,
	"赛":  bible.Isaiah,
	"耶":  bible.Jeremiah,
	"哀":  bible.Lamentations,
	"結":  bible.Ezekiel,
	"结":  bible.Ezekiel,
	"但":  bible.Daniel,
	"何":  bible.Hosea,
	"珥":  bible.Joel,
	"摩":  bible.Amos,
	"俄":  bible.Obadiah,
	"拿":  bible.Jonah,
	"弥":  bible.Micah,
	"彌":  bible.Micah,
	"鴻":  bible.Nahum,
	"鸿":  bible.Nahum,
	"哈":  bible.Habakkuk,
	"番":  bible.Zephaniah,
	"該":  bible.Haggai,
	"该":  bible.Haggai,
	"亚":  bible.Zechariah,
	"亞":  bible.Zechariah,
	"玛":  bible.Malachi,
	"瑪":  bible.Malachi,
	"太":  bible.Matthew,
	"可":  bible.Mark,
	"路":  bible.Luke,
	"約":  bible.John,
	"约":  bible.John,
	"徒":  bible.Acts,
	"罗":  bible.Romans,
	"羅":  bible.Romans,
	"林前": bible.Corinthians1,
	"林后": bible.Corinthians2,
	"林後": bible.Corinthians2,
	"加":  bible.Galatians,
	"弗":  bible.Ephesians,
	"腓":  bible.Philippians,
	"西":  bible.Colossians,
	"帖前": bible.Thessalonians1,
	"帖后": bible.Thessalonians2,
	"帖後": bible.Thessalonians2,
	"提前": bible.Timothy1,
	"提后": bible.Timothy2,
	"提後": bible.Timothy2,
	"多":  bible.Titus,
	"門":  bible.Philemon,
	"门":  bible.Philemon,
	"來":  bible.Hebrews,
	"来":  bible.Hebrews,
	"雅":  bible.James,
	"彼前": bible.Peter1,
	"彼后": bible.Peter2,
	"彼後": bible.Peter2,
	"約一": bible.John1,
	"约一": bible.John1,
	"約二": bible.John2,
	"约二": bible.John2,
	"約三": bible.John3,
	"约三": bible.John3,
	"犹":  bible.Jude,
	"猶":  bible.Jude,
	"启":  bible.Revelation,
	"啟":  bible.Revelation,
}

var chineseParser *parser

func init() {
	fullDict := new(dict)
	for name, id := range chineseBookNames {
		fullDict.add(name, id, true)
	}

	abbrDict := new(dict)
	for name, id := range chineseBookAbbrs {
		abbrDict.add(name, id, true)
	}

	numDict := new(dict)
	for i := 0; i < 200; i++ {
		numDict.add(strconv.Itoa(i), i, true)
	}

	chapterDict := new(dict)
	chapterDict.add("章", nil, true)
	chapterDict.add("篇", nil, true)
	verseDict := new(dict)
	verseDict.add("节", nil, true)
	verseDict.add("節", nil, true)

	ss := &spaceSkipper{true}
	ccp := new(charCheckerParser)

	chineseParser = newParser([]tokenizerParser{
		fullDict, abbrDict, numDict, chapterDict, verseDict, ss, ccp,
	})

	for _, p := range chinesePatterns {
		err := chineseParser.addRule(p.pattern, p.ref)
		if err != nil {
			panic(fmt.Sprintf("Error in parsing pattern %s: %v", p.pattern, err))
		}
	}
}
