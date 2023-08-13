package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./nmap_scanner <IP>")
		return
	}

	ip := os.Args[1]

	// Run nmap command
	cmd := exec.Command("nmap", "-sC", "-sV", "-T5", "-p-", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running nmap:", err)
		return
	}

	// Get the current directory name
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	dirName := filepath.Base(currentDir)

	// Generate the output filename
	outputFile := fmt.Sprintf("nmap_%s.txt", dirName)

	// Write output to a file
	err = ioutil.WriteFile(outputFile, output, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Nmap output saved to %s\n", outputFile)
}
