package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	sum, err := calcular()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}

func calcular() (int, error) {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)

	var sum int = 0

	for scanner.Scan() {
		line := scanner.Text()
		var numbers string
		for _, num := range line {
			numbers += re.FindString(string(num))
		}

		first := string(numbers[0])
		last := string(numbers[len(numbers)-1])
		num, err := strconv.Atoi(first + last)
		if err != nil {
			return -1, errors.New("erro ao converter a string em número inteiro")
		}

		sum += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum, nil
}
