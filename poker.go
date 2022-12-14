package holdem

import "strings"

// Poker 扑克类
// |16bit rankBit||8bit rankPrime||4bit rank||4bit suit|
type Poker int32

var (
	ranks        = [13]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	ranksString  = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	ranksPrime   = [13]int32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	suits        = [4]int32{1, 2, 4, 8}
	suitsString  = [9]string{"", "Hearts", "Diamonds", "", "Clubs", "", "", "", "Spades"}
	suitsUnicode = [9]string{"", "\u2665", "\u2666", "", "\u2663", "", "", "", "\u2660"}
)

var (
	ranksMap = map[string]int32{}
	suitsMap = map[string]int32{}
)

// 初始化对照字典
func init() {
	for _, i := range ranks {
		ranksMap[ranksString[i]] = ranks[i]
	}

	for _, i := range suits {
		suitsMap[suitsString[i]] = i
	}
}

func NewPoker(rank int32, suit int32) Poker {
	rankBit := int32(1) << rank
	rankPrime := ranksPrime[rank]
	return Poker(suit | (rank << 4) | (rankPrime << 8) | (rankBit << 16))
}

// NewPokerFromString 从字符串生成扑克，输入如"10 Hearts"
func NewPokerFromString(str string) Poker {
	s := strings.Split(str, " ")
	suit := suitsMap[s[1]]
	rank := ranksMap[s[0]]
	return NewPoker(rank, suit)
}

// AllPoker 生成一付52张扑克
func AllPoker() []Poker {
	var pokers []Poker
	for _, s := range suits {
		for _, r := range ranks {
			pokers = append(pokers, NewPoker(r, s))
		}
	}
	return pokers
}

func (p Poker) Suit() int32 {
	return int32(p) & 0xF
}

func (p Poker) Rank() int32 {
	return (int32(p) >> 4) & 0xF
}

func (p Poker) RankBit() int32 {
	return (int32(p) >> 16) & 0xFFFF
}

func (p Poker) RankPrime() int32 {
	return (int32(p) >> 8) & 0xFF
}

func (p Poker) String() string {
	return ranksString[p.Rank()] + " " + suitsString[p.Suit()]
}

func (p Poker) Unicode() string {
	return suitsUnicode[p.Suit()] + ranksString[p.Rank()]
}
