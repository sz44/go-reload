package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
)

// TODO: make build part of ProcessManager
func build() {
	cmd := exec.Command("go", "build", "-o", "./server/main", "./server")
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
	pm.cmd = exec.Command("./server/main")
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

func handleProcess(pm *ProcessManager) {
	pm.Stop()
	build()
	time.Sleep(50 * time.Millisecond)
	pm.Start()
}
func Watch(watcher *fsnotify.Watcher, pm *ProcessManager) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Println("event: ", event)
			if event.Has(fsnotify.Write) {
				pm.Stop()
				build()
				pm.Start()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error", err)
		}
	}
}

func main() {
	// initial build
	build()
	// initial run
	serverProcess := NewProcessManager()
	serverProcess.Start()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	fmt.Println("Starting live reload watcher...")
	go Watch(watcher, serverProcess)
	err = watcher.Add("./server")
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
