package printer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clearOsMap map[string]func()

func init() {
	clearOsMap := make(map[string]func())
	clearOsMap["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error while clearing terminal screen:", err)
		}
	}
	clearOsMap["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error while clearing terminal screen:", err)
		}
	}
}

type StdoutPrinter struct {
}

func (s *StdoutPrinter) Print(msg string) {
	fmt.Println(msg)
}

func (s *StdoutPrinter) Clear() {
	value, ok := clearOsMap[runtime.GOOS]
	if ok {
		value()
	} else {
		fmt.Println("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
