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

	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
	// part 2
	/*
		file, err := os.Open("cube_games_outcomes.txt")
		if err != nil {
			log.Fatalln("Error opening file: ", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		gamesMinCubesByColorMap := make(map[int]map[string]int, 100) // we know how many games are in the outcomes text file, so we can pre-allocate the capacity of the map to help the compiler
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

			// initialize the inner map in gamesMinCubesByColorMap
			gamesMinCubesByColorMap[gameNumber] = make(map[string]int, 3) // glancing at the outcomes text file, we can make a guestimation of 3 to help the compiler pre-allocate

			// discard the "Game n: " prefix
			re2 := regexp.MustCompile("Game [0-9]+: ")
			split := re2.Split(unmodifiedGameOutcome, -1)
			withoutGamePrefix := split[1]
			commaSeparatedOutcomes := strings.ReplaceAll(withoutGamePrefix, ";", ",")

			asArray := strings.Split(commaSeparatedOutcomes, ",")

			//highestCubeCount := 0
			for index := range asArray {
				cubeCountAndColorArray := strings.Split(strings.TrimSpace(asArray[index]), " ")
				countStr := cubeCountAndColorArray[0]
				count, err := strconv.Atoi(countStr)
				if err != nil {
					log.Fatal(err)
				}
				color := cubeCountAndColorArray[1]

				if gamesMinCubesByColorMap[gameNumber][color] < count {
					gamesMinCubesByColorMap[gameNumber][color] = count
				}
			}
		}

		sum := 0
		for _, outcomes := range gamesMinCubesByColorMap {
			//fmt.Println(outcomes)
			val := 1
			for _, count := range outcomes {
				val = val * count
			}

			sum += val
		}

		log.Printf("Done; sum of power of sets: %d\n", sum)
	*/
}
