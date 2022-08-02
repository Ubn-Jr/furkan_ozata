// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	// lose weight or get weight
	var focus string = "lose weight"
	var gender string = "female"
	var weight int = 80
	var lengthCM int = 183
	var age int = 30

	if gender == "male" && focus == "lose weight" {
		getMaleCalorie := maleCalorieCalculator(weight, lengthCM, age)
		fmt.Println(getMaleCalorie)
	} else if gender == "male" && focus == "get weight" {
		getMaleCalorie := maleCalorieCalculator(weight, lengthCM, age) + 1000
		fmt.Println(getMaleCalorie)
	} else if gender == "female" && focus == "lose weight" {
		getFemalaCalorie := femaleCalorieCalculator(weight, lengthCM, age)
		fmt.Println(getFemalaCalorie)
	} else if gender == "female" && focus == "get weight" {
		getFemalaCalorie := femaleCalorieCalculator(weight, lengthCM, age) + 1000
		fmt.Println(getFemalaCalorie)
	} else {
		fmt.Println("eksik veya hatali giris yaptiniz")
	}
}

func maleCalorieCalculator(weight int, lenghtCM int, age int) int {
	maleCalorie := (665 + (10 * weight) + (2 * lenghtCM) - (5 * age))
	return maleCalorie

}
func femaleCalorieCalculator(weight int, lengthCM int, age int) int {
	famaleCalorie := (66 + (10 * weight) + (2 * lengthCM) - (5 * age))
	return famaleCalorie
}
