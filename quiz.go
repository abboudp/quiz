package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Question struct {
	q string
	a string
}

func main() {
	f, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	data, err := r.ReadAll()
	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	questions := make([]Question, 0)
	correct := 0

	for _, i := range data {
		var line Question
		line.q = i[0]
		line.a = i[1]
		questions = append(questions, line)
	}
	timer := time.NewTimer(5 * time.Second)
	go func() {
		<-timer.C
		fmt.Println(fmt.Sprintf("You got %d correct", correct))
		os.Exit(1)
	}()

	for _, i := range questions {
		fmt.Println("What is", i.q, "equal to?")
		var input string
		_, err := fmt.Scanln(&input)

		if err != nil {
			log.Fatal(err)
		}

		if input == i.a {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect...")
		}
	}
	fmt.Println(fmt.Sprintf("You got %d correct", correct))
}
