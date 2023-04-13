package main

import (
	"fmt"
	"go-powershell/backend"
)

func main() {
	back := &backend.Local{}

	shell, err := New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	stdout, _, err := shell.Execute("Get-WmiObject -Class Win32_Processor")
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout)
}
