package main

import (
	"fmt"
	"math/rand"
	"time"
)

type player struct {
	id    int
	point int
	dice  int
}

func Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func GamePlay(p []player) []player {
	// fmt.Println("gameplay")
	for i, v := range p {
		if v.dice > 0 {
			res := Roll()
			if res == 1 {
				next := (i + 1) % len(p)
				p[next].dice += 1
				p[i].dice -= 1
			} else if res == 6 {
				p[i].point += 1
				p[i].dice -= 1
			}
		}
	}
	return p
}

func FinishedPlayer(p []player) []player {
	// fmt.Println("finishplayer")
	var new []player
	for i, v := range p {
		if v.dice > 0 {
			new = append(new, p[i])
		}
	}
	return new
}

func Start(p []player) {
	round := 1
	for len(p) > 1 {
		resplay := GamePlay(p)
		fmt.Println("round ", round, ":", p)
		round += 1
		p = FinishedPlayer(resplay)
	}
	if p != nil {
		fmt.Println("winner : player", p[0].id, " point :", p[0].point)
	} else {
		fmt.Println("no winner")
	}
}

func main() {
	n := 3
	m := 4
	players := make([]player, n)
	for i := 0; i < len(players); i++ {
		players[i].id = i + 1
		players[i].dice = m
	}
	Start(players)
}
