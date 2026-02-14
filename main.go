package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	pathFlag := flag.String("p", "problems.csv", "file path")
	timeFlag := flag.Int("t", 2, "time duration in seconds")
	shuffleFlag := flag.Bool("s", false, "is shuffle")

	flag.Parse()

	var rightAnswersAmount int

	file, err := os.Open(*pathFlag)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	if *shuffleFlag {
		records = shuffle(records)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("start? y/n ")
	s, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	if strings.TrimSpace(s) != "y" {
		os.Exit(1)
	}

	for _, row := range records {
		timer := time.AfterFunc(time.Duration(*timeFlag)*time.Second, func() {
			fmt.Println("\nno time left!")
			printResults(rightAnswersAmount, len(records))

			os.Exit(1)
		})

		question := row[0]
		answer := row[1]

		fmt.Printf("%s = ", question)

		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		timer.Stop()

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if answer == input {
			rightAnswersAmount += 1
		}
	}

	printResults(rightAnswersAmount, len(records))
}

func printResults(answers, total int) {
	fmt.Printf("right: %d\nwrong: %d\ntotal:%d\n", answers, total-answers, total)
}

func shuffle(arr [][]string) [][]string {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
