package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

// func isTmux() bool {
// 	// for testing, set TMUX_TEST true/false
// 	if os.Getenv("TMUX_TEST") == "false" {
// 		return false
// 	}
// 	if os.Getenv("TMUX_TEST") == "true" {
// 		return true
// 	}
// 	// otherwise, determine from TERM, TMUX variables
// 	return (os.Getenv("TERM") == "screen" || len(os.Getenv("TMUX")) > 0)
// }

// func headerEscape() string {
// 	if isTmux() {
// 		return "\x1bPtmux;\x1b\x1b]1337"
// 	}
// 	return "\x1b]1337"
// }

// func footerEscape() string {
// 	if isTmux() {
// 		return "\a\x1b\\\n"
// 	}
// 	return "\a\033[AHolaaaa\n"
// }

func printableImage(url string) string {
    // open file
    file, _ := os.Open(url)
    defer file.Close()
    // get file info
    info, _ := file.Stat()
    // read file
    data := make([]byte, info.Size())
    file.Read(data)
    // convert to base64
    str := base64.StdEncoding.EncodeToString(data)
    // print image
    // fmt.Printf("%s;File=inline=1;width=2:%s%s", headerEscape(), str, footerEscape())
    // fmt.Printf("\033]1337;File=inline=1;width=2:%s\a", str)
    return fmt.Sprintf("\033]1337;File=inline=1;width=2:%s\a", str)
}

func printImage(url string) {
    fmt.Printf("%s", printableImage(url))
}
