package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	}
	if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type MyInt int

func okTypeAssertion() error {
	var i any

	var mine MyInt = 20

	i = mine

	i2, ok := i.(int)
	if !ok {
		fmt.Printf("i2 %v \n", i2)
		return fmt.Errorf("unexpected type for i: %v", i)
	}

	fmt.Println(i2 + 1)

	return nil
}

type Team struct {
	name        string
	playerNames []string
}

type Wins struct {
	name     string
	teamWins int
}

type League struct {
	teams []*Team
	wins  []Wins
}

func (l *League) MatchResult(firstTeam, secondTeam string, firstScore, secondScore int) {
	if firstScore > secondScore {
		for i, v := range l.wins {
			if v.name == firstTeam {
				l.wins[i].teamWins++
			}
		}
	}
	if firstScore < secondScore {
		for i, v := range l.wins {
			if v.name == secondTeam {
				l.wins[i].teamWins++
			}
		}
	}
}

func (l League) Ranking() []string {
	slices.SortFunc(l.wins, func(a, b Wins) int {
		return b.teamWins - a.teamWins
	})
	res := []string{}
	for _, v := range l.wins {
		res = append(res, v.name)
	}
	return res
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	ranks := r.Ranking()
	for _, v := range ranks {
		io.WriteString(w, fmt.Sprintln(v))
	}
}

func main() {
	var it *IntTree
	it = it.Insert(21)
	it = it.Insert(213)
	it = it.Insert(54)
	it = it.Insert(1)
	it = it.Insert(12323)
	it = it.Insert(119)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(1))
	err := okTypeAssertion()
	if err != nil {
		fmt.Println(err)
	}

	teamA := &Team{
		name:        "Eagles",
		playerNames: []string{"Alice", "Bob", "Charlie"},
	}
	teamB := &Team{
		name:        "Tigers",
		playerNames: []string{"Dave", "Eve", "Frank"},
	}
	teamC := &Team{
		name:        "Wolves",
		playerNames: []string{"Grace", "Hank", "Ivy"},
	}

	// Create the league
	league := &League{
		teams: []*Team{teamA, teamB, teamC},
		wins: []Wins{
			{name: teamA.name, teamWins: 0},
			{name: teamB.name, teamWins: 0},
			{name: teamC.name, teamWins: 0},
		},
	}

	// Simulate some match results
	league.MatchResult(teamA.name, teamB.name, 3, 1) // Eagles win
	league.MatchResult(teamC.name, teamA.name, 2, 0) // Wolves win
	league.MatchResult(teamB.name, teamB.name, 4, 2) // Tigers win
	league.MatchResult(teamA.name, teamC.name, 1, 3) // Wolves win
	league.MatchResult(teamA.name, teamC.name, 1, 3) // Wolves win
	league.MatchResult(teamA.name, teamC.name, 1, 3) // Wolves win
	league.MatchResult(teamB.name, teamA.name, 2, 2) // Draw (no wins recorded)

	// Print rankings
	fmt.Println("Rankings:", league.Ranking())
	RankPrinter(league, os.Stdout)
}
