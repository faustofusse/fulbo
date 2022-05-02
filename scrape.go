package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

var baseUrl string = "https://www.promiedos.com.ar/"

func findString(s []string, comparer func(string) bool) string {
    for _, v := range s {
        if comparer(v) {
            return v
        }
    }
    return ""
}

func extractTeam(tr *colly.HTMLElement, column int) team {
    srcExtractor := func(_ int, img *goquery.Selection) string { return img.AttrOr("src", "(Imagen equipo)") }
    homeImages := tr.DOM.Find(fmt.Sprintf("td:nth-child(%d) img", column)).Map(srcExtractor)
    isCountry := func(s string) bool { return strings.Contains(s, "/ps/") }
    isClub := func(s string) bool { return !isCountry(s) }
    return team {
        Name: tr.DOM.Find(fmt.Sprintf("td:nth-child(%d) span", column)).Text(),
        Image: findString(homeImages, isClub),
        CountryImage: findString(homeImages, isCountry),
    }
}

func extractLeague(container *colly.HTMLElement, table *colly.HTMLElement) league {
    number, _ := strconv.Atoi(table.DOM.Parent().Parent().AttrOr("class", "-1"))
    return league {
        Number: number,
        Name: container.DOM.Find(fmt.Sprintf("div.cuadrono%d", number)).AttrOr("copa", "(Nombre liga)"),
        Image: table.DOM.Find("tr.tituloin img").AttrOr("src", "(Imagen liga)"),
        Url: table.DOM.Find("tr.tituloin a").AttrOr("href", "(Url liga)"),
        Games: []game{},
    }
}

func extractGame(tr *colly.HTMLElement) game {
    extractInt := func(column int) int { result, _ := strconv.Atoi(tr.DOM.Find(fmt.Sprintf("td:nth-child(%d)", column)).Text()); return result }
    return game{
        Time: strings.TrimSpace( tr.DOM.Find("td:nth-child(1)").Text() ),
        HomeTeam: extractTeam(tr, 2),
        AwayTeam: extractTeam(tr, 5),
        HomeScore: extractInt(3),
        AwayScore: extractInt(4),
        Url: tr.DOM.Find("td:nth-child(5) > a").AttrOr("href", "(Url partido)"),
    }
}

func extractGames(table *colly.HTMLElement) []game {
    games := []game{}
    table.ForEach("tr", func(i int, tr *colly.HTMLElement) {
        // Si la fila no es la de un partido, skipeo
        class := tr.DOM.AttrOr("class", "")
        if i < 2 || class == "choy" || class == "goles" || class == "diapart"{ return }
        // Extraigo el partido
        newGame := extractGame(tr)
        // Agrego el partido a la liga
        games = append(games, newGame)
    })
    return games
}

func extractLeagues() []league {
    leagues := []league{}
    c := colly.NewCollector()
    c.OnHTML("div#principal", func(container *colly.HTMLElement) {
        container.ForEach("div#partidos table", func(i int, table *colly.HTMLElement) {
            // Extraigo la liga
            newLeague := extractLeague(container, table)
            // Recorro la tabla de partidos
            newLeague.Games = extractGames(table)
            // Agrego la liga a la lista
            leagues = append(leagues, newLeague)
        })
    })
    c.Visit(baseUrl)
    return leagues
}


