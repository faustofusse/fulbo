package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var baseDir string = "/Users/faustofusse/Documents/Software/terminal/fulbo/"

func isTmux() bool {
	// for testing, set TMUX_TEST true/false
	if os.Getenv("TMUX_TEST") == "false" {
		return false
	}
	if os.Getenv("TMUX_TEST") == "true" {
		return true
	}
	// otherwise, determine from TERM, TMUX variables
	return (os.Getenv("TERM") == "screen" || len(os.Getenv("TMUX")) > 0)
}

func headerEscape() string {
	if isTmux() {
		return "\x1bPtmux;\x1b\x1b]1337"
	}
	return "\x1b]1337"
}

func footerEscape() string {
	if isTmux() {
		return "\a\x1b\\"
	}
	// return "\a\033[A\n"
    return "\a"
}

func printableImage(url string) string {
    // open file
    file, err := os.Open(baseDir + url)
    if err != nil && errors.Is(err, os.ErrNotExist) {
        file, _ = os.Create(url)
        response, _ := http.Get(baseUrl + url)
        io.Copy(file, response.Body)
        response.Body.Close()
    }
    defer file.Close()
    // get file info
    info, _ := file.Stat()
    // read file
    data := make([]byte, info.Size())
    file.Read(data)
    // convert to base64
    str := base64.StdEncoding.EncodeToString(data)
    // print image
    return fmt.Sprintf("%s;File=inline=1;width=2:%s%s", headerEscape(), str, footerEscape())
}

func printImage(url string) {
    fmt.Printf("%s", printableImage(url))
}
