package holdem

var (
	//同花查找表
	flushTable map[int32]int32
	//杂色查找表
	unsuitedTable map[int32]int32
)

const (
	// RoyalStraightFlush 皇家同花顺，1
	RoyalStraightFlush = 7462
	// StraightFlush 同花顺，9
	StraightFlush = 7453
	// FourOfAKind 四条，13*12=156
	FourOfAKind = 7297
	// FullHouse 葫芦，13*12=156
	FullHouse = 7141
	// Flush 同花，C(5,13)-10=1277
	Flush = 5864
	// Straight 顺子，10
	Straight = 5854
	// ThreeOfAKind 三条，13*C(2,12)=858
	ThreeOfAKind = 4996
	// TwoPair 两对，C(2,13)*11=858
	TwoPair = 4138
	// Pair 对子，13*C(3,12)=2860
	Pair = 1278
	// HighCard 乌龙，C(5,13)-10=1277
	HighCard = 1
)

func init() {
	flushTable = map[int32]int32{}
	unsuitedTable = map[int32]int32{}

	//顺子，从TJQKA开始
	straights := [10]int32{
		0b1000000001111,
		0b0000000011111,
		0b0000000111110,
		0b0000001111100,
		0b0000011111000,
		0b0000111110000,
		0b0001111100000,
		0b0011111000000,
		0b0111110000000,
		0b1111100000000,
	}
	serials := [1277]int32{}

	//非顺子序列，从23456开始向上生成，参考 https://graphics.stanford.edu/~seander/bithacks.html#NextBitPermutation
	serial := int32(0b0000000011111)
	for i := 0; i < 1277; {
		serial = nextSerial(serial)
		isStraight := false
		for _, straight := range straights {
			if straight^serial == 0 {
				isStraight = true
				break
			}
		}

		if !isStraight {
			serials[i] = serial
			i++
		}
	}

	//插入同花顺与顺子
	for i, s := range straights {
		p := serialToPrime(s)
		flushTable[p] = int32(StraightFlush + i)
		unsuitedTable[p] = int32(Straight + i)
	}

	//插入同花和乌龙
	for i, s := range serials {
		p := serialToPrime(s)
		flushTable[p] = int32(Flush + i)
		unsuitedTable[p] = int32(HighCard + i)
	}

	//插入四条与葫芦，从22223与22233开始生成
	s := 0
	for major := 0; major < 13; major++ {
		for sub := 0; sub < 13; sub++ {
			if sub != major {
				p1 := ranksPrime[major] * ranksPrime[major] * ranksPrime[major] *
					ranksPrime[major] * ranksPrime[sub]
				p2 := ranksPrime[major] * ranksPrime[major] * ranksPrime[major] *
					ranksPrime[sub] * ranksPrime[sub]
				unsuitedTable[p1] = int32(FourOfAKind + s)
				unsuitedTable[p2] = int32(FullHouse + s)
				s++
			}
		}
	}

	//插入三条，从222234开始生成
	s = 0
	for major := 0; major < 13; major++ {
		for sub1 := 0; sub1 < 13; sub1++ {
			if sub1 != major {
				for sub2 := 0; sub2 < sub1; sub2++ {
					if sub2 != major {
						p := ranksPrime[major] * ranksPrime[major] * ranksPrime[major] *
							ranksPrime[sub1] * ranksPrime[sub2]
						unsuitedTable[p] = int32(ThreeOfAKind + s)
						s++
					}
				}
			}
		}
	}

	//插入两对，从22334开始生成
	s = 0
	for major1 := 0; major1 < 13; major1++ {
		for major2 := 0; major2 < major1; major2++ {
			for sub := 0; sub < 13; sub++ {
				if sub != major1 && sub != major2 {
					p := ranksPrime[major1] * ranksPrime[major1] * ranksPrime[major2] *
						ranksPrime[major2] * ranksPrime[sub]
					unsuitedTable[p] = int32(TwoPair + s)
					s++
				}
			}
		}
	}

	//插入对子，从22345开始生成
	s = 0
	for major := 0; major < 13; major++ {
		for sub1 := 0; sub1 < 13; sub1++ {
			if sub1 != major {
				for sub2 := 0; sub2 < sub1; sub2++ {
					if sub2 != major {
						for sub3 := 0; sub3 < sub2; sub3++ {
							if sub3 != major {
								p := ranksPrime[major] * ranksPrime[major] * ranksPrime[sub1] *
									ranksPrime[sub2] * ranksPrime[sub3]
								unsuitedTable[p] = int32(Pair + s)
								s++
							}
						}
					}
				}
			}
		}
	}
}

func nextSerial(a int32) int32 {
	t := (a | (a - 1)) + 1
	return t | ((((t & -t) / (a & -a)) >> 1) - 1)
}

func serialToPrime(s int32) int32 {
	p := int32(1)
	for i := 0; i < 13; i++ {
		if s&(1<<i) != 0 {
			p *= ranksPrime[i]
		}
	}
	return p
}

func pokersToPrime(pokers []Poker) int32 {
	p := int32(1)
	for _, poker := range pokers {
		p *= poker.RankPrime()
	}
	return p
}

func serialScore(pokers []Poker) int32 {
	//判断是否为同花
	if (pokers[0]&pokers[1]&pokers[2]&pokers[3]&pokers[4])&0xF == 0 {
		return unsuitedTable[pokersToPrime(pokers)]
	} else {
		return flushTable[pokersToPrime(pokers)]
	}
}

func Score(pokers []Poker) int32 {
	maxS := int32(0)

	for i := 0; i < 7; i++ {
		for j := 0; j < i; j++ {
			ps := make([]Poker, 7)
			copy(ps, pokers)
			ps = append(ps[:i], ps[i+1:]...)
			ps = append(ps[:j], ps[j+1:]...)
			s := serialScore(ps)
			if s > maxS {
				maxS = s
			}
		}
	}
	return maxS
}

func ScoreType(s int32) string {
	switch {
	case s >= RoyalStraightFlush:
		return "Royal Straight Flush"
	case s >= StraightFlush:
		return "Straight Flush"
	case s >= FourOfAKind:
		return "Four of a Kind"
	case s >= FullHouse:
		return "Full House"
	case s >= Flush:
		return "Flush"
	case s >= Straight:
		return "Straight"
	case s >= ThreeOfAKind:
		return "Three of a Kind"
	case s >= TwoPair:
		return "Two Pair"
	case s >= Pair:
		return "Pair"
	case s >= HighCard:
		return "High Card"
	default:
		return "Unknown"
	}
}
