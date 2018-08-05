package main
import "statistics"

func main() {
	numbers := []int{1, 3, 7, 24}
	sum, avg := statistics.GetStats(numbers)
	println("sum =", sum)
	println("avg =", avg)
}