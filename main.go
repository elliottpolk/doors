// Copyright 2016 Elliott Polk. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/elliottpolk/doors/game"
	"github.com/urfave/cli"
)

const (
	GamesFlag string = "games"
	IterFlag  string = "iters"
)

func main() {
	app := cli.NewApp()
	app.Name = "confgr"
	app.Usage = "a simple test of the Monty Hall problem"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  fmt.Sprintf("%s, g", GamesFlag),
			Value: 10,
			Usage: "number of games to simulate",
		},
		cli.IntFlag{
			Name:  fmt.Sprintf("%s, i", IterFlag),
			Value: 1,
			Usage: "number of iterations",
		},
	}
	app.Action = func(context *cli.Context) {
		context.Command.VisibleFlags()

		cnt := context.Int(GamesFlag)
		iters := context.Int(IterFlag)

		var wg sync.WaitGroup

		totalWins := 0
		for i := 0; i < iters; i++ {
			wg.Add(1)
			go func() {
				wins := game.Play(cnt)
				totalWins += wins
				wg.Done()
			}()
		}

		wg.Wait()

		avg := totalWins / iters
		per := float64(avg) / float64(cnt)

		fmt.Printf("total wins:     %d\n", totalWins)
		fmt.Printf("avg wins:       %d\n", avg)
		fmt.Printf("avg percentage: %v\n", per)
	}

	app.Run(os.Args)
}
