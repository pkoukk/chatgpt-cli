package main

import "github.com/pkoukk/chatgpt-cli/ui"

func main() {
	cli := ui.NewCLI()
	cli.Start()
	defer cli.Close()
}
