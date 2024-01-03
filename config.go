package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
)

type Config interface {
	Load() error
}

type GitConfig struct {
	User struct {
		Name  string `ini:"name"`
		Email string `ini:"email"`
	} `ini:"user"`
}

func (gc *GitConfig) Load() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error loading home dir: %s\n", err)
	}
	gitConfigPath := filepath.Join(homeDir, ".gitconfig")

	config, err := ini.Load(gitConfigPath)
	if err != nil {
		return fmt.Errorf("error loading git config: %s\n", err)
	}

	err = config.MapTo(&gc)
	if err != nil {
		return fmt.Errorf("error parsing .gitconfig: %s\n", err)
	}

	return nil
}
