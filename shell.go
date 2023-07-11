package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	red      = "\033[31m"
	Blue     = "\033[34m"
	Purple   = "\033[35m"
	green    = "\033[32m"
	yellowBg = "\033[43m"
	reset    = "\033[0m"
)

func DirPath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error ", err)
	}
	desktopPath := filepath.Join(currentUser.HomeDir, "Desktop")
	golangPath := filepath.Join(desktopPath, "golang")

	return golangPath, err
}

func sysDetail() (string, string, error) {
	// Return User home directory and current directory of the user
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf(err.Error())
	}
	username := os.Getenv("USERNAME")
	if username == "" {
		username = os.Getenv("LOGNAME")
	}

	if username == "" {
		// If the environment variables are not set, use the current user
		user, err := user.Current()
		if err != nil {
			fmt.Println("Error getting username:", err)
			return "", "", fmt.Errorf("error getting username: %v", err)
		}
		username = user.Username

	}

	return username, hostname, err

}

// func ViewCurrentDir() {

// }

func main() {

	reader := bufio.NewReader(os.Stdin)
	hostname, username, err := sysDetail()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for {

		fmt.Println()
		fmt.Printf(green+"%s@%s "+Purple+"Hamsel ", hostname, username)
		fmt.Print(reset + "$")
		fmt.Printf(reset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Exit the shell if the user enters "exit" or "quit"
		if input == "exit" || input == "quit" {
			break
		}

		args := strings.Split(input, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
