package main

import (
	"bufio"
	"os"

	"github.com/nikolaichub/command-parser-event-loop/engine"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open("input.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := engine.Parse(commandLine) // parse the line to get a Command -> &printCommand{arg: “hello”}
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
