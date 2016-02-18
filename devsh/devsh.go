package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("cwd:", cwd)
	cmd := exec.Command("gnome-terminal", "-x", "bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr 
	cmd.Run()
}
