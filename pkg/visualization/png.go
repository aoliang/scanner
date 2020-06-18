package visualization

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func PrintDag(packageImports map[string][]string, dotFile, pngFile string) {
	if err := os.Remove(dotFile); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	f, err := os.Create(dotFile)
	if err != nil {
		panic(err)
	}


	//打印dot数据结构
	buf := bufio.NewWriterSize(f, 1024)
	buf.WriteString("digraph G{\n")
	for packageName, imports := range packageImports {
		for i := range imports {
			buf.WriteString(fmt.Sprintf("\"%s\" -> \"%s\";\n", packageName, imports[i]))
		}
	}
	buf.WriteString("}")
	buf.Flush()

	//dot转换为png
	cmd := exec.Command("dot", "-Tpng", dotFile, "-o", pngFile)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	log.Println(string(output))
}
