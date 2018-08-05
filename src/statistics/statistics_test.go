package statistics

import (
  "fmt"
  "testing"
)

func TestSum(t *testing.T) {
  nums := []int{1, 2, 3}
  sum := Sum(nums)
  if (sum != 6) {
    t.Error("expected sum to be 6")
  }
}

func ExampleSun() {
  nums := []int{1, 2, 3}
  fmt.Println("sum is", Sum(nums))
  // Output:
  // sum is 6
}
