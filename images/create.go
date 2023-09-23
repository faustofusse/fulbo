package main

import (
	"os"
)

func main() {
    file, _ := os.Create("images.txt")
    defer file.Close()
    file.WriteString("hola\n")
    file.WriteString("adios\n")
}
