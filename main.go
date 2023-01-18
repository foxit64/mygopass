package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/akamensky/argparse"
)

const (
	lowerChars   string = "abcdefghijklmnopqrstuvwxyz"
	upperChars   string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  string = "1234567890"
	specialChars string = "!@#$%^&*()+"
	versionMajor int    = 0
	versionMinor int    = 3
	versionPatch int    = 0
)

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
	argLength := parser.Int("l", "length", &argparse.Options{Default: 32, Help: "Password length"})
	argCount := parser.Int("c", "count", &argparse.Options{Default: 4, Help: "Password count"})
	argNumbers := parser.Flag("n", "numbers", &argparse.Options{Default: false, Help: "Use number chars"})
	argSpecial := parser.Flag("s", "special", &argparse.Options{Default: false, Help: "Use special chars"})
	argLower := parser.Flag("o", "lower", &argparse.Options{Default: false, Help: "Use only lower chars"})
	argUpper := parser.Flag("u", "upper", &argparse.Options{Default: false, Help: "Use only UPPER chars"})
	argVersion := parser.Flag("", "version", &argparse.Options{Help: "Show version"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *argVersion {
		fmt.Printf("my-gopass v%v.%v.%v", versionMajor, versionMinor, versionPatch)
		os.Exit(0)
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
