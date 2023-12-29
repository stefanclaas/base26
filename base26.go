package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var decodeFlag = flag.Bool("d", false, "Decode input")
var lineLength = flag.Int("l", -1, "Number of characters per line")

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encode(b byte) string {
	return string(alphabet[b%26]) + string(alphabet[(b/26)%26])
}

func decode(s string) byte {
	return byte(indexOf(s[0]) + 26*indexOf(s[1]))
}

func indexOf(char byte) int {
	for i, c := range []byte(alphabet) {
		if c == char {
			return i
		}
	}
	return -1
}

func main() {
	flag.Parse()
	bytes, _ := ioutil.ReadAll(os.Stdin)
	if *decodeFlag {
		s := strings.ReplaceAll(string(bytes), "\n", "")
		for i := 0; i < len(s); i += 2 {
			fmt.Printf("%c", decode(s[i:i+2]))
		}
		fmt.Println()
	} else {
		var line string
		for _, b := range bytes {
			line += encode(b)
			if *lineLength != -1 && len(line) >= *lineLength {
				fmt.Println(line)
				line = ""
			}
		}
		if len(line) > 0 {
			fmt.Print(line)
		}
		if *lineLength != -1 || len(line) > 0 {
			fmt.Println()
		}
	}
}
