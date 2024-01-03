package main

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
)

type GitConfig struct {
	User struct {
		Name  string `ini:"name"`
		Email string `ini:"email"`
	} `ini:"user"`
}

func LoadGitConfig() GitConfig {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error loading home dir: %s\n", err)
	}
	gitConfigPath := filepath.Join(homeDir, ".gitconfig")

	config, err := ini.Load(gitConfigPath)
	if err != nil {
		log.Fatalf("error loading git config: %s\n", err)
	}

	var gitConfig GitConfig
	err = config.MapTo(&gitConfig)
	if err != nil {
		log.Fatalf("error parsing .gitconfig: %s\n", err)
	}

	return gitConfig
}
