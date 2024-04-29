package main

import "fmt"

func getTestData() string {
	return "111111"
}
func getDate() string {
	return "2024/4/28"
}
func main() {

	fmt.Println("data============", getTestData(),
		getDate())
}
