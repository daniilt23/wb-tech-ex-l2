package service

import (
	"ex-13/models"
	"fmt"
	"strings"
)

func Logic(data []string, finalColumns []int, flags models.Flag) {
	for _, val := range data {
		var sSlice []string
		var delimiter string
		if flags.DFlag != "" {
			delimiter = flags.DFlag
			sSlice = strings.Split(val, delimiter)
		} else {
			delimiter = "    "
			sSlice = strings.Split(val, delimiter)
		}
		if flags.SFlag {
			if len(sSlice) == 1 {
				continue
			}
		}
		answer := []string{}
		for _, value := range finalColumns {
			if value <= len(sSlice) {
				answer = append(answer, sSlice[value-1])
			}
		}
		fmt.Println(strings.Join(answer, delimiter))
	}
}
