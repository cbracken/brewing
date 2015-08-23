package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type UnitSystem int

const (
	Us UnitSystem = iota
	Imperial
)

const usGallon = 3.78541
const imperialGallon = 4.54609

func usage() {
	fmt.Fprintln(os.Stderr, "usage: g2l [-imperial] gallons")
}

// GallonsToLitres returns the specifed volume in litres.
func GallonsToLitres(g float64, s UnitSystem) float64 {
	if s == Us {
		return g * usGallon
	}
	return g * imperialGallon
}

func main() {
	flag.Usage = usage
	imp := flag.Bool("imperial", false, "use imperial gallons")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(64) // EX_USAGE
	}

	g, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(65) // EX_DATAERR
	}

	s := Us
	if *imp {
		s = Imperial
	}

	l := GallonsToLitres(g, s)
	lStr := strconv.FormatFloat(l, 'f', 3, 64)
	fmt.Println(lStr)
}
