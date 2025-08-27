package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/simple_cli/internal/db"
)

func Add() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter user name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("User name cannot be empty.")
		return
	}

	_, err := db.DB.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		fmt.Printf("Failed to add user: %v\n", err)
		return
	}

	fmt.Printf("User '%s' added successfully!\n", name)
}