package main

// ├ │ └  ── ╭──╮ ╰───╯   

import (
	"fmt"
	"strings"
)

func BoxWidth(homeTeamLength int, awayTeamLength int) int {
    return 2 + 5 + 3 + homeTeamLength + 13 + awayTeamLength + 2
}

func (l *league) Header(homeTeamLength int, awayTeamLength int) string {
    return fmt.Sprintf("╭─ %s %s %s╮",
        printableImage(l.Image),
        l.Name,
        strings.Repeat("─", BoxWidth(homeTeamLength, awayTeamLength) - len(l.Name) - 8),
    )
}

func (l *league) Footer(homeTeamLength int, awayTeamLength int) string {
    return fmt.Sprintf("╰%s╯",
        strings.Repeat("─", BoxWidth(homeTeamLength, awayTeamLength) - 2),
    )
}

func (g *game) ToString(homeTeamLength int, awayTeamLength int) string {
    return fmt.Sprintf("│ %s%s │ %s%s %s %d - %d %s %s%s │",
        strings.Repeat(" ", 5 - len(g.Time)),
        g.Time,
        strings.Repeat(" ", homeTeamLength - len(g.HomeTeam.Name)),
        g.HomeTeam.Name,
        printableImage(g.HomeTeam.Image),
        g.HomeScore, g.AwayScore,
        printableImage(g.AwayTeam.Image),
        g.AwayTeam.Name,
        strings.Repeat(" ", awayTeamLength - len(g.AwayTeam.Name)),
    )
}

func LongestTeam(leagues []league, away bool) int {
    longest := 0
    for _, l := range leagues {
        for _, g := range l.Games {
            var length int
            if away {
                length = len(g.AwayTeam.Name)
            } else {
                length = len(g.HomeTeam.Name)
            }
            if length > longest {
                longest = length
            }
        }
    }
    return longest
}

func printLeagues(leagues []league) {
    longestHomeTeam := LongestTeam(leagues, false)
    longestAwayTeam := LongestTeam(leagues, true)
    for _, l := range leagues {
        fmt.Printf("%s\n", l.Header(longestHomeTeam, longestAwayTeam))
        for _, g := range l.Games {
            fmt.Printf("%s\n", g.ToString(longestHomeTeam, longestAwayTeam))
        }
        fmt.Printf("%s\n", l.Footer(longestHomeTeam, longestAwayTeam))
    }
}

func main() {
    leagues := extractLeagues()
    printLeagues(leagues)
}
