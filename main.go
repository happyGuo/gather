package main

import (
	"fmt"
	"os"
	//"gather/collector"
	//"gather/agent"
	"gather/agent"
	"log"
)

func main() {

	launchPath := os.Args[1]
	_, err := os.Stat(launchPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("start...")
	agent.FsNotify(launchPath)

}
