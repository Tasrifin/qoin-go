package app

import (
	"fmt"
	"log"

	"github.com/Tasrifin/qoin-go/helpers"
)

func StartAPP() {
	var totalPlayer, totalDice int

	//Input Total Players
	for {
		fmt.Print("Masukkan total player : ")

		data, err := helpers.CheckInput()
		if err != nil {
			log.Println("Error : ", err)
			continue
		}

		totalPlayer = data
		break
	}

	//Input Total Cube
	for {
		fmt.Print("Masukkan total dadu : ")

		data, err := helpers.CheckInput()
		if err != nil {
			log.Println("Error : ", err)
			continue
		}

		totalDice = data
		break
	}

	//Play Game
	PlayGame(totalPlayer, totalDice)
}

func PlayGame(totalPlayer, totalDice int) {
	countGame := 0
	gameDatas := make(map[int][]int)
	gameRules := make(map[int]int)
	gameScore := make(map[int]int)

	//Set Rules
	for i := 1; i <= totalPlayer; i++ {
		gameRules[i] = totalDice
		gameScore[i] = 0
	}

	for {
		countGame++
		fmt.Println("\n========================================")
		fmt.Printf("\nGiliran %d Lempar Dadu : \n", countGame)

		for i := 1; i <= totalPlayer; i++ {
			limitDice := gameRules[i]

			for j := 1; j <= limitDice; j++ {
				getRandomNumber := helpers.GenerateRandomNumber()
				gameDatas[i] = append(gameDatas[i], getRandomNumber)
			}

			fmt.Printf("    Pemain #%v (%v) : %v \n", i, gameScore[i], gameDatas[i])
		}

		//Validate
		availablePlayers := 0
		gameDatas, gameScore, gameRules = ValidateResult(gameDatas, gameScore, gameRules, totalPlayer)

		fmt.Println("\nSetelah Evaluasi : ")
		for i := 1; i <= totalPlayer; i++ {
			fmt.Printf("    Pemain #%v (%v) : %v \n", i, gameScore[i], gameDatas[i])

			//Check
			lenData := len(gameDatas[i])
			if lenData > 0 {
				availablePlayers++
			}
		}

		//Reset Game Data
		gameDatas = make(map[int][]int)

		if availablePlayers <= 1 {
			break
		} else {
			continue
		}
	}
}

func ValidateResult(gameDatas map[int][]int, gameScore map[int]int, gameRules map[int]int, totalPlayer int) (finalData map[int][]int, finalScore map[int]int, finalRules map[int]int) {
	for i := 1; i <= totalPlayer; i++ {
		//Set Next Player
		nextPlayer := i + 1
		if i == totalPlayer {
			nextPlayer = 1
		}

		data := gameDatas[i]
		for _, v := range data {
			if v == 6 || v == 1 {
				//Renew Data
				gameDatas[i] = ReNewData(gameDatas[i])
				if v == 1 {
					//Give to Next Player
					if len(gameDatas[nextPlayer]) > 0 {
						gameDatas[nextPlayer] = append(gameDatas[nextPlayer], v)
					}
				} else {
					//Add Score
					gameScore[i]++
				}
			}
		}

		//Renew Rule
		gameRules[i] = len(gameDatas[i])
		if i == totalPlayer {
			gameRules[1] = len(gameDatas[1])
		}
	}

	return gameDatas, gameScore, gameRules
}

func ReNewData(data []int) (result []int) {
	for _, v := range data {
		if v != 6 && v != 1 {
			result = append(result, v)
		}
	}

	return
}
