package main

import (
	"bufio"
	"os"

	"github.com/Ararni/cat-event-loop/engine"
)

func main() {
	Loop := new(engine.Loop)
	Loop.Start()
	if input, err := os.Open("input.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := engine.Parse(commandLine) // parse the line to get a Command
			Loop.Post(cmd)
		}
	}
	Loop.AwaitFinish()
}
