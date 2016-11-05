# Doors

**_Doors_** is intended to be a super simple test of the [Monty Hall Problem](https://en.wikipedia.org/wiki/Monty_Hall_problem), specifially to [Marilyn vos Savant](https://en.wikipedia.org/wiki/Marilyn_vos_Savant)'s suggested solution of switching'every time. Please **note**, this is by no means commentary. The intent is purely for fun of testing.

## Dependencies

[Go](https://golang.org/dl/)

## Build

```bash
$ go get -v ./...
$ go build
```



## Usage

```bash
$ ./doors -g 3000 -i 10

# sample output
total wins:     20017
avg wins:       2001
avg percentage: 0.667
```

When run, a defined number of _games_ are generated, each with 3 doors. Of the 3 doors, 2 have a **goat** and 1 has a **car** behind them. In the original spirit of the problem, a contestant picks a door. In this case, we rely on a (psuedo-)randomly generated number between 0 and 2. The host then opens a door, of the 2 not selected, showing a goat. For this… we loop through the doors ignoring the the one previusly selected, looking for the first one with a goat. Per the suggested solution, we should then select the remaining door. This is achieved by looping through the doors and selecting the door that was neither selected nor the "opened door". If this final door is a "car" door, then we have a win.

For the outputs:

* **total wins** is the aggregate sum of game wins for _all_ iterations. 
* **avg wins** is the average of the aggregate wins… i.e. _total wins / iterations_
* **avg percentage** is the average percent of game wins (in decimals)… i.e. _avg wins / games_

It's worth noting, there may be no reason to include iterations. The intent here was to try and run more games of the same count to see variations in the results.

