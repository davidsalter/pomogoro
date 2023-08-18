package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	durationMins := flag.Int("duration", 30, "Concentration time in minutes")
	flag.Parse()

	duration := *durationMins * 60
	fmt.Println("Starting new session")
	for duration > 0 {
		time.Sleep(1 * time.Second)
		duration--
		displayRemainingTime(duration)
	}
	fmt.Println("\nCongratulations. You've finished the session.")
}

func displayRemainingTime(duration int) {
	minutes := duration / 60
	seconds := duration % 60
	fmt.Printf("\033[0GTime remaining: %02d:%02d", minutes, seconds)
}
