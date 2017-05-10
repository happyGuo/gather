package agent

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	//	"github.com/hpcloud/tail"
	"runtime"
)

var (
	server = "127.0.0.1:9876"
)

func FsNotify(path string) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	fileOffset := make(map[string]int64)
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					filePath := event.Name
					if runtime.GOOS == "windows" {
						filePath = event.Name[0 : len(event.Name)-12]
					}

					log.Println("modified file:", filePath)
					var offset int64 = 0
					if _, ok := fileOffset[filePath]; ok {
						offset = fileOffset[filePath]

					}
					log.Println("offset:", fileOffset)
					curOffset := readLine(filePath, offset)

					fileOffset[filePath] = curOffset
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func readLine(fileName string, offset int64) int64 {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("failed to open")

	}

	file.Seek(offset, os.SEEK_SET)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //每次读取一行

		if err != nil {
			break // 读完或发生错误
		}
		fmt.Printf(str)
		sendToCollector(str)
	}
	curOffset, _ := file.Seek(0, os.SEEK_CUR)
	return curOffset
}

func sendToCollector(str string) {
	sender := NewSender()
	sender.SendData(str)

}
