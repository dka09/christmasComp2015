package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	jan := findPossibilities(31)
	apr := findPossibilities(30)
	may := findPossibilities(31)

	aMap := createMapForA(jan[:], apr[:], may[:])

	for _, v := range aMap {
		if check(v[0], v[1], v[2]) {
			break
		}
	}

	endTime := time.Since(startTime)
	fmt.Println(endTime.Nanoseconds() / 1000)
}

// Find all possible numbers divisible by the num provided that are 3 digits in length.
// Function takes an int for the num and returns a slice.

func findPossibilities(num int) []int {

	possibilities := []int{}
	var possibility int
	var counter int

	for possibility < 1000 {
		possibility = counter * num

		if possibility < 1000 && possibility > 100 {
			if repeats(possibility) && removeZero(possibility, num) {
				possibilities = append(possibilities, possibility)
			}
		}

		counter++
	}

	return possibilities

}

// Checks to see if the int provided contains repeated digits.
// Returns true if the int does not contain repeated digits else false.

func repeats(x int) bool {
	var counts [10]int
	for x > 0 {
		counts[x%10]++
		if counts[x%10] > 1 {
			return false
		}
		x /= 10
	}
	return true
}

// Checks to see if the provided int contains a 0 unless it is a number for APR.
// The digit 0 must be in APR as it is the only month divisible by 30.
// Returns true if it does not contain a 0 else false.

func removeZero(x int, num int) bool {
	if num == 30 {
		return true
	}

	for x > 0 {
		if x%10 == 0 {
			return false
		}
		x /= 10
	}
	return true
}

// Create a map and return the map.
// KEY 	= possible numbers to represent letters.
// VALUE = A slice containing 3 other slices for possible numbers for JAN, APR, MAY.

func createMapForA(jan []int, apr []int, may []int) map[int][][]int {
	//for each number in array, check a's that match over months
	aMap := map[int][][]int{
		1: [][]int{},
		2: [][]int{},
		3: [][]int{},
		4: [][]int{},
		5: [][]int{},
		6: [][]int{},
		7: [][]int{},
		8: [][]int{},
		9: [][]int{},
	}

	// Add the three slices to the value of each key of the map.

	for k, _ := range aMap {
		aMap[k] = append(aMap[k], []int{})
		aMap[k] = append(aMap[k], []int{})
		aMap[k] = append(aMap[k], []int{})
	}

	// For each monnth JAN, APR, MAY, loop through each possible number
	// and add it to the relevant map key depending on the number representing A.

	for _, v := range jan {
		val := strconv.Itoa(v)
		valInt, _ := strconv.Atoi(string(val[1]))
		aMap[valInt][0] = append(aMap[valInt][0], v)
	}

	for _, v := range apr {
		val := strconv.Itoa(v)
		valInt, _ := strconv.Atoi(string(val[0]))
		aMap[valInt][1] = append(aMap[valInt][1], v)
	}

	for _, v := range may {
		val := strconv.Itoa(v)
		valInt, _ := strconv.Atoi(string(val[1]))
		aMap[valInt][2] = append(aMap[valInt][2], v)
	}

	return aMap
}

// Loop through all possible combinations provided as slices for each month.
// (values taken from map created in createMapForA)

func check(jan []int, apr []int, may []int) bool {
	allPossibleSolutions := []string{}

	for i := 0; i < len(jan); i++ {
		for j := 0; j < len(apr); j++ {
			for k := 0; k < len(may); k++ {
				if compareAllMonths(jan[i], apr[j], may[k]) {

					var buffer bytes.Buffer
					buffer.WriteString(strconv.Itoa(jan[i]))
					buffer.WriteString(strconv.Itoa(apr[j]))
					buffer.WriteString(strconv.Itoa(may[k]))

					if findFeb(buffer.String()) {
						// fmt.Println("jan",jan[i])
						// fmt.Println("apr",apr[j])
						// fmt.Println("may",may[k])
						return true
					}

					allPossibleSolutions = append(allPossibleSolutions, buffer.String())
				}
			}
		}
	}
	return false
}

// Compare all combinations of number for the months JAN, APR, MAY

func compareAllMonths(j int, a int, m int) bool {
	if !compareTwoMonths(j, m, "jan", "may") {
		return false
	}
	if !compareTwoMonths(j, a, "jan", "apr") {
		return false
	}
	if !compareTwoMonths(a, m, "apr", "may") {
		return false
	}

	return true
}

// Compare two months and check if any letters are present in both.
// Check if letter in secondMonth is present in firstMonth, if not, check if number in secondNum is in firstNum

func compareTwoMonths(first int, second int, firstMonth string, secondMonth string) bool {
	firstNum, secondNum := strconv.Itoa(first), strconv.Itoa(second)

	m := map[string]string{
		string(firstMonth[0]): string(firstNum[0]),
		string(firstMonth[1]): string(firstNum[1]),
		string(firstMonth[2]): string(firstNum[2]),
	}

	for index, value := range secondMonth {
		va := string(value)
		if _, ok := m[va]; !ok {
			for _, v := range m {
				if v == string(secondNum[index]) {
					return false
				}
			}
		}
	}

	return true
}

// Find the numbers left for FEB from the possible combination found for JAN, APR, MAY

func findFeb(array string) bool {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	possibleNumbersForFeb := []string{}
	var buffer bytes.Buffer

	for _, v := range numbers {
		if strings.Index(array, v) < 0 {
			buffer.WriteString(v)
		}
	}

	possibleNumbersForFeb = append(possibleNumbersForFeb, buffer.String())

	return checkFebNumbers(possibleNumbersForFeb)
}

// Check all combinations of the numbers for FEB

func checkFebNumbers(array []string) bool {
	for _, value := range array {

		one, two, three := string(value[0]), string(value[1]), string(value[2])

		if combinations(one, two, three) {
			return true
		}
		if combinations(one, three, two) {
			return true
		}
		if combinations(two, one, three) {
			return true
		}
		if combinations(two, three, one) {
			return true
		}
		if combinations(three, one, two) {
			return true
		}
		if combinations(three, two, one) {
			return true
		}

	}

	return false
}

// Check if the certain combination of numbers is divisible by 28

func combinations(a string, b string, c string) bool {
	var num int

	var buffer bytes.Buffer

	buffer.WriteString(a)
	buffer.WriteString(b)
	buffer.WriteString(c)

	num, _ = strconv.Atoi(buffer.String())
	if num%28 == 0 {
		fmt.Println(num)
		return true
	}
	return false
}
