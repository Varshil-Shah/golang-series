package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza shop!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How many pizza(s) would you like to have?")

	input, _ := reader.ReadString('\n')
	fmt.Println("Number of pizza to be ordered, ", input)

	count, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Here is you total pizza count with 1 pizza free", count+1)
	}
}
