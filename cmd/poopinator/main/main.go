package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	pooToken = "💩"
	delim    = "🤢"
)

var isPoop = regexp.MustCompile(fmt.Sprintf("^[%v|%v]+$", pooToken, delim)).MatchString

func main() {
	if len(os.Args) == 2 && os.Args[1] == "help" {
		printHelp()
		return
	} else if len(os.Args) < 3 {
		fmt.Printf("Not enough arguments!\nUsage:\n")
		printHelp()
		return
	}

	action := os.Args[1]
	val := os.Args[2]

	if action != "-d" && action != "-e" && len(val) == 0 {
		panic("Wrong arguments")
	}

	if action == "-e" {
		fmt.Println(poopinate([]byte(val)))
	} else if action == "-d" {
		res, err := depoopinate(val)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return
		}
		fmt.Println(string(res))
	}
	return
}

func poopinate(bts []byte) string {
	var result []byte

	for _, bt := range bts {
		result = append(result, []byte(strings.Repeat(pooToken, int(bt))+delim)...)
	}
	return strings.TrimSuffix(string(result), delim)
}

func depoopinate(pileOfPoo string) ([]byte, error) {
	if !isPoop(pileOfPoo) {
		return nil, fmt.Errorf("Not a complete Poo!")
	}

	var result []byte
	for _, poo := range strings.Split(pileOfPoo, delim) {
		result = append(result, uint8(strings.Count(poo, pooToken)))
	}
	return result, nil
}

func printHelp() {
	fmt.Printf("-e abcd \t-> 💩💩💩\n-d 💩💩💩\t-> abcd")
}
