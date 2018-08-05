package statistics

func Average(numbers []int) float32 {
  return float32(Sum(numbers)) / float32(len(numbers))
}

func Sum(numbers []int) int {
  sum := 0
  for _, number := range numbers {
    sum += number
  }
  return sum
}

func GetStats(numbers []int) (int, float32) {
  return Sum(numbers), Average(numbers)
}

