package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/katre/techdice/dice"
	"github.com/katre/techdice/parser"
)

var seed = flag.Int64("seed", time.Now().UnixNano(), "Seed the RNG")

func main() {
	// Create a new roller.
	roller := dice.New(*seed)
	parser := parser.New(roller)

	// Accumulate all the args into a string.
	input := readArgs(os.Args[1:])

	// Now roll those dice!
	result, err := parser.Roll(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	fmt.Printf("%s\n", result.Describe())
}

func readArgs(args []string) string {
	var buf bytes.Buffer

	for i, arg := range args {
		if i != 0 {
			fmt.Fprint(&buf, " ")
		}
		fmt.Fprintf(&buf, "%s", arg)
	}

	return buf.String()
}
