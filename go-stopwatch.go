package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	tm "github.com/buger/goterm"
	"github.com/frozzare/go-beeper"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func playMp3(songName string) error {
	f, err := os.Open(songName)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

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
			if len(os.Args) > 2 && os.Args[2] == "break" || len(os.Args) > 3 &&  os.Args[3] == "break" {
				break
			}
		}
		//tm.MoveCursorUp(1)
		tm.Flush()
		if elapsedSecs >= timer {
			if len(os.Args) > 2 && os.Args[2] == "playDing" {
				if err := playMp3("/Users/abhijit/Documents/go-stopwatch/Ding.mp3"); err != nil {
					log.Fatal(err)
				}
			} else {
				beeper.Beep()
			}
		}
	}
}
