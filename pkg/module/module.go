package module

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var module string

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open("go.mod")
	if err != nil {
		panic("go.mod not exist. pwd=" + dir)
	}

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		module = strings.TrimSpace(strings.TrimLeft(line, "module"))
		if module == "" {
			continue
		}
		if module != "" {
			break
		}
	}

	if module == "" {
		panic("go.mod has no module declaration")
	}
}

func Get() string {
	return module
}
