// package declaration for a Go program
package main

// syntax for importing multiple packages
import (
	"fmt"
	"math"
)

// function declaration
func main() {
	// variable declarations
  var stationName string
  var nextTrainTime int8
  var isUptownTrain bool
  
	// variable assignments
  stationName = "Union Square"
  nextTrainTime = 12
  isUptownTrain = false
  
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
  
  stationName = "Grand Central"
  nextTrainTime = 3
  isUptownTrain = true
  
  fmt.Println("")
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
}