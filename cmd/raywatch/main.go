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
	poll       = flag.Duration("poll", 10*time.Millisecond, "poll interval")
	tresh      = flag.Duration("trhesh", 500*time.Millisecond, "poll interval")
	flagWidth  = flag.Int("w", 400, "image width")
	flagHeight = flag.Int("h", 300, "image height")
)

var (
	cmd *exec.Cmd
)

func main() {

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("need one argument")
	}
	fname := flag.Arg(0)

	Print(exec.Command("x-www-browser", "localhost:3700").Start())
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
	build(fname)
	goServe()
}

const Executable = "/tmp/bruteray-scene"

func build(fname string) error {
	start := time.Now()
	build := exec.Command("go", "build", "-o", Executable, fname)
	build.Stderr = os.Stderr
	err := build.Run()
	log.Println("build", ok(err), time.Since(start).Round(time.Millisecond))
	return err
}

func ok(err error) string {
	if err == nil {
		return "OK"
	}
	return err.Error()
}

func goServe() {
	if cmd != nil {
		Print(cmd.Process.Kill())
		cmd.Process.Wait()
	}
	cmd = exec.Command(Executable, fmt.Sprintf("-w=%v", *flagWidth), fmt.Sprintf("-h=%v", *flagHeight))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	log.Println("serve", ok(err))
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
