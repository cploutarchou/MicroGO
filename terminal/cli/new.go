package main

import (
	"errors"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
	"strings"
)

func createNew(applicationName string) {
	applicationName = strings.TrimSpace(applicationName)
	applicationName = strings.ToLower(applicationName)
	if applicationName == "" {
		gracefullyExit(errors.New("No project name specified! "))
	}

	// sanitize the application name
	if strings.Contains(applicationName, "/") {
		exploded := strings.SplitAfter(applicationName, "/")
		applicationName = exploded[len(exploded)-1]
	}
	log.Println("Application name: ", applicationName)
	// git clone skeleton application
	color.Green("\tCloning skeleton application from git repository...")

	_, err := git.PlainClone("./"+applicationName, false, &git.CloneOptions{
		URL:      "https://github.com/cploutarchou/microGo_skeleton_app.git",
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		gracefullyExit(err)
	}
	//remove the .git directory

	// create a new .env file

	// create a makefile for the application

	// update the go.mod file

	// update the existing .go files with th correct package names

	// run go mod tidy

}
