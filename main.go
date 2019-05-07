package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Unable to open file %s: %s\n", path, err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			lines = append(lines, scanner.Text())
		}
	}
	return lines
}

func main() {
	sleep := flag.Uint("sleep", 50, "number of microseconds to sleep between hashes")
	conc := flag.Uint("threads", 2, "amount of concurrent goroutines that should run at the same time")
	file := flag.String("file", "", "input file")

	flag.Parse()

	if len(*file) == 0 {
		fmt.Println("You must specify a file to open")
		os.Exit(1)
	}

	targets = readFile(*file)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Total run time:", time.Since(now), i, "hashes")
		os.Exit(0)
	}()
	Start(*sleep, *conc)
}
