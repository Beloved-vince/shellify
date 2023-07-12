package main

import (
	"bufio"
	"fmt"

	"github.com/logrusorgru/aurora"

	// "golang/package"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func DirPath() (string, error) {
	//  return path to current working directory

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting directory path: ", err)
	}
	prefix := "C:\\Users\\yyyy"

	// Remove the prefix from the path
	trimmedPath := strings.TrimPrefix(dir, prefix)

	return trimmedPath, err
}

func sysDetail() (string, string, error) {
	// Return PC host name and  username and return LOGNAME if
	// username is empty

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
const (
	SHELL_NAME = "SHELLIFY"
	reset      = "\033[0m"
	green      = aurora.Green
	Purple     = aurora.Magenta
	Yellow     = aurora.Yellow
	// green      = color.FgGreen.Render
	// purple     = color.FgMagenta.Render
	// yellow     = color.FgYellow.Render
	// reset      = color.Reset
)

func main() {
	dir, _ := os.Getwd()
	reader := bufio.NewReader(os.Stdin)
	hostname, username, _ := sysDetail()
	dir_path, err := DirPath()

	// reset := aurora.Reset
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for {

		fmt.Println()
		fmt.Print(aurora.Sprintf("%s%s%s %s %s", green(hostname), green("@"), green(username), Purple(SHELL_NAME), Yellow("~"+dir_path)))
		fmt.Print(reset + " $ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Exit the shell if the user enters "exit" or "quit"
		if input == "exit" || input == "quit" {
			break
		}

		if input == "list" || input == "ls" {
			files := listDirectory(dir)
			for _, file := range files {
				fmt.Print(file.Name() + "\t")
			}
		}

		args := strings.Split(input, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err := cmd.Run(); err != nil {
			continue
		}
	}
}

func listDirectory(dirPath string) []os.DirEntry {
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		return nil
	}
	return dir
}
