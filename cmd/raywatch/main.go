/*
Command raywatch watches a source file defining a bruteray scene,
and starts rendering the scene each time the file is changed.
The result is rendered in a browser.

E.g.:
	raywatch scenes/myscene.go
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	poll        = flag.Duration("poll", 10*time.Millisecond, "poll interval")
	tresh       = flag.Duration("trhesh", 500*time.Millisecond, "poll interval")
	flagWidth   = flag.Int("w", 960, "image width")
	flagHeight  = flag.Int("h", 540, "image height")
	flagBrowser = flag.String("browser", "x-www-browser", "display in this browser")
)

var (
	cmd *exec.Cmd
)

const Executable = "/tmp/bruteray-scene"

func main() {

	log.SetFlags(0)
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("raywatch: need one argument")
	}
	fname := flag.Arg(0)

	if *flagBrowser != "" {
		Print(exec.Command(*flagBrowser, "localhost:3700").Start())
	}
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
	log.Println("raywatch:", fname, "modified")
	kill()
	err := build(fname)
	if err != nil {
		return
	} else {
		goServe()
	}
}

func build(fname string) error {
	start := time.Now()
	build := exec.Command("go", "build", "-o", Executable, fname)
	build.Stderr = os.Stderr
	err := build.Run()
	log.Println("raywatch: build", ok(err), time.Since(start).Round(time.Millisecond))
	return err
}

func goServe() {
	cmd = exec.Command(Executable, fmt.Sprintf("-w=%v", *flagWidth), fmt.Sprintf("-h=%v", *flagHeight))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	log.Println("raywatch: serve", ok(err))
}

func kill() {
	if cmd != nil {
		log.Println("raywatch: kill", cmd.Path)
		err := cmd.Process.Kill()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Process.Wait()
		cmd = nil
	}
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

func ok(err error) string {
	if err == nil {
		return "OK"
	}
	return err.Error()
}

func Print(err error) {
	if err != nil {
		log.Println(err)
	}
}
