package meta

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var project string

func init() {
	f, err := os.Open("go.mod")
	if err != nil {
		panic("go.mod not exist.")
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
		project = strings.TrimSpace(strings.TrimLeft(line, "module"))
		if project == "" {
			continue
		}
	}

	if project == "" {
		panic("go.mod has no module declaration")
	}
}

func GetProject() string {
	return project
}
