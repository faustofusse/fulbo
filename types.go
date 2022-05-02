package main

type league struct {
    Number int
    Name string
    Image string
    Games []game
    Url string
}

type game struct {
    Time string
    HomeTeam team
    AwayTeam team
    HomeScore int
    AwayScore int
    Url string
}

type team struct {
    Name string
    Image string
    CountryImage string
}

