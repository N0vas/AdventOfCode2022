package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

const inputFile string = "input.txt"

type Elf struct {
	id       int
	calories int
}

var elfs []Elf

func main() {
	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	elfs = make([]Elf, 0)
	var currentElf int = 1 //start with 1 since the task starts counting at 1
	var currentCalories int = 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			addElf(currentElf, currentCalories)

			currentElf++
			currentCalories = 0

			continue
		}
		intValue, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}

		currentCalories += intValue
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//catch last elf
	addElf(currentElf, currentCalories)

	//get elf number with most calories
	// Sort descending by calories preserving name order
	sort.SliceStable(elfs, func(i, j int) bool { return elfs[i].calories > elfs[j].calories })
	log.Println(elfs[0].calories)

}

func addElf(id int, calories int) {
	elfs = append(elfs, Elf{id, calories})
}
