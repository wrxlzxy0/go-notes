package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var spisok []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "стоп" {
			break
		}
		spisok = append(spisok, line)
	}
	skleika := strings.Join(spisok, "\n")
	err := os.WriteFile("notes.txt", []byte(skleika), 0644)
	if err != nil {
		fmt.Println("ошибка", err)
	}
	file, oshibka := os.ReadFile("notes.txt")
	if oshibka != nil {
		fmt.Println("ошибка", oshibka)
		return
	}
	во
	fmt.Println("--- сохранено: ---")
	fmt.Println(string(file))
}
