package main

import (
	"fmt"
	"strconv"
)

func main() {
	stringSlice := []string{"4#hello", "is", "everyone", "Hearing", "Me?"}
	fmt.Println(encode(stringSlice))
	fmt.Println(decode(encode(stringSlice)))

}

func encode(strs []string) string {
	retString := ""
	for _, str := range strs {
		retString += fmt.Sprint(len(str))
		retString += "#"
		retString += str
	}

	return retString

}

func decode(str string) []string {
	i := 0

	retArr := []string{}
	numToStringStore := ""
	for i < len(str) {
		for str[i] != '#' {
			numToStringStore += string(str[i])
			i++
		}
		num, err := strconv.Atoi(numToStringStore)
		if err != nil {
			panic(err)
		}
		fmt.Println(i)
		retArr = append(retArr, str[i+1:i+num+1])
		i = i + num + 1
		numToStringStore = ""
	}

	return retArr
}
