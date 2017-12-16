package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/elberth90/advent_of_code_2017/day_13/firewall"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(byteData)
	data = strings.Trim(data, "\n")
	result := firewall.CaughtHowManyTimes(data)
	fmt.Printf("Result: `%d`\n", result)
	result = firewall.DonNotGetCaught(data)
	fmt.Printf("Result: `%d`\n", result)

}
