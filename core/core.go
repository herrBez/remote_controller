package core

import (
  "fmt"
  "os/exec"
  "bytes"
)


// Simple wrapper that execute a command and prints the output result in the
// command line
// https://stackoverflow.com/questions/18159704/how-to-debug-exit-status-1-error-when-running-exec-command-in-golang
func ExecuteCommand(cmd *exec.Cmd) { 
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
	  fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	  return
	}
	fmt.Println("Result: " + out.String())   
}
