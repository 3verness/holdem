package holdem

import "testing"

func TestNewPoker(t *testing.T) {
	type args struct {
		rank int32
		suit int32
	}
	tests := []struct {
		name string
		args args
		want Poker
	}{
		{
			"A Hearts",
			args{int32(12), int32(1)},
			Poker(0b00010000000000000010100111000001),
		},
		{
			"K Spades",
			args{int32(11), int32(8)},
			Poker(0b00001000000000000010010110111000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoker(tt.args.rank, tt.args.suit); got != tt.want {
				t.Errorf("NewPoker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPokerFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Poker
	}{
		{
			"A Hearts",
			args{"A Hearts"},
			Poker(0b00010000000000000010100111000001),
		},
		{
			"K Spades",
			args{"K Spades"},
			Poker(0b00001000000000000010010110111000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPokerFromString(tt.args.str); got != tt.want {
				t.Errorf("NewPokerFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoker_Unicode(t *testing.T) {
	tests := []struct {
		name string
		p    Poker
		want string
	}{
		{
			"A Hearts",
			Poker(0b00010000000000000010100111000001),
			"♥A",
		},
		{
			"K Spades",
			Poker(0b00001000000000000010010110111000),
			"♠K",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Unicode(); got != tt.want {
				t.Errorf("Unicode() = %v, want %v", got, tt.want)
			}
		})
	}
}
