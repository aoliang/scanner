package browser

import (
	"log"
	"os/exec"
)

//调用mac open命令打开浏览器。不支持windows。
func Request(link string) {
	err := exec.Command(`open`, link).Start()
	if err != nil {
		log.Println("open browser error.", err)
	}
}
