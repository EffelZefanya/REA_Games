package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Inputter struct{
	scanner *bufio.Reader
}

func NewInputter() *Inputter{
	return &Inputter{
		scanner: bufio.NewReader(os.Stdin),
	}
}

func (in *Inputter) ReadInput(prompt string) string {
	for {
		fmt.Print(prompt)

		input, err := in.scanner.ReadString('\n')
		if err != nil {
			fmt.Println("[Error] There's an error reading input:", err)
			continue
		}

		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == ""{
			fmt.Println("[Error] Please input a value")
			continue
		}

		return trimmedInput
	}
}

func (in *Inputter) ReadFloat(prompt string) float64 {
	for {
		input := in.ReadInput(prompt)

		floatInput, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("[Error] The given input isn't a float number, please try again.")
			continue
		}

		return floatInput
	}
}

func (in *Inputter) ReadInt(prompt string) int {
	for {
		input := in.ReadInput(prompt)

		intInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("[Error] The given input isn't an integer, please try again.")
			continue
		}

		return intInput
	}
}
