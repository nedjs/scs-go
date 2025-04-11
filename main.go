package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	// not sure we still need to do this for optimal performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	options := ScsOptions{
		Parallel:      false,
		Deterministic: false,
	}
	timeOnly := false
	verbose := false
	wordsArgOffset := 1
	for ; wordsArgOffset < len(os.Args); wordsArgOffset++ {
		arg := os.Args[wordsArgOffset]
		if arg[0] == '-' {
			if arg == "-h" || arg == "--help" {
				printHelp()
				os.Exit(0)
			} else {
				flags := strings.Split(arg[1:], "")
				for _, flag := range flags {
					if flag == "v" {
						verbose = true
					} else if flag == "t" {
						timeOnly = true
					} else if flag == "p" {
						options.Parallel = true
					} else if flag == "d" {
						options.Deterministic = true
					} else {
						fmt.Printf("Unparseable flag: %s\n", arg)
						printHelp()
						os.Exit(1)
					}
				}
			}
		} else {
			break
		}
	}

	words := os.Args[wordsArgOffset:]

	if len(words) < 2 {
		fmt.Printf("Please provide at least two words to compare\n")
		printHelp()
		os.Exit(1)
	}

	start := time.Now()
	result := scs(words, options)
	elapsed := time.Since(start)
	valid, _ := validate(result, words)
	if !valid {
		fmt.Printf("Invalid result: %s\n", result)
		os.Exit(1)
	}

	if timeOnly {
		fmt.Printf("%d\n", elapsed.Nanoseconds())
	} else if verbose {
		fmt.Printf("Shortest common superstring: %s\n", result)
		fmt.Printf("Length of scs: %d\n", len(result))
		fmt.Printf("Elapsed time: %s\n", elapsed)
	} else {
		fmt.Printf("%s\n", result)
	}
}

func printHelp() {
	fmt.Printf("Usage: scs [-flags] <word1> <word2> [<word3> ...]\n")
	fmt.Printf("Flags:\n")
	fmt.Printf("  -v: Verbose output\n")
	fmt.Printf("  -t: Only print the elapsed time in nanoseconds\n")
	fmt.Printf("  -p: Use parallel processing\n")
	fmt.Printf("  -d: Deterministic output (same words will always produce the same result regardless of order)\n")
	fmt.Printf("  -h, --help: Show this help message\n")
	fmt.Printf("\nExample:\n")
	fmt.Printf("  scs -pv jack apple maven\n")
}
