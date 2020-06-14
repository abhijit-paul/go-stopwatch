package main

import (
	"fmt"
	"os"
	"time"

	tm "github.com/buger/goterm"
	"github.com/frozzare/go-beeper"
)

func main() {
	clitimer := os.Args[1]
	timerD, _ := time.ParseDuration(clitimer)
	timer := int(timerD.Seconds())
	c := time.Tick(1 * time.Second)
	inow := time.Now()
	fmt.Println("Starting Timer")
	for now := range c {
		elapsedSecs := int(now.Sub(inow).Seconds())
		remainingTime := timer - elapsedSecs
		if remainingTime >= 0 {
			fmt.Printf("\r\033[KRemaining time: %d %s", (timer - elapsedSecs), "s")
		} else {
			fmt.Printf("\r\033[KTime UP!")
		}
		//tm.MoveCursorUp(1)
		tm.Flush()
		if elapsedSecs >= timer {
			beeper.Beep()
		}
	}
}
