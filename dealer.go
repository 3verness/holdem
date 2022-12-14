package holdem

import (
	"math/rand"
	"time"
)

type Dealer struct {
	pokers []Poker
}

// Shuffle 洗牌
func (d *Dealer) Shuffle() {
	d.pokers = AllPoker()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.pokers), func(i, j int) {
		d.pokers[i], d.pokers[j] = d.pokers[j], d.pokers[i]
	})
}

// Cut 切牌
func (d *Dealer) Cut() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(len(d.pokers))
	buf := make([]Poker, len(d.pokers))
	copy(buf, d.pokers)
	d.pokers = append(buf[n:], buf[:n]...)
}

// Deal 发n张牌
func (d *Dealer) Deal(n int) []Poker {
	ps := make([]Poker, n)
	copy(ps, d.pokers[:n])
	d.pokers = d.pokers[n:]
	return ps
}
