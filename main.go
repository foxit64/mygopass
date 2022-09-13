package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/akamensky/argparse"
)

const defaultChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const specialChars = "1234567890+*$%&()[]!:=@#{}"

func RandStringBytes(n int, chars string) string {
	rsb := make([]byte, n)
	for i := range rsb {
		rsb[i] = chars[rand.Intn(len(chars))]
	}
	return string(rsb)
}

func main() {
	// random seed
	rand.Seed(time.Now().UnixNano())

	// create parser objects
	parser := argparse.NewParser("my-gopass", "simple password generator written in GO")

	// creating arguments
	argLength := parser.Int("l", "length", &argparse.Options{Required: false, Default: 32, Help: "Password length"})
	argCount := parser.Int("c", "count", &argparse.Options{Required: false, Default: 4, Help: "Password count"})
	argSpecial := parser.Flag("s", "special", &argparse.Options{Required: false, Default: false, Help: "Use special chars"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// loop count
	for i := 1; i <= *argCount; i++ {
		if *argSpecial {
			var chars = defaultChars + specialChars
			fmt.Println(RandStringBytes(*argLength, chars))
		} else {
			var chars = defaultChars
			fmt.Println(RandStringBytes(*argLength, chars))
		}
	}

}
