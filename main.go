package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/akamensky/argparse"
)

const defaultChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ#$?!:=@"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = defaultChars[rand.Intn(len(defaultChars))]
	}
	return string(b)
}

func main() {
	// random seed
	rand.Seed(time.Now().UnixNano())

	// create parser objects
	parser := argparse.NewParser("my-gopass", "simple password generator written in GO")

	// creating arguments
	argLength := parser.Int("l", "length", &argparse.Options{Required: false, Default: 32, Help: "Password length"})
	argCount := parser.Int("c", "count", &argparse.Options{Required: false, Default: 4, Help: "Password count"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// loop count
	for i := 1; i <= *argCount; i++ {
		fmt.Println(RandStringBytes(*argLength))
	}

}
