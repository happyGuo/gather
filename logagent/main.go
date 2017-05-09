package main

import (
	"./agent"
)

func main() {

	notifyPath := os.Args[1]
	_, err := os.Stat(notifyPath)
	if err != nil {
		log.Fatal(err)
	}

	agent.FsNotify(notifyPath)

}
