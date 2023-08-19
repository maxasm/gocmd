package main

import (
	"bytes"
	"os/exec"
	"fmt"
	"os"
)

func main() {
	cmd_name := "find"
	cmd_args := []string{"-type", "d"}
	
	cmd := exec.Command(cmd_name, cmd_args...) 	
	
	var std_out bytes.Buffer
	var std_err bytes.Buffer
	
	cmd.Stdout = &std_out	
	cmd.Stderr = &std_err
	
	err := cmd.Run()
	
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("\n ---- output ----\n\n%s\n", std_out.String())
	
	if std_err.Len() > 0 {
		fmt.Printf("\n ---- error ----\n\n%s\n", std_err.String())
	}
}
