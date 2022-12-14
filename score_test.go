package holdem

import "testing"

func Test_serialScore(t *testing.T) {
	type args struct {
		pokers []Poker
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"Royal Straight Flush",
			args{[]Poker{
				NewPokerFromString("A Hearts"),
				NewPokerFromString("K Hearts"),
				NewPokerFromString("Q Hearts"),
				NewPokerFromString("J Hearts"),
				NewPokerFromString("10 Hearts"),
			}},
			7462,
		},
		{
			"Max Straight Flush",
			args{[]Poker{
				NewPokerFromString("9 Hearts"),
				NewPokerFromString("K Hearts"),
				NewPokerFromString("Q Hearts"),
				NewPokerFromString("J Hearts"),
				NewPokerFromString("10 Hearts"),
			}},
			7461,
		},
		{
			"Min Straight Flush",
			args{[]Poker{
				NewPokerFromString("A Hearts"),
				NewPokerFromString("2 Hearts"),
				NewPokerFromString("3 Hearts"),
				NewPokerFromString("4 Hearts"),
				NewPokerFromString("5 Hearts"),
			}},
			7453,
		},
		{
			"Max Four of a Kind",
			args{[]Poker{
				NewPokerFromString("A Hearts"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("A Diamonds"),
				NewPokerFromString("K Hearts"),
			}},
			7452,
		},
		{
			"Min Four of a Kind",
			args{[]Poker{
				NewPokerFromString("2 Hearts"),
				NewPokerFromString("2 Spades"),
				NewPokerFromString("2 Clubs"),
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Hearts"),
			}},
			7297,
		},
		{
			"Max Full House",
			args{[]Poker{
				NewPokerFromString("A Hearts"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("K Diamonds"),
				NewPokerFromString("K Hearts"),
			}},
			7296,
		},
		{
			"Min Full House",
			args{[]Poker{
				NewPokerFromString("2 Hearts"),
				NewPokerFromString("2 Spades"),
				NewPokerFromString("2 Clubs"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("3 Hearts"),
			}},
			7141,
		},
		{
			"Max Flush",
			args{[]Poker{
				NewPokerFromString("9 Hearts"),
				NewPokerFromString("K Hearts"),
				NewPokerFromString("Q Hearts"),
				NewPokerFromString("J Hearts"),
				NewPokerFromString("A Hearts"),
			}},
			7140,
		},
		{
			"Min Flush",
			args{[]Poker{
				NewPokerFromString("7 Hearts"),
				NewPokerFromString("2 Hearts"),
				NewPokerFromString("3 Hearts"),
				NewPokerFromString("4 Hearts"),
				NewPokerFromString("5 Hearts"),
			}},
			5864,
		},
		{
			"Max Straight",
			args{[]Poker{
				NewPokerFromString("A Hearts"),
				NewPokerFromString("K Hearts"),
				NewPokerFromString("Q Hearts"),
				NewPokerFromString("J Hearts"),
				NewPokerFromString("10 Spades"),
			}},
			5863,
		},
		{
			"Min Straight",
			args{[]Poker{
				NewPokerFromString("A Hearts"),
				NewPokerFromString("2 Hearts"),
				NewPokerFromString("3 Hearts"),
				NewPokerFromString("4 Hearts"),
				NewPokerFromString("5 Spades"),
			}},
			5854,
		},
		{
			"Max Three of a Kind",
			args{[]Poker{
				NewPokerFromString("A Hearts"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("Q Diamonds"),
				NewPokerFromString("K Hearts"),
			}},
			5853,
		},
		{
			"Min Three of a Kind",
			args{[]Poker{
				NewPokerFromString("2 Hearts"),
				NewPokerFromString("2 Spades"),
				NewPokerFromString("2 Clubs"),
				NewPokerFromString("4 Diamonds"),
				NewPokerFromString("3 Hearts"),
			}},
			4996,
		},
		{
			"Max Two Pair",
			args{[]Poker{
				NewPokerFromString("A Hearts"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("K Clubs"),
				NewPokerFromString("Q Diamonds"),
				NewPokerFromString("K Hearts"),
			}},
			4995,
		},
		{
			"Min Two Pair",
			args{[]Poker{
				NewPokerFromString("2 Hearts"),
				NewPokerFromString("2 Spades"),
				NewPokerFromString("3 Clubs"),
				NewPokerFromString("4 Diamonds"),
				NewPokerFromString("3 Hearts"),
			}},
			4138,
		},
		{
			"Max Pair",
			args{[]Poker{
				NewPokerFromString("A Hearts"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("K Clubs"),
				NewPokerFromString("Q Diamonds"),
				NewPokerFromString("J Hearts"),
			}},
			4137,
		},
		{
			"Min Pair",
			args{[]Poker{
				NewPokerFromString("2 Hearts"),
				NewPokerFromString("2 Spades"),
				NewPokerFromString("3 Clubs"),
				NewPokerFromString("4 Diamonds"),
				NewPokerFromString("5 Hearts"),
			}},
			1278,
		},
		{
			"Max High Card",
			args{[]Poker{
				NewPokerFromString("9 Hearts"),
				NewPokerFromString("K Hearts"),
				NewPokerFromString("Q Hearts"),
				NewPokerFromString("J Hearts"),
				NewPokerFromString("A Spades"),
			}},
			1277,
		},
		{
			"Min High Card",
			args{[]Poker{
				NewPokerFromString("7 Clubs"),
				NewPokerFromString("2 Hearts"),
				NewPokerFromString("3 Hearts"),
				NewPokerFromString("4 Hearts"),
				NewPokerFromString("5 Hearts"),
			}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serialScore(tt.args.pokers); got != tt.want {
				t.Errorf("serialScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScore(t *testing.T) {
	type args struct {
		pokers []Poker
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"High Card",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("K Spades"),
				NewPokerFromString("J Clubs"),
				NewPokerFromString("7 Hearts"),
				NewPokerFromString("5 Diamonds"),
			}},
			1211,
		},
		{
			"Pair",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("J Clubs"),
				NewPokerFromString("7 Hearts"),
				NewPokerFromString("5 Diamonds"),
			}},
			4015,
		},
		{
			"Two Pair",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("J Clubs"),
				NewPokerFromString("J Diamonds"),
				NewPokerFromString("5 Diamonds"),
			}},
			4966,
		},
		{
			"Three of a Kind",
			args{[]Poker{
				NewPokerFromString("2 Clubs"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("A Diamonds"),
				NewPokerFromString("J Diamonds"),
				NewPokerFromString("5 Diamonds"),
			}},
			5827,
		},
		{
			"Straight",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("K Clubs"),
				NewPokerFromString("Q Hearts"),
				NewPokerFromString("J Diamonds"),
				NewPokerFromString("10 Diamonds"),
			}},
			5863,
		},
		{
			"Flush",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("10 Spades"),
				NewPokerFromString("7 Spades"),
				NewPokerFromString("4 Spades"),
				NewPokerFromString("3 Spades"),
				NewPokerFromString("2 Spades"),
			}},
			5921,
		},
		{
			"Full House",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("4 Spades"),
				NewPokerFromString("4 Clubs"),
				NewPokerFromString("4 Diamonds"),
				NewPokerFromString("2 Spades"),
				NewPokerFromString("2 Hearts"),
			}},
			7165,
		},
		{
			"Four of a Kind",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("A Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("5 Hearts"),
			}},
			7444,
		},
		{
			"Straight Flush",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("10 Spades"),
				NewPokerFromString("J Spades"),
				NewPokerFromString("Q Spades"),
				NewPokerFromString("K Spades"),
				NewPokerFromString("9 Spades"),
			}},
			7461,
		},
		{
			"Royal Straight Flush",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("10 Spades"),
				NewPokerFromString("J Spades"),
				NewPokerFromString("Q Spades"),
				NewPokerFromString("K Spades"),
				NewPokerFromString("A Spades"),
			}},
			7462,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Score(tt.args.pokers); got != tt.want {
				t.Errorf("Score() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreType(t *testing.T) {
	type args struct {
		s []Poker
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"High Card",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("K Spades"),
				NewPokerFromString("J Clubs"),
				NewPokerFromString("7 Hearts"),
				NewPokerFromString("5 Diamonds"),
			}},
			"High Card",
		},
		{
			"Pair",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("J Clubs"),
				NewPokerFromString("7 Hearts"),
				NewPokerFromString("5 Diamonds"),
			}},
			"Pair",
		},
		{
			"Two Pair",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("J Clubs"),
				NewPokerFromString("J Diamonds"),
				NewPokerFromString("5 Diamonds"),
			}},
			"Two Pair",
		},
		{
			"Three of a Kind",
			args{[]Poker{
				NewPokerFromString("2 Clubs"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("A Diamonds"),
				NewPokerFromString("J Diamonds"),
				NewPokerFromString("5 Diamonds"),
			}},
			"Three of a Kind",
		},
		{
			"Straight",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("K Clubs"),
				NewPokerFromString("Q Hearts"),
				NewPokerFromString("J Diamonds"),
				NewPokerFromString("10 Diamonds"),
			}},
			"Straight",
		},
		{
			"Flush",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("10 Spades"),
				NewPokerFromString("7 Spades"),
				NewPokerFromString("4 Spades"),
				NewPokerFromString("3 Spades"),
				NewPokerFromString("2 Spades"),
			}},
			"Flush",
		},
		{
			"Full House",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("4 Spades"),
				NewPokerFromString("4 Clubs"),
				NewPokerFromString("4 Diamonds"),
				NewPokerFromString("2 Spades"),
				NewPokerFromString("2 Hearts"),
			}},
			"Full House",
		},
		{
			"Four of a Kind",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("A Clubs"),
				NewPokerFromString("A Diamonds"),
				NewPokerFromString("A Spades"),
				NewPokerFromString("5 Hearts"),
			}},
			"Four of a Kind",
		},
		{
			"Straight Flush",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("10 Spades"),
				NewPokerFromString("J Spades"),
				NewPokerFromString("Q Spades"),
				NewPokerFromString("K Spades"),
				NewPokerFromString("9 Spades"),
			}},
			"Straight Flush",
		},
		{
			"Royal Straight Flush",
			args{[]Poker{
				NewPokerFromString("2 Diamonds"),
				NewPokerFromString("3 Diamonds"),
				NewPokerFromString("10 Spades"),
				NewPokerFromString("J Spades"),
				NewPokerFromString("Q Spades"),
				NewPokerFromString("K Spades"),
				NewPokerFromString("A Spades"),
			}},
			"Royal Straight Flush",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScoreType(Score(tt.args.s)); got != tt.want {
				t.Errorf("ScoreType() = %v, want %v", got, tt.want)
			}
		})
	}
}
