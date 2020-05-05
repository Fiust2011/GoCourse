package main

import "fmt"

type Automobile struct {
	CarBody           string
	Mark              string
	Year              int
	BodyVolume        int
	EngineRunning     string
	WindowsOpen       string
	HowFullBodyVolume string
}

func main() {

	auto1 := Automobile{
		CarBody:           "car",
		Mark:              "Lada",
		Year:              2020,
		BodyVolume:        200,
		EngineRunning:     "yes",
		WindowsOpen:       "yes",
		HowFullBodyVolume: "full",
	}

	auto2 := Automobile{
		CarBody:           "truck",
		Mark:              "Gaz",
		Year:              2005,
		BodyVolume:        1000,
		EngineRunning:     "no",
		WindowsOpen:       "no",
		HowFullBodyVolume: "empty",
	}

	fmt.Println(auto1)
	fmt.Println(auto2)
}
