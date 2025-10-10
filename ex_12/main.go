package main

import (
	"bufio"
	"context"
	"ex-12/service"
	"fmt"
	"log"
	"os"
	"os/signal"

	"strings"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal("error open file")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileInfo := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileInfo = append(fileInfo, scanner.Text())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		cancel()
	}()
	readCh := make(chan string, 1)

	go func() {
		defer close(readCh)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			readCh <- scanner.Text()
		}
	}()

	read(ctx, readCh, fileInfo)
}

func read(ctx context.Context, readCh chan string, fileInfo []string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("program end with ctrl+c")
			return
		case text, ok := <-readCh:
			if ok {
				validate(text, fileInfo)
			} else {
				return
			}
		}
	}
}

func validate(s string, fileInfo []string) {
	if len(s) == 0 {
		fmt.Println("no command")
		return
	}

	sSlice := strings.Split(s, " ")
	if sSlice[0] != "grep" {
		fmt.Printf("no command %s available\n", sSlice[0])
		return
	}

	if sSlice[1][0] != '-' {
		var exists bool
		for i, val := range fileInfo {
			if ok := strings.Contains(val, sSlice[1]); ok {
				fmt.Println(fileInfo[i])
				exists = true
			}
		}
		if !exists {
			fmt.Println("no substring in file")
		}
	} else {
		service.Logic(sSlice, fileInfo)
	}
}
