package main

import (
	"fmt"
	"math"
)

func calcBmi() {
	var heightCentimeter, weightKilogram float64

	fmt.Print("YourHeight in centimeter : ")
	fmt.Scan(&heightCentimeter)

	fmt.Print("YourWeight in kilogram : ")
	fmt.Scan(&weightKilogram)

	centiToMeter := math.Round(heightCentimeter*100) / 10000

	bmi := weightKilogram / (math.Pow(centiToMeter, 2))

	fmt.Println("Your BMI is ", math.Round(bmi*100)/100)

	if bmi < 18.5 {
		fmt.Println("BMI Category is Underweight.")
	} else if 18.5 <= bmi && bmi < 25 {
		fmt.Println("BMI Category is Normal weight.")
	} else if 25 <= bmi && bmi < 30 {
		fmt.Println("BMI Category is Overweight.")
	} else if 30 <= bmi {
		fmt.Println("BMI Category is Obesity")
	}
}

func main() {
	calcBmi()
}
