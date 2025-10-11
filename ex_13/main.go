package main

import (
	"bufio"
	"ex-13/models"
	"ex-13/service"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fFlagPtr := flag.String("f", "", "fields")
	dFlagPtr := flag.String("d", "    ", "delimiter")
	sFlagPtr := flag.Bool("s", false, "separated")

	flag.Parse()

	flags := models.Flag{
		FFlag: *fFlagPtr,
		DFlag: *dFlagPtr,
		SFlag: *sFlagPtr,
	}

	finalColumns := takeColumns(flags.FFlag)

	data := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	service.Logic(data, finalColumns, flags)
}

func takeColumns(fFlag string) []int {
	if fFlag == "" {
		log.Fatal("cut: you must specify a list of fields")
	}
	columns := strings.Split(fFlag, ",")
	finalColumns := []int{}
	for _, val := range columns {
		if len(val) == 1 {
			iVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			finalColumns = append(finalColumns, iVal)
		} else {
			startIndex, err1 := strconv.Atoi(string(val[0]))
			endIndex, err2 := strconv.Atoi(string(val[2]))
			if err1 != nil || err2 != nil {
				log.Fatal("not integer")
			}
			for i := startIndex; i <= endIndex; i++ {
				finalColumns = append(finalColumns, i)
			}
		}
	}
	if finalColumns[0] < 1 {
		log.Fatal("cut: [-bcf] list: values may not include zero")
	}

	return finalColumns
}
