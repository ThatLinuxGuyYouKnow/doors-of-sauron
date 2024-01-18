package main
import (
	"fmt"
	"math/rand"
	"time"
)

var usersGuess int
var actualDoor int = rand.Intn(3)
var usersInitialGuessIsTheDoor bool

func generateRandomNumberThatIsntZero() int {
	rand.Seed(time.Now().UnixNano())
	actualSeedDoor := rand.Intn(3) + 1 // Adding 1 to avoid getting 0, as Intn(3) generates numbers in the range [0, 3)

	actualDoor = actualSeedDoor
	return actualDoor
}
func extractTheWildcardDoorNotInGuess(usersGuess int, correctDoor int) int {
	// Create a map to track which numbers are present
	presentNumbers := make(map[int]bool)

	// Populate the map with the given inputs
	presentNumbers[usersGuess] = true
	presentNumbers[correctDoor] = true

	// Find the missing number
	missingNumber := 0
	for i := 1; i <= 3; i++ {
		if !presentNumbers[i] {
			missingNumber = i
			break
		}
	}

	return missingNumber
}

func randomDoorPicker(usersGuess int, correctDoor int) int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 1
	randomNumber := rand.Intn(2)

	// Define your two options
	option1 := usersGuess
	option2 := correctDoor

	// Make the random choice and return it
	if randomNumber == 0 {
		return option1
	} else {
		return option2
	}
}

func secondGuessSequence(correctDoorGuess int, usersPreviousDoorGuess int) {
	println("Welcome to the final round player! congrats on getting this far!")
	println("In the previous round, you picked the door ", usersPreviousDoorGuess)
	var wildcardDoor = extractTheWildcardDoorNotInGuess(usersPreviousDoorGuess, correctDoorGuess)
	println("now player, i can assure you that the answer is *NOT* %d", wildcardDoor)
	println("now that you know this,")
	println("are you staying with your previous guess, or do wou want to change it")
	println("y for yes, n for no")
	var isUserChosingToChangeGuess string

	fmt.Scanln(&isUserChosingToChangeGuess)

	switch isUserChosingToChangeGuess {
	case "y":

		fmt.Println("alright then")
		fmt.Println("..........")
		fmt.Println("you'reeee.........")
		if correctDoorGuess == usersPreviousDoorGuess {
			println("Correct! you win!!!")
		}

	case "n":
		fmt.Println("bold choice!")
		fmt.Println("..........")
		fmt.Println("you'reeee.........")
		randomDoorPicker(usersPreviousDoorGuess, correctDoorGuess)
	default:
		println("I dont know what %d means, you quit? alright then ", isUserChosingToChangeGuess)
	}

}
func secondGuesssSequenceWhenOriginalDoorPickWasCorrect(userGuessWhichIsAlsoCorrectGuess int) {
	var missing1, missing2 int
	switch userGuessWhichIsAlsoCorrectGuess {
	case 1:
		missing1, missing2 = 2, 3
	case 2:
		missing1, missing2 = 1, 3
	case 3:
		missing1, missing2 = 1, 2
	}
	var wildcardDoor = randomDoorPicker(missing1, missing2)
	println("Welcome to the final round player! congrats on getting this far!")
	println("In the previous round, you picked the door ", userGuessWhichIsAlsoCorrectGuess)
	println("now player, i can assure you that the answer is *NOT* %d", wildcardDoor)
	println("now that you know this,")
	println("are you staying with your previous guess, or do wou want to change it")
	println("y for yes, n for no")
	var isUserChosingToChangeGuess string

	fmt.Scanln(&isUserChosingToChangeGuess)

	switch isUserChosingToChangeGuess {
	case "y":

		fmt.Println("alright then")
		fmt.Println("..........")
		fmt.Println("you'reeee.........")

	

	case "n":
		fmt.Println("bold choice!")
		fmt.Println("..........")
		fmt.Println("you'reeee.........")
		println("wrong!, you should have stuck with your lasy answer")
	default:
		println("I dont know what %d means, you quit? alright then ", isUserChosingToChangeGuess)
	}

}

func main() {
	generateRandomNumberThatIsntZero()
	println("Welcome player!")
	println(" Here are the rules")
	println("There are 3 doors, behind 1 door is the grandprize of a million dollars, the other 2? dissapointment(it's goats, you have goats behind the other two doors)")
	println("In this first round, you can pick 1 door, when you do this, the announcer will open another door for you, which definetly does *NOT* have the prize behind it ")
	println("Now pick your first door, 1,2 or 3")
	fmt.Scanln(&usersGuess)
	if usersGuess == actualDoor {
		usersInitialGuessIsTheDoor = true
		secondGuesssSequenceWhenOriginalDoorPickWasCorrect(usersGuess)
	} else {
		usersInitialGuessIsTheDoor = false
		secondGuessSequence(actualDoor, usersGuess)

	}
	// user has technically picked wrongly
	// when user picks one door, 'open' 1 other door -comment logic lane 1,

}
