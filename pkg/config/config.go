package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/ini.v1"
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
	BasePath   string
	ClearCache bool
	Year       int
	GitConfig  GitConfig
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
		return fmt.Errorf("[Err]: loading home dir: %s\n", err)
	}
	gitConfigPath := filepath.Join(homeDir, ".gitconfig")

	gConf, err := ini.Load(gitConfigPath)
	if err != nil {
		return fmt.Errorf("[Err]: loading git config: %s\n", err)
	}

	err = gConf.MapTo(&config.GitConfig)
	if err != nil {
		return fmt.Errorf("[Err]: parsing .gitconfig: %s\n", err)
	}

	return nil
}

func (config *Config) loadFlags() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("[Err]: error getting current dir: %w", err)
	}

	flag.StringVar(&config.BasePath, "p", currentDir, "directory to cacluate stats for")
	flag.BoolVar(&config.ClearCache, "c", false, "clear the cached repos list")
	flag.IntVar(&config.Year, "y", time.Now().Year(), "year to be aggregated")
	flag.Parse()

	config.BasePath, err = filepath.Abs(config.BasePath)
	if err != nil {
		return fmt.Errorf("[Err]: finding the directory provided!, %w", err)
	}

	return nil
}
