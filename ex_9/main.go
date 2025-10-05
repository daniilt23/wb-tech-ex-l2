package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	answer, err := CheckCondition(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

func CheckCondition(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	sSplit := strings.Split(s, "")
	count := 0
	for _, val := range sSplit {
		if isDigit(val) {
			count++
		}
	}

	if count == len(s) {
		return "", errors.New("строка не может состоять только из цифр")
	}

	return strings.Join(stringUnpacking(sSplit), ""), nil
}

func stringUnpacking(sSplit []string) []string {
	answer := []string{}
	answer = append(answer, sSplit[0])
	for i := 1; i < len(sSplit); i++ {
		if sSplit[i] == "\\" { // доп условие на escape-последовательности
			if (i+2 < len(sSplit) && sSplit[i+2] == "\\") || (i+2 == len(sSplit)) {
				answer = append(answer, sSplit[i+1])
			}
			i++
			continue
		}
		if isDigit(sSplit[i]) {
			num, _ := strconv.Atoi(sSplit[i])
			if isDigit(sSplit[i-1]) {
				num++
			}
			pushSlice := strings.Repeat(sSplit[i-1], num-1)
			answer = append(answer, pushSlice)
		} else {
			answer = append(answer, sSplit[i])
		}
	}

	return answer
}

func isDigit(s string) bool {
	return s >= "0" && s <= "9"
}
