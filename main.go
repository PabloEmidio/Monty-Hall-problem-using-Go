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


func swithDoor(playerChoice int, doorWasOpenned int) int{
	if (playerChoice == 2 && doorWasOpenned == 3) || (playerChoice == 3 && doorWasOpenned == 2){
		return 1
	} else if (playerChoice == 1 && doorWasOpenned == 3) || (playerChoice == 3 && doorWasOpenned == 1){
		return 2
	} else if (playerChoice == 1 && doorWasOpenned == 2) || (playerChoice == 2 && doorWasOpenned == 1){
		return 3
	} else {
		return 0
	}
}


func main(){
	rand.Seed(time.Now().UnixNano())

	var carPosition, doorWillBeOpenned, playerChoice, playerFinalChoice, playerWon int
	var winningsPercentage float32

	for i:=1; i < simulationTimes; i++ {
		carPosition = getRandomDoor()
		playerChoice = getRandomDoor()

		for {
			doorWillBeOpenned = getRandomDoor()

			if doorWillBeOpenned != carPosition && doorWillBeOpenned != playerChoice {
				break
			}
		}

		playerFinalChoice = swithDoor(playerChoice, doorWillBeOpenned)

		if playerFinalChoice == 0 {
			fmt.Println("An error occurred") // Indicate a bad combination between playerChoice and doorWillBeOpenned
			return
		}

		if playerFinalChoice == carPosition {
			playerWon++
		}

	}

	winningsPercentage = float32(playerWon)  / float32(simulationTimes) * 100.0

	fmt.Printf("%d matches done, switching the door it'd win %d times, %.1f of the matches.\n", simulationTimes, playerWon, winningsPercentage)

}
