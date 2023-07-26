package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var isValidEmojis func(s string) bool

func init() {
	emojis := make([]string, 0, 16)
	for i := 0; i < 16; i++ {
		emojis = append(emojis, string(emgen(uint8(i))))
	}
	isValidEmojis = regexp.MustCompile(fmt.Sprintf("^[%v]+$", strings.Join(emojis, "|"))).MatchString
}

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
		fmt.Println(encode([]byte(val)))
	} else if action == "-d" {
		res, err := decode(val)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return
		}
		fmt.Println(string(res))
	}
	return
}

func encode(bts []byte) string {
	result := make([]byte, 0, len(bts)*8)
	for _, b := range bts {
		result = append(result, emgen(b>>4)...)
		result = append(result, emgen(b&0x0f)...)
	}

	return string(result)
}

func decode(str string) ([]byte, error) {
	bts := []byte(str)
	if !isValidEmojis(str) || len(bts)%8 != 0 {
		return nil, fmt.Errorf("Not a valid emojis!")
	}
	result := make([]byte, 0, len(bts)/8)
	for i := 3; i < len(bts); i += 8 {
		v := ((bts[i] - 129) << 4) | (bts[i+4] - 129)
		result = append(result, v)
	}
	return result, nil
}

func emgen(b4 uint8) []byte {
	return []byte{240, 159, 152, b4 + 129}
}

func printHelp() {
	fmt.Printf("-e abcd \t-> 😅👉👌\n-d 😅👉👌\t-> abcd")
}
