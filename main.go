package main

import (
	"fmt"
	"os"
	//"gather/collector"
	//"gather/agent"
	"log"
	"gather/agent"
)

func main() {

	launchPath := os.Args[1]
	_, err := os.Stat(launchPath)
	if err != nil {
		log.Fatal(err)
	}
	agent.FsNotify(launchPath)
	fmt.Println("start...")

}