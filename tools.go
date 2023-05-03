package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func RunAgentBinaryFile() (int, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return 0, err
	}

	binaryFileName := ""
	switch runtime.GOOS {
	case "linux":
		binaryFileName = "HelloWorldGoAgent_binary_build"
	case "windows":
		binaryFileName = "HelloWorldGoAgent_binary_build.exe"
	default:
		fmt.Println("Unsupported operating system")
		os.Exit(1)
	}
	binaryFilePath := filepath.Join(dir, binaryFileName)
	log.Println("logPath", binaryFilePath)

	// Set the path to your binary file
	cmd := exec.Command(binaryFilePath)

	// Start the process
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting process:", err)
		return 0, err
	}

	// Get the process ID
	pid := cmd.Process.Pid
	fmt.Println("Process started with PID:", pid)

	// * Wait for the process to finish, we don't need to wait for it because we don't want this process to be blocked
	// if err := cmd.Wait(); err != nil {
	// 	fmt.Println("Process exited with error:", err)
	// } else {
	// 	fmt.Println("Process exited successfully")
	// }

	return pid, nil
}
