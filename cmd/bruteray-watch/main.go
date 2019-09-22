package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

var (
	flagSleep = flag.Duration("i", 30*time.Millisecond, "poll time")
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fatal("need at least 1 argument: go file to watch")
	}
	watch(flag.Arg(0), flag.Args()[1:])
}

func watch(fname string, args []string) {
	var ts time.Time
	for {
		time.Sleep(*flagSleep)
		info, err := os.Stat(fname)
		check(err)
		if info.ModTime() != ts {
			rerun(fname, args)
			ts = info.ModTime()
		}
	}
}

var last *exec.Cmd

func rerun(fname string, args []string) {
	if last != nil {
		if err := last.Process.Kill(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := last.Wait(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	if err := compile(fname); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	dir := path.Dir(fname)
	base := path.Base(fname)
	noExt := base[:len(base)-len(path.Ext(base))]
	allArgs := append([]string{"-o", noExt + ".jpg"}, args...)
	cmd := exec.Command("./"+noExt, allArgs...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	out, err := exec.Command("renice", "20", fmt.Sprint(cmd.Process.Pid)).CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "renice: %s", out)
	}

	last = cmd
}

func compile(fname string) error {
	dir := path.Dir(fname)
	cmd := exec.Command("go", "build", path.Base(fname))
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func check(e error) {
	if e != nil {
		fatal(e)
	}
}

func fatal(x ...interface{}) {
	fmt.Fprintln(os.Stderr, x...)
	os.Exit(1)
}
