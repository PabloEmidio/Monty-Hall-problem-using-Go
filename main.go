package main

import (
	"fmt"
	"math/rand"
	"time"
)


const simulationTimes = 500
const problemTotalDoors = 3


func getRandomDoor() int {
	return rand.Intn(problemTotalDoors) + 1
}


func switchDoor(playerChoice *int, doorOpenned int){
	if (*playerChoice == 2 && doorOpenned == 3) || (*playerChoice == 3 && doorOpenned == 2){
		*playerChoice = 1
	} else if (*playerChoice == 1 && doorOpenned == 3) || (*playerChoice == 3 && doorOpenned == 1){
		*playerChoice = 2
	} else if (*playerChoice == 1 && doorOpenned == 2) || (*playerChoice == 2 && doorOpenned == 1){
		*playerChoice = 3
	} else {
		*playerChoice = 0
	}
}


func main(){
	rand.Seed(time.Now().UnixNano())

	var carPosition, doorOpenned, playerChoice, playerWon int
	var winningsPercentage float32

	for i:=1; i < simulationTimes; i++ {
		carPosition = getRandomDoor()
		playerChoice = getRandomDoor()

		for {
			doorOpenned = getRandomDoor()

			if doorOpenned != carPosition && doorOpenned != playerChoice {
				break
			}
		}

		switchDoor(&playerChoice, doorOpenned)

		if playerChoice == 0 {
			fmt.Println("An error occurred") // Indicate a bad combination between playerChoice and doorOpenned
			return
		}

		if playerChoice == carPosition {
			playerWon++
		}

	}

	winningsPercentage = float32(playerWon)  / float32(simulationTimes) * 100.0

	fmt.Printf("%d matches done, switching the door it'd win %d times, %.1f of the matches.\n", simulationTimes, playerWon, winningsPercentage)

}
