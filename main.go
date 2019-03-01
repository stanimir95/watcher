package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var dirFlag string
	if dirFlag == "" {
		dirFlag = currentDir
	}
	flag.StringVar(&dirFlag, "dir", "", "directory to watch")
	flag.Parse()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Printf("EVENT! MOFO %#v\n", event)

			case err := <-watcher.Errors:
				fmt.Println("Error", err)
			}
		}
	}()
	if err := watcher.Add(dirFlag); err != nil {
		fmt.Println("Error", err)
	}
	<-done
}
