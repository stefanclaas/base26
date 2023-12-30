package main

import (
	"bufio"
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
	if *decodeFlag {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		s := strings.ReplaceAll(string(bytes), "\n", "")
		var result []byte
		for i := 0; i < len(s); i += 2 {
			result = append(result, decode(s[i:i+2]))
		}
		fmt.Printf("%s\n", string(result))
	} else {
		reader := bufio.NewReader(os.Stdin)
		var line string
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				break
			}
			for _, b := range []byte(string(r)) {
				line += encode(b)
			}
		}
		if *lineLength != -1 {
			for i := 0; i < len(line); i += *lineLength {
				end := i + *lineLength
				if end > len(line) {
					end = len(line)
				}
				fmt.Println(line[i:end])
			}
		} else {
			fmt.Println(line)
		}
	}
}

