package config

import (
	"log"
	"path/filepath"

	pathhandler "dido/internal/pathhandler"
)

type Config struct {
	appearance *AppearanceLoader
}

func NewConfig() *Config {
	appearance := NewAppearanceLoader()
	err := appearance.Unmarshal()
	if err != nil {
		log.Fatalf("Config/Load appearance: %v", err)
	}
	return &Config{
		appearance: appearance,
	}
}

func (c *Config) Appearance() *appearance {
	return c.appearance.Get()
}

func (c *Config) FontPath() (string, error) {
	fontsDir, err := pathhandler.FontDirPath()
	if err != nil {
		return "", err
	}
	fontName := filepath.Join(fontsDir, c.appearance.FontName)
	return fontName, nil
}
