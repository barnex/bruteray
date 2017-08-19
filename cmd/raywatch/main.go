package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	poll  = flag.Duration("poll", 10*time.Millisecond, "poll interval")
	tresh = flag.Duration("trhesh", 500*time.Millisecond, "poll interval")
)

var (
	cmd *exec.Cmd
)

func main() {
	log.SetFlags(0)

	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("need one argument")
	}
	fname := flag.Arg(0)

	watch(fname)
}

func watch(fname string) {
	var prev os.FileInfo
	for range time.Tick(*poll) {
		fi, err := os.Stat(fname)
		if err == nil && !fiEq(fi, prev) {
			trigger(fname)
		}
		prev = fi
	}
}

func trigger(fname string) {
	log.Println("triggered", fname)
	err := build(fname)
	if err == nil {
		goServe()
	}
}

const Executable = "/tmp/bruteray-scene"

func build(fname string) error {
	err := exec.Command("go", "build", "-o", Executable, fname).Run()
	log.Println("build", err)
	//Print(err)
	return err
}

func goServe() {
	if cmd != nil {
		log.Println("killing previous...")
		Print(cmd.Process.Kill())
		cmd.Process.Wait()
		log.Println("...killed")
	}
	log.Println("starting")
	cmd = exec.Command(Executable)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	Print(cmd.Start())
}

// FileInfo equality.
func fiEq(a, b os.FileInfo) bool {
	switch {
	default:
		panic("bug")
	case a == nil && b == nil:
		return true
	case a != nil && b == nil:
		return false
	case a == nil && b != nil:
		return false
	case a != nil && b != nil:
		return a.ModTime() == b.ModTime() && a.Size() == b.Size()
	}
}

func Print(err error) {
	if err != nil {
		log.Println(err)
	}
}
