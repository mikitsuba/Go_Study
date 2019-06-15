package main

import "fmt"

func main() {
	favSport := "badminton"
	switch favSport {
	case "snowboarding":
		fmt.Println("go to the mountains!")
	case "diving":
		fmt.Println("lets go together to koh Tao!")
	case "baseball":
		fmt.Println("which team is your favorite?")
	case "soccer":
		fmt.Println("Project BJ!")
	default:
		fmt.Println("you and I dont have anything in commmon.")
	}
}
