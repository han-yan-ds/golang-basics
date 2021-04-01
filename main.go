// package declaration for a Go program
package main

// syntax for importing multiple packages
import (
	"fmt"
	m "math" // alias m
)

// function declaration
func main() {
	// variable declarations
  const stationName string = "Union Square"
  var nextTrainTime int8
  var isUptownTrain bool
  
	// variable assignments
  nextTrainTime = 12
  isUptownTrain = false
  
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
  
  nextTrainTime = 3
  isUptownTrain = true
  
  fmt.Println("")
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
}