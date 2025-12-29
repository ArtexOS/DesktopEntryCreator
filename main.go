package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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
	var appName, execPath, iconPath, categories string
	var err error
	s := bufio.NewReader(os.Stdin)
	fields := []struct {
		label  string
		target *string
	}{
		{"Название", &appName},
		{"Путь к приложению", &execPath},
		{"Путь к иконке", &iconPath},
		{"Категория", &categories},
	}

	for _, f := range fields {
		val, err := userInput(s, "Введите"+" "+f.label+":")
		if err != nil || strings.TrimSpace(val) == "" {
			fmt.Println("Ошибка ввода или пустое значение!")
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
		fmt.Printf("Файла %s не существует\n", execPath)
		return
	}

	_, err = os.Stat(iconPath)
	if os.IsNotExist(err) {
		fmt.Printf("Файла %s не существует\n", iconPath)
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
	fmt.Println("Файл", fileName+".desktop", "был успешно создан по пути", filePath)
}
