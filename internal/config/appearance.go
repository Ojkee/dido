package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"dido/internal/pathhandler"
)

type appearance struct {
	FontName string `json:"font_name"`
}

type AppearanceLoader struct {
	fileName string
	appearance
}

func NewAppearanceLoader() *AppearanceLoader {
	return &AppearanceLoader{
		fileName:   "appearance.json",
		appearance: appearance{},
	}
}

func (a *AppearanceLoader) Unmarshal() error {
	fontDir, err := pathhandler.ProjectPath()
	if err != nil {
		return err
	}
	jsonPath := filepath.Join(fontDir, "cfg", a.fileName)
	jsonFile, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	json.Unmarshal(jsonFile, &a.appearance)
	return nil
}

func (a *AppearanceLoader) Get() *appearance {
	return &a.appearance
}
