package main

import (
	"flag"
	"fmt"
)

type flags struct {
	file string
}

var cmd_flags flags

func init() {
	flag.StringVar(&cmd_flags.file, "f", "", "A file to read the input from")
	flag.Parse()
}

func main() {
	lines, lineLimit := getText()

	spacing := ""
	line := ""
	for i := 0; i < lineLimit+4; i++ {
		spacing += " "
	}
	for i := 2; i < lineLimit+4; i++ {
		line += "_"
	}

	//print the speech bubble
	fmt.Printf("  %s  \n", line)
	fmt.Printf("/%s\\\n", spacing)

	for i := 0; i < len(lines); i++ {
		for j := len(lines[i]); j <= lineLimit; j++ {
			lines[i] += string(' ')
		}
		fmt.Printf("|  %s  |\n", lines[i])
	}

	fmt.Printf("\\%s__/\n", line)

	//print the amogus
	fmt.Printf("%s", amogus)

}
