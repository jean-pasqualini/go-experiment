package main

import (
	"fmt"
	"github.com/bcicen/grmon/agent"
	"github.com/jean-pasqualini/go-experiment/fifo"
	"github.com/jean-pasqualini/go-experiment/queue"
	"github.com/manifoldco/promptui"
)

func main() {
	grmon.Start()

	queue.Handle()
}

func out() {
	prompt := promptui.Select{
		Label: "Select handler",
		Items: []string{"queue", "fifo"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "queue":
		queue.Handle()
		break
	case "fifo":
		fifo.Handle()
		break
	}
}
