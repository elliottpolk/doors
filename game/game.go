// Copyright 2016 Elliott Polk. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package game

import "github.com/elliottpolk/doors/rand"

const DefaultGameCount int = 10

type Game struct {
	Doors [DoorCount]*Door
}

func get(cnt int) []*Game {
	games := make([]*Game, cnt)

	for i := 0; i < cnt; i++ {
		game := &Game{
			[DoorCount]*Door{
				&Door{Goat},
				&Door{Goat},
				&Door{Goat},
			},
		}

		game.Doors[rand.Intn(DoorCount)].Is = Car
		games[i] = game
	}

	return games
}

func Play(cnt int) int {
	games := get(cnt)

	wins := 0

	for _, game := range games {
		doors := game.Doors

		selected := rand.Intn(DoorCount)

		var opened int
		for i, door := range doors {
			if i != selected && door.Is == Goat {
				opened = i
				break
			}
		}

		for i, door := range doors {
			if i != selected && i != opened {
				if door.Is == Car {
					wins++
				}
				break
			}
		}
	}

	return wins
}
