package brave

func addBook(p *bookParser, b BookId, names ...string) {
	for _, n := range names {
		p.AddBook(b, n)
	}
}

func newChineseBookParser() *bookParser {
	p := newBookParser()
	addBook(p, Genesis, "创世记", "創世記", "创", "創")
	addBook(p, Exodus, "出埃及记", "出埃及記", "出", "出")
	addBook(p, Leviticus, "利未记", "利未記", "利", "利")
	addBook(p, Numbers, "民数记", "民數記", "民", "民")
	addBook(p, Deuteronomy, "申命记", "申命記", "申", "申")
	addBook(p, Joshua, "约书亚记", "約書亞記", "书", "書")
	addBook(p, Judges, "士师记", "士師記", "士", "士")
	addBook(p, Ruth, "路得记", "路得記", "得", "得")
	addBook(p, Samuel1, "撒母耳记上", "撒母耳記上", "撒", "	撒上")
	addBook(p, Samuel2, "撒母耳记下", "撒母耳記下", "撒", "	撒下")
	addBook(p, Kings1, "列王纪上", "列王紀上", "王", "	王上")
	addBook(p, Kings2, "列王纪下", "列王紀下", "王", "	王下")
	addBook(p, Chronicles1, "历代志上", "歷代志上", "代", "	代上")
	addBook(p, Chronicles2, "历代志下", "歷代志下", "代", "	代下")
	addBook(p, Ezra, "以斯拉记", "以斯拉記", "拉", "拉")
	addBook(p, Nehemiah, "尼希米记", "尼希米記", "尼", "尼")
	addBook(p, Esther, "以斯帖记", "以斯帖記", "斯", "斯")
	addBook(p, Job, "约伯记", "約伯記", "伯", "伯")
	addBook(p, Psalm, "诗篇", "詩篇", "诗", "詩")
	addBook(p, Proverbs, "箴言", "箴言", "箴", "箴")
	addBook(p, Ecclesiastes, "传道书", "傳道書", "传", "傳")
	addBook(p, SongOfSongs, "雅歌", "雅歌", "歌", "歌")
	addBook(p, Isaiah, "以赛亚书", "以賽亞書", "赛", "賽")
	addBook(p, Jeremiah, "耶利米书", "耶利米書", "耶", "耶")
	addBook(p, Lamentations, "耶利米哀歌", "耶利米哀歌", "哀", "哀")
	addBook(p, Ezekiel, "以西结书", "以西結書", "结", "結")
	addBook(p, Daniel, "但以理书", "但以理書", "但", "但")
	addBook(p, Hosea, "何西阿书", "何西阿書", "何", "何")
	addBook(p, Joel, "约珥书", "約珥書", "珥", "珥")
	addBook(p, Amos, "阿摩司书", "阿摩司書", "摩", "摩")
	addBook(p, Obadiah, "俄巴底亚书", "俄巴底亞書", "俄", "俄")
	addBook(p, Jonah, "约拿书", "約拿書", "拿", "拿")
	addBook(p, Micah, "弥迦书", "彌迦書", "弥", "彌")
	addBook(p, Nahum, "那鸿书", "那鴻書", "鸿", "鴻")
	addBook(p, Habakkuk, "哈巴谷书", "哈巴谷書", "哈", "哈")
	addBook(p, Zephaniah, "西番雅书", "西番雅書", "番", "番")
	addBook(p, Haggai, "哈该书", "哈該書", "该", "該")
	addBook(p, Zechariah, "撒迦利亚书", "撒迦利亞書", "亚", "亞")
	addBook(p, Malachi, "玛拉基书", "瑪拉基書", "玛", "瑪")
	addBook(p, Matthew, "马太福音", "馬太福音", "太", "太")
	addBook(p, Mark, "马可福音", "馬可福音", "可", "可")
	addBook(p, Luke, "路加福音", "路加福音", "路", "路")
	addBook(p, John, "约翰福音", "約翰福音", "约", "約")
	addBook(p, Acts, "使徒行传", "使徒行傳", "徒", "徒")
	addBook(p, Romans, "罗马书", "羅馬書", "羅", "羅")
	addBook(p, Corinthians1, "哥林多前书", "哥林多前書", "林", "	林前")
	addBook(p, Corinthians2, "哥林多后书", "哥林多後書", "林", "	林後")
	addBook(p, Galatians, "加拉太书", "加拉太書", "加", "加")
	addBook(p, Ephesians, "以弗所书", "以弗所書", "弗", "弗")
	addBook(p, Philippians, "腓立比书", "腓立比書", "腓", "腓")
	addBook(p, Colossians, "歌罗西书", "歌羅西書", "西", "西")
	addBook(p, Thessalonians1, "帖撒罗尼迦前书", "帖撒羅尼迦前書", "帖", "	帖前")
	addBook(p, Thessalonians2, "帖撒罗尼迦后书", "帖撒羅尼迦後書", "帖", "	帖後")
	addBook(p, Timothy1, "提摩太前书", "提摩太前書", "提", "	提前")
	addBook(p, Timothy2, "提摩太后书", "提摩太後書", "提", "	提後")
	addBook(p, Titus, "提多书", "提多書", "多", "多")
	addBook(p, Philemon, "腓利门书", "腓利門書", "门", "門")
	addBook(p, Hebrews, "希伯来书", "希伯來書", "來", "來")
	addBook(p, James, "雅各书", "雅各書", "雅", "雅")
	addBook(p, Peter1, "彼得前书", "彼得前書", "彼", "	彼前")
	addBook(p, Peter2, "彼得后书", "彼得後書", "彼", "	彼後")
	addBook(p, John1, "约翰一书", "約翰一書", "约", "	約一")
	addBook(p, John2, "约翰二书", "約翰二書", "约", "	約二")
	addBook(p, John3, "约翰三书", "約翰三書", "约", "	約三")
	addBook(p, Jude, "犹大书", "猶大書", "犹", "猶")
	addBook(p, Revelation, "启示录", "啟示錄", "启", "啟")
	return p
}

