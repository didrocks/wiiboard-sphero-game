package main

import (
	"fmt"
	"os"

	"github.com/didrocks/wiiboard-sphero-game/internal/wiiboard"
)

func main() {
	board, err := wiiboard.Detect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	battery, err := board.Battery()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(battery)
	go board.Listen()
	board.Calibrate()
	for event := range board.Events {
		fmt.Printf("%+v\n", event)
	}
}
