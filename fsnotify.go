package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
)

func buildMain() {
	cmd := exec.Command("go", "build", "-o main", ".")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}

	fmt.Println("Build completed successfully.")
}

type MainExe struct {
	cmd    *exec.Cmd
	status string
}

func NewMainExe() *MainExe {
	// cmd := exec.Command("gofromgo/main")
	cmd := exec.Command("./servermain")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return &MainExe{cmd, "created"}
}

func (main *MainExe) Start() {
	log.Println("starting program...")
	err := main.cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func (main *MainExe) Stop() {
	log.Println("stopping program...")
	err := main.cmd.Process.Kill()
	if err != nil {
		log.Println("err: ", err)
	}
}

func test() {
	prog := NewMainExe()
	prog.Start()
	time.Sleep(1 * time.Second)
	prog.Stop()
	time.Sleep(50 * time.Millisecond)
}

func runMain() {
	// cmd := exec.Command("gofromgo/main", "../oop/adapter.go")
	cmd := exec.Command("gofromgo/main")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
	cmd.Process.Kill()
	time.Sleep(5 * time.Second)
}

func prog() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event: ", event)
				if event.Has(fsnotify.Write) {
					// stop
					// build with command go build .
					// run the binary
					log.Println("modified file:", event.Name)

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan any)
}
