package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Inputter struct {
	scanner *bufio.Reader
}

func NewInputter() *Inputter {
	return &Inputter{
		scanner: bufio.NewReader(os.Stdin),
	}
}

func (in *Inputter) ReadInputOldVal(prompt string, oldval string) string {
	for {
		fmt.Printf("%s\n", prompt)
		fmt.Printf("%s\r", oldval)
		input, err := in.scanner.ReadString('\n')
		if err != nil {
			fmt.Println("[Error] There's an error reading input:", err)
			continue
		}

		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "" {
			return oldval
		}

		return trimmedInput
	}
}

func (in *Inputter) ReadFloatOldVal(prompt string, oldval float64) float64 {
	for {
		fmt.Printf("%s\n", prompt)
		fmt.Printf("%f\r", oldval)
		input, err := in.scanner.ReadString('\n')
		if err != nil {
			fmt.Println("[Error] There's an error reading input:", err)
			continue
		}
		fmt.Println(input)
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "" {
			return oldval
		}

		floatInput, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("[Error] The given input isn't a float number, please try again.")
			continue
		}

		return floatInput
	}
}

func (in *Inputter) ReadIntOldVal(prompt string, oldval int) int {
	for {
		fmt.Printf("%s\n", prompt)
		fmt.Printf("%d\r", oldval)
		input, err := in.scanner.ReadString('\n')
		if err != nil {
			fmt.Println("[Error] There's an error reading input:", err)
			continue
		}

		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "" {
			return oldval
		}

		intInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("[Error] The given input isn't an integer, please try again.")
			continue
		}

		return intInput
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
		if trimmedInput == "" {
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
