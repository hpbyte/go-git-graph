package main

import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

type ConfigLoader[T any] interface {
	Load() (T, error)
}

type GitConfig struct {
	User struct {
		Name  string `ini:"name"`
		Email string `ini:"email"`
	} `ini:"user"`
}

type Config struct {
	BasePath  string
	GitConfig GitConfig
}

func (c *Config) Load() (Config, error) {
	if err := c.loadFlags(); err != nil {
		return Config{}, err
	}

	if err := c.loadGitConfig(); err != nil {
		return Config{}, err
	}

	return *c, nil
}

func (config *Config) loadGitConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error loading home dir: %s\n", err)
	}
	gitConfigPath := filepath.Join(homeDir, ".gitconfig")

	gConf, err := ini.Load(gitConfigPath)
	if err != nil {
		return fmt.Errorf("error loading git config: %s\n", err)
	}

	err = gConf.MapTo(&config.GitConfig)
	if err != nil {
		return fmt.Errorf("error parsing .gitconfig: %s\n", err)
	}

	return nil
}

func (config *Config) loadFlags() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current dir: %w", err)
	}

	flag.StringVar(&config.BasePath, "p", currentDir, "directory to cacluate stats for")
	flag.Parse()

	config.BasePath, err = filepath.Abs(config.BasePath)
	if err != nil {
		return fmt.Errorf("cannot find the directory provided!, %w", err)
	}

	return nil
}
