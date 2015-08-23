package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: abv og fg")
}

// Abv computes the alcohol by volume given an original gravity, og, and a final
// gravity, fg.
func Abv(og float64, fg float64) float64 {
	// Hall, Michael. "Brew by the Numbers: Add Up What's in Your Beer." Zymurgy
	// Summer 1995, vol. 18, no. 2.
	return (76.08 * (og - fg) / (1.775 - og)) * (fg / 0.794)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		usage()
		os.Exit(64) // EX_USAGE
	}
	og, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(65) // EX_DATAERR
	}

	fg, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(65) // EX_DATAERR
	}
	abv := Abv(og, fg)
	abvStr := strconv.FormatFloat(abv, 'f', 3, 64)
	fmt.Println(abvStr)
}
