package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// testkjorer evig lokke med 4 forskjellige terminal kommandoer
	indefLoop()
}

func indefLoop() {

	//setup for aa sende signal medinger
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGHUP,
		syscall.SIGILL)

	for {
		fmt.Println("Indefinite loop running ...")
		s := <-c
		switch s {
		case syscall.SIGINT:
			fmt.Println("Process terminated - SIGINT")
		case syscall.SIGQUIT:
			fmt.Println("Terminal quit - SIGQUIT")
		case syscall.SIGHUP:
			fmt.Println("Hangup - SIGHUP")
		case syscall.SIGILL:
			fmt.Println("Illegal instruction - SIGILL")
		}
		break
	}
}
