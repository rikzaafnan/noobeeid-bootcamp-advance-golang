package main

import "fmt"

func main() {

	var car = map[string]string{}
	car["name"] = "BWM"
	car["color"] = "Black"

	getDataString := getString(car)

	showString(getDataString)

}

func getString(car map[string]string) string {

	message := fmt.Sprintf("Mobil %s berwarna %s", car["name"], car["color"])
	return message
}

func showString(dataString string) {
	fmt.Println(dataString)
}
