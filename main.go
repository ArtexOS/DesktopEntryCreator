package main

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed locales/*.json
var localesFS embed.FS

type Translation struct {
	Prompts  PromptTexts  `json:"prompts"`
	Labels   LabelTexts   `json:"labels"`
	Errors   ErrorTexts   `json:"errors"`
	Messages MessageTexts `json:"messages"`
}

type PromptTexts struct {
	Name     string `json:"name"`
	ExecPath string `json:"exec_path"`
	IconPath string `json:"icon_path"`
	Category string `json:"category"`
}

type LabelTexts struct {
	Name     string `json:"name"`
	ExecPath string `json:"exec_path"`
	IconPath string `json:"icon_path"`
	Category string `json:"category"`
}

type ErrorTexts struct {
	InputEmpty   string `json:"input_empty"`
	FileNotFound string `json:"file_not_found"`
	HomeDir      string `json:"home_dir"`
	WriteFile    string `json:"write_file"`
}

type MessageTexts struct {
	Success string `json:"success"`
}

func loadTranslation() Translation {
	langEnv := os.Getenv("LANG")
	fileName := "locales/en.json"

	if strings.HasPrefix(langEnv, "ru") {
		fileName = "locales/ru.json"
	}

	data, err := localesFS.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading embedded file %s: %v\n", fileName, err)
		return Translation{}
	}

	var t Translation
	err = json.Unmarshal(data, &t)
	if err != nil {
		fmt.Println("Критическая ошибка: файл не найден!", err)
		return Translation{}
	}

	return t
}

func userInput(reader *bufio.Reader, text string) (string, error) {
	fmt.Println(text)
	variable, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(variable), nil
}

func tildeExpansion(path string) (string, error) {
	if !strings.HasPrefix(path, "~/") {
		return path, nil
	}

	path = strings.TrimPrefix(path, "~")
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homePath, path), nil
}

func main() {
	
	lang := loadTranslation()

	var appName, execPath, iconPath, categories string
	var err error
	s := bufio.NewReader(os.Stdin)

	fields := []struct {
		label  string
		prompt string
		target *string
	}{
		{lang.Labels.Name, lang.Prompts.Name, &appName},
		{lang.Labels.ExecPath, lang.Prompts.ExecPath, &execPath},
		{lang.Labels.IconPath, lang.Prompts.IconPath, &iconPath},
		{lang.Labels.Category, lang.Prompts.Category, &categories},
	}

	for _, f := range fields {
		val, err := userInput(s, f.prompt)
		if err != nil || strings.TrimSpace(val) == "" {
			fmt.Println(lang.Errors.InputEmpty)
			return
		}
		*f.target = val
	}

	execPath, err = tildeExpansion(execPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	iconPath, err = tildeExpansion(iconPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = os.Stat(execPath)
	if os.IsNotExist(err) {
		fmt.Printf(lang.Errors.FileNotFound, execPath)
		return
	}

	_, err = os.Stat(iconPath)
	if os.IsNotExist(err) {
		fmt.Printf(lang.Errors.FileNotFound, iconPath)
		return
	}

	content := `
[Desktop Entry]
Type=Application
Version=1.0
Name=%s
Exec=%s
Icon=%s
Terminal=false
Categories=%s
`
	finalContent := fmt.Sprintf(content, appName, execPath, iconPath, categories)
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := strings.ReplaceAll(strings.ToLower(appName), " ", "-")
	filePath := filepath.Join(homePath, ".local", "share", "applications", fileName+".desktop")

	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile(filePath, []byte(finalContent), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(lang.Messages.Success, fileName, filePath)
}
