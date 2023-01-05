package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/akamensky/argparse"
)

const lowerChars = "abcdefghijklmnopqrstuvwxyz"
const upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberChars = "1234567890"
const specialChars = "!@#$%^&*()+"

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
	argNumbers := parser.Flag("n", "numbers", &argparse.Options{Required: false, Default: false, Help: "Use number chars"})
	argSpecial := parser.Flag("s", "special", &argparse.Options{Required: false, Default: false, Help: "Use special chars"})
	argLower := parser.Flag("o", "lower", &argparse.Options{Required: false, Default: false, Help: "Use only lower chars"})
	argUpper := parser.Flag("u", "upper", &argparse.Options{Required: false, Default: false, Help: "Use only UPPER chars"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// loop count and check args
	var chars = lowerChars + upperChars
	for i := 1; i <= *argCount; i++ {
		if *argSpecial {
			chars += specialChars
		}
		if *argNumbers {
			chars += numberChars
		}
		if *argLower {
			chars = strings.ToLower(chars)
		}
		if *argUpper {
			chars = strings.ToUpper(chars)
		}
		fmt.Println(RandStringBytes(*argLength, chars))
	}

}
