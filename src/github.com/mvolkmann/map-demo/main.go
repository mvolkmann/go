package main

import "fmt"

func main() {
	type playerScoreMap map[string]int

	scoreMap := playerScoreMap{"Mark": 90, "Tami": 92}
	scoreMap["Amanda"] = 83
	scoreMap["Jeremy"] = 95

	myScore := scoreMap["Mark"]
	fmt.Println("my score =", myScore)
	fmt.Printf("scoreMap = %+v\n", scoreMap)
}
