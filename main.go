package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"os/exec"
)

func standardizeSpaces(s string) string {
	return strings.Replace(strings.Join(strings.Fields(s), " "), " ", "-", -1)
}

func main() {
	fmt.Print("title: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	title := scanner.Text()

	date := time.Now().Format("20060102")
	
	filepath := "posts/" + date + "-" + standardizeSpaces(title) + ".md"

	cmd := exec.Command("hugo", "new", filepath) 
	cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
	cmd.Run()
}
