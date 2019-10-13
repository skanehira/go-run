package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/skanehira/go-logger"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	debug    = flag.Bool("debug", false, "print debug log")
	contents = flag.String("c", "", "source code")
)

func execute(contents string) int {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Errorf("%s", err)
		return 1
	}

	file, err := os.Create(filepath.Join(dir, "main.go"))
	if err != nil {
		log.Errorf("%s", err)
		return 1
	}
	log.Debugf("temp file: %s", file.Name())

	defer file.Close()
	defer os.RemoveAll(dir)

	if _, err := file.WriteString(contents); err != nil {
		log.Errorf("%s", err)
		return 1
	}

	command := exec.Command("go", "run", file.Name())
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	if err := command.Run(); err != nil {
		log.Errorf("%s", err)
		return 1
	}

	return 0
}

func run() int {
	if *contents != "" {
		return execute(*contents)
	}
	if !terminal.IsTerminal(0) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Errorf("%s", err)
			return 1
		}
		return execute(string(b))
	}
	return 0
}

func init() {
	flag.Parse()
	if *debug {
		log.SetMinLevel(log.DEBUG)
	}
	log.SetFlags(log.Lshortfile)
}

func main() {
	os.Exit(run())
}
