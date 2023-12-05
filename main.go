package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	red   = "red"
	green = "green"
	blue  = "blue"
)

var gameConfiguration = map[string]int{
	red:   12,
	green: 13,
	blue:  14,
}

func main() {
	file, err := os.Open("cube_games_outcomes.txt")
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		unmodifiedGameOutcome := scanner.Text()
		// extract the game number
		splitGame := strings.Split(unmodifiedGameOutcome, ":")
		re := regexp.MustCompile("[0-9]+")
		gameNumberStr := re.FindString(splitGame[0])
		gameNumber, err := strconv.Atoi(gameNumberStr)
		if err != nil {
			log.Fatal(err)
		}

		// discard the "Game n: " prefix
		re2 := regexp.MustCompile("Game [0-9]+: ")
		split := re2.Split(unmodifiedGameOutcome, -1)
		withoutGamePrefix := split[1]
		commaSeparatedOutcomes := strings.ReplaceAll(withoutGamePrefix, ";", ",")

		asArray := strings.Split(commaSeparatedOutcomes, ",")

		gameIsPossible := true
		for index := range asArray {
			if !gameIsPossible {
				break
			}

			cubeCountAndColorArray := strings.Split(strings.TrimSpace(asArray[index]), " ")
			countStr := cubeCountAndColorArray[0]
			count, err := strconv.Atoi(countStr)
			if err != nil {
				log.Fatal(err)
			}
			color := cubeCountAndColorArray[1]

			if count > gameConfiguration[color] {
				gameIsPossible = false
				break
			}
		}

		if gameIsPossible {
			sum += gameNumber
		}
	}

	log.Printf("Done; sum of possible game numbers: %d\n", sum)
}
