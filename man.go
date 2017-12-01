package main

import "os"

import "os/exec"

func main() {
	page := ""
	if len(os.Args) > 1 {
		page = os.Args[1]
	}
	cmd := exec.Command("open", "-a", "/Applications/Bwana.app", "man://"+page)
	cmd.Run()
}
