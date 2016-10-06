package main

import "github.com/fsnotify/fsnotify"
import "github.com/pkg4go/pathx"
import "github.com/pkg4go/execx"
import "log"

func watchDir(dir, cmd string, args ...string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	path := pathx.Resolve("", dir)

	log.Printf("watching dir: %s \n", path)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					out, err := execx.Run(cmd, args...)
					if err != nil {
						log.Println("error:", err.Error())
					} else {
						log.Println(out)
					}
				}
			case err := <-watcher.Errors:
				log.Println("error:", err.Error())
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		panic(err)
	}
	<-done
}
