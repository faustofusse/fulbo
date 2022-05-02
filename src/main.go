package main

import (
	"fmt"
)

func (g *game) ToString() string {
    return fmt.Sprintf("%s: %s %s (%d) - (%d) %s %s", g.Time, g.HomeTeam.Name, printableImage(g.HomeTeam.Image), g.HomeScore, g.AwayScore, printableImage(g.AwayTeam.Image), g.AwayTeam.Name)
}

func printLeagues(leagues []league) {
    for _, l := range leagues {
        fmt.Printf("%d: %s\n", l.Number, l.Name)
        for _, g := range l.Games {
            fmt.Printf("\t%s\n", g.ToString())
        }
    }
}

func main() {
    leagues := extractLeagues()
    printLeagues(leagues)
}
