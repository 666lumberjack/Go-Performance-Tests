package main

import "os"

func main() {
	mode := os.Args[1]
	switch mode {
	case "minVsMulti":
		minVsMultiCompare()
	}
}
