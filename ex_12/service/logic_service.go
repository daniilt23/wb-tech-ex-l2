package service

import (
	"fmt"
	"strconv"
	"strings"
)

func Logic(sSlice []string, fileInfo []string) {
	flag := sSlice[1][1]
	switch flag {
	case 'A':
		A(sSlice, fileInfo)
	case 'B':
		B(sSlice, fileInfo)
	case 'C':
		C(sSlice, fileInfo)
	case 'c':
		c(sSlice, fileInfo)
	case 'i':
		i(sSlice, fileInfo)
	case 'v':
		v(sSlice, fileInfo)
	case 'F':
		F(sSlice, fileInfo)
	case 'n':
		n(sSlice, fileInfo)
	default:
		fmt.Println("no flag:", string(sSlice[1][1]))
	}
}

func A(sSlice []string, fileInfo []string) {
	n, err := strconv.Atoi(sSlice[2])
	if err != nil {
		fmt.Println("N need to be int")
		return
	}
	for i, val := range fileInfo {
		if ok := strings.Contains(val, sSlice[3]); ok {
			endIndex := min(len(fileInfo), i+n)
			for j := i; j <= endIndex; j++ {
				fmt.Println(fileInfo[j])
			}
			break
		}
	}
}

func B(sSlice []string, fileInfo []string) {
	n, err := strconv.Atoi(sSlice[2])
	if err != nil {
		fmt.Println("N need to be int")
		return
	}
	for i, val := range fileInfo {
		if ok := strings.Contains(val, sSlice[3]); ok {
			startIndex := max(i-n, 0)
			for j := startIndex; j <= i; j++ {
				fmt.Println(fileInfo[j])
			}
			break
		}
	}
}

func C(sSlice []string, fileInfo []string) {
	n, err := strconv.Atoi(sSlice[2])
	if err != nil {
		fmt.Println("N need to be int")
		return
	}
	for i, val := range fileInfo {
		if ok := strings.Contains(val, sSlice[3]); ok {
			startIndex := max(i-n, 0)
			endIndex := min(len(fileInfo), i+n)
			for j := startIndex; j <= endIndex; j++ {
				fmt.Println(fileInfo[j])
			}
			break
		}
	}
}

func c(sSlice []string, fileInfo []string) {
	if len := len(sSlice); len < 3 {
		fmt.Println("there are no substring to find")
		return
	}
	count := 0
	for _, val := range fileInfo {
		if ok := strings.Contains(val, sSlice[2]); ok {
			count++
		}
	}
	fmt.Println(count)
}

func i(sSlice []string, fileInfo []string) {
	if len := len(sSlice); len < 3 {
		fmt.Println("there are no substring to find")
		return
	}

	for _, val := range fileInfo {
		if ok := strings.Contains(strings.ToLower(val), strings.ToLower(sSlice[2])); ok {
			fmt.Println(val)
		}
	}
}

func v(sSlice []string, fileInfo []string) {
	if len := len(sSlice); len < 3 {
		fmt.Println("there are no substring to find")
		return
	}

	for _, val := range fileInfo {
		if ok := strings.Contains(val, sSlice[2]); !ok {
			fmt.Println(val)
		}
	}
}

func F(sSlice []string, fileInfo []string) {
	if len := len(sSlice); len < 3 {
		fmt.Println("there are no substring to find")
		return
	}

	for _, val := range fileInfo {
		if val == sSlice[2] {
			fmt.Println(val)
		}
	}
}

func n(sSlice []string, fileInfo []string) {
	if len := len(sSlice); len < 3 {
		fmt.Println("there are no substring to find")
		return
	}

	for i, val := range fileInfo {
		if ok := strings.Contains(val, sSlice[2]); ok {
			fmt.Printf("%d. %s\n", i, val)
		}
	}
}
