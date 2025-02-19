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
		log.Fatalf("Error executing command: %v\n", err)
		// fmt.Printf("Error executing command: %v\n", err)
		// return
	}

	fmt.Println("Build completed successfully.")
}

type ProcessManager struct {
	cmd *exec.Cmd
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{}
}

func (pm *ProcessManager) Start() {
	if pm.cmd != nil && pm.cmd.Process != nil {
		log.Println("Process already running!")
		return
	}
	pm.cmd = exec.Command("./servermain")
	pm.cmd.Stderr = os.Stderr
	pm.cmd.Stdout = os.Stdout

	log.Println("Starting process")
	err := pm.cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func (pm *ProcessManager) Stop() {
	if pm.cmd == nil || pm.cmd.Process == nil {
		log.Println("No process running!")
		return
	}
	log.Println("Stopping process")
	err := pm.cmd.Process.Kill()
	if err != nil {
		log.Println("err: ", err)
	}

	pm.cmd.Wait()
	pm.cmd = nil
}

func test() {
	prog := NewProcessManager()
	prog.Start()
	time.Sleep(1 * time.Second)
	prog.Stop()
	time.Sleep(50 * time.Millisecond)
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

func main() {
	// initial build
	// buildMain()
	// initial run
	serverProcess := NewProcessManager()
	serverProcess.Start()
	time.Sleep(2 * time.Second)
	serverProcess.Stop()
	time.Sleep(2 * time.Second)
	serverProcess.Start()
}
