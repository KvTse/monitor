package main

import (
	"github.com/shirou/gopsutil/process"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	for {
		select {
		case <-time.After(time.Second):
			p, err := process.NewProcess(6740)
			if err != nil {
				log.Println(err)
				break
			}
			files, err := p.OpenFiles()
			if err != nil {
				log.Println(err)
				break
			}
			for _, file := range files {
				log.Println(file.String())
			}
		}
	}
}
