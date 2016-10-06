package main

import "github.com/pkg4go/execx"
import "github.com/jinzhu/now"
import "time"
import "log"

func watchTime(pattern, cmd string, args ...string) {
	log.Printf("watching time: %s \n", pattern)

	done := make(chan bool)
	go func() {
		timer := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-timer.C:
				if match(pattern) {
					out, err := execx.Run(cmd, args...)
					if err != nil {
						log.Println("error:", err.Error())
					} else {
						log.Println(out)
					}
				}
			}
		}
	}()

	<-done
}

func match(pattern string) bool {
	n := time.Now().Format("2006-01-02 15:04:05")

	t := now.MustParse(pattern).Format("2006-01-02 15:04:05")

	return n == t
}
