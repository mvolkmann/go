package main

func main() {
	//var myMap = map[string]int{"alpha": 1, "beta": 2}
	myMap := map[string]int{"alpha": 1, "beta": 2}
	for key, value := range myMap {
		println(key, "=", value)
	}
}