package main

import (
	"fmt"
	"math/rand"
)

type diceProfile [6]int
type actionDice struct {
	diceProfile diceProfile
	results     [100]int
	total       int
	ammo        int
}

func main() {
	myResults := actionDice{
		total:       100000,
		ammo:        3,
		diceProfile: diceProfile{0, 0, 0, 0, 0, 1},
	}

	for index := 0; index < myResults.total; index++ {
		fail := 0
		count := 0
		for index := 0; fail < myResults.ammo; index++ {
			randNum := rand.Intn(6)
			count++
			checkSuccess := myResults.diceProfile[randNum]
			if checkSuccess == 1 {
				fail++
			}
		}
		myResults.results[count]++
	}

	for i, num := range myResults.results {
		fmt.Println(i, num)
	}
}
