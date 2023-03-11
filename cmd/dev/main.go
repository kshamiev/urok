package main

// https://stackoverflow.com/questions/71418671/restart-or-shutdown-golang-apps-programmatically

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/radovskyb/watcher"
)

func main() {
	log.Println("APP START OR RESTART FROM CHANGE BINARY")
	err := restart()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour * 1)
}

func restart() error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	t_watch := watcher.New()
	t_watch.FilterOps(watcher.Write)
	err = t_watch.Add(self)
	if err != nil {
		return err
	}
	go func() {
		var err error
		for {
			_ = <-t_watch.Event
			log.Println("RESTART")
			args := os.Args
			env := os.Environ()
			// Windows does not support exec syscall.
			if runtime.GOOS == "windows" {
				cmd := exec.Command(self, args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				cmd.Env = env
				err = cmd.Run()
				if err == nil {
					os.Exit(0)
				}
			} else {
				err = syscall.Exec(self, args, env)
			}
			log.Println(err)
			// TODO logging error
		}
	}()
	go func() {
		_ = t_watch.Start(time.Second * 1)
	}()
	return nil
}
