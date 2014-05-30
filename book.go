package brave

import (
	"github.com/saintfish/trie.go"
)

type BookId byte

const (
	InvalidBook BookId = iota

	Genesis
	Exodus
	Leviticus
	Numbers
	Deuteronomy
	Joshua
	Judges
	Ruth
	Samuel1
	Samuel2
	Kings1
	Kings2
	Chronicles1
	Chronicles2
	Ezra
	Nehemiah
	Esther
	Job
	Psalm
	Proverbs
	Ecclesiastes
	SongOfSongs
	Isaiah
	Jeremiah
	Lamentations
	Ezekiel
	Daniel
	Hosea
	Joel
	Amos
	Obadiah
	Jonah
	Micah
	Nahum
	Habakkuk
	Zephaniah
	Haggai
	Zechariah
	Malachi
	Matthew
	Mark
	Luke
	John
	Acts
	Romans
	Corinthians1
	Corinthians2
	Galatians
	Ephesians
	Philippians
	Colossians
	Thessalonians1
	Thessalonians2
	Timothy1
	Timothy2
	Titus
	Philemon
	Hebrews
	James
	Peter1
	Peter2
	John1
	John2
	John3
	Jude
	Revelation

	FirstBook = Genesis
	LastBook  = Revelation
	FirstOT   = Genesis
	LastOT    = Malachi
	FirstNT   = Matthew
	LastNT    = Revelation
)
const NumBooks = int(LastBook - FirstBook + 1)

type bookParser struct {
	trie *trie.Trie
}

func newBookParser() *bookParser {
	return &bookParser{trie.NewTrie()}
}

func (this *bookParser) AddBook(b BookId, names ...string) {
	for _, n := range names {
		this.trie.Add([]byte(n), b)
	}
}

type Book struct {
	Book       BookId
	Annotation Annotation
}

func (this *bookParser) parseFrom(input string, pos int) *Book {
	// TODO: Make trie work with string input
	prefix, found := this.trie.MatchLongestPrefixString(input[pos:])
	if found {
		return &Book{
			Book: prefix.Value.(BookId),
			Annotation: Annotation{
				Begin: pos,
				End:   pos + prefix.PrefixLength,
			},
		}
	}
	return nil
}

/*
func (this *bookParser) Parse(content string) []BookMatch {
	contentBytes := []byte(content)
	result := []BookMatch{}
	for i := range(content) {
		this.parsePrefix(contentBytes, i, &result)
	}
	return result
}
*/
