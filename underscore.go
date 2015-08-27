package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  // sum
  someNums := []float64{1, 2, 3, 4, 5, 6}
  fmt.Println(average(someNums))

  // max
  maxNums := []float64{12, 9, 2, 13, 56, 3, 2}
  fmt.Println(max(maxNums))

  // min
  minNums := []float64{12, 9, 2, 13, 56, 3, 2}
  fmt.Println(min(minNums))

  // contains
  someInts := []int{1, 2, 3, 4, 5}
  fmt.Println(contains(someInts, 2))
  fmt.Println(contains(someInts, 6))

  //shuffle
  shuffled := shuffle(someInts)
  fmt.Printf("array: %v\n", someInts)
  fmt.Printf("shuffled: %v\n", shuffled)

}

func average(nums []float64) float64{
  total := 0.0

  for i := 0; i < len(nums); i++ {
    total += nums[i]
  }

  return float64(total) / float64(len(nums))
}

func max(nums []float64) float64{
  max := nums[0]

  for i := 1; i < len(nums); i++ {
    if (nums[i] > max) {
      max = nums[i]
    }
  }

  return max
}

func min(nums []float64) float64{
  min := nums[0]

  for i := 1; i < len(nums); i++ {
    if (nums[i] < min) {
      min = nums[i]
    }
  }

  return min
}

func contains(s []int, e int) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}

func shuffle(nums []int) []int{
  shuffled := make([]int, len(nums))
  used := make([]int, len(nums))
  fmt.Printf("%v\n", used)

  // loop for length of input
  for i:= 0; i < len(nums); i++ {
    // get a random index in the slice
    rand.Seed(time.Now().UnixNano())
    ind := rand.Intn(len(nums))
    // if in used, repick
    for contains(used, ind) {
      ind = rand.Intn(len(nums))
    }
    // add index to used
    used[i] = ind
    // add value at index to shuffled
    shuffled[i] = nums[ind]
  }

  return shuffled
}
