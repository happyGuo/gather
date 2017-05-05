package agent

import (
	"log"
	"github.com/fsnotify/fsnotify"
	//"github.com/hpcloud/tail"
	//"fmt"
)

func FsNotify(path string)  {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					filePath := event.Name[0:len(event.Name)-12]
					log.Println("modified file:", filePath)



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
