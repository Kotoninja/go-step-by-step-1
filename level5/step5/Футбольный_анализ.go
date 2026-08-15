package main

import (
	"sort"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
	newPlayer := &Player{Name: name, Goals: goals, Misses: misses, Assists: assists}
	newPlayer.CalculateRating()
	return *newPlayer
}

func (p *Player) CalculateRating() {
	if p.Misses == 0 {
		p.Rating = float64(p.Goals) + float64(p.Assists) / 2
	} else {
		p.Rating = (float64(p.Goals) + float64(p.Assists) / 2) / float64(p.Misses)
	}
}

func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		firstPlayer := players[i]
		secondPlayer := players[j]
		if firstPlayer.Rating == secondPlayer.Rating {
			return firstPlayer.Name < secondPlayer.Name
		}
		return float64(firstPlayer.Goals) > float64(secondPlayer.Goals)
	})

	return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		firstPlayer := players[i]
		secondPlayer := players[j]
		if firstPlayer.Rating == secondPlayer.Rating {
			return firstPlayer.Name < secondPlayer.Name
		}
		return firstPlayer.Rating > secondPlayer.Rating
	})

	return players
}

func gmSort(players []Player) []Player {
    sort.Slice(players, func(i, j int) bool {
        ratioI := float64(players[i].Goals) / float64(max(players[i].Misses, 1))
        ratioJ := float64(players[j].Goals) / float64(max(players[j].Misses, 1))
        
        if ratioI == ratioJ {
            return players[i].Name < players[j].Name
        }
        return ratioI > ratioJ
    })
    return players
}