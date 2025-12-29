![Build Status](https://github.com/ArtexOS/DesktopEntryCreator/actions/workflows/release.yml/badge.svg)

# [RU] Desktop Entry Generator

CLI-утилита на языке **Go** для быстрого создания ярлыков приложений (`.desktop` файлов) в Linux. Оптимизирована для окружений на базе Debian/Ubuntu (протестировано на Linux Mint).

## Функции

* **Интерактивная настройка**: Пошаговый ввод названия, пути к исполняемому файлу, иконки и категории.
* **Расширение путей**: Автоматическое преобразование тильды в полный путь (например, `~/` в `/home/user/`).
* **Валидация**: Проверка существования файлов перед созданием ярлыка для предотвращения появления нерабочих пунктов меню.
* **Совместимость с XDG**: Автоматическое создание директории `~/.local/share/applications/`, если она отсутствует.

## Установка

### Из исходников

Если у вас установлен Go:

```bash
go run main.go
```

### Компиляция и установка

Чтобы использовать утилиту как системную команду:

1. Скомпилируйте бинарный файл:
```bash
go build -o desktop-gen main.go
```


2. Переместите его в директорию bin:
```bash
sudo mv desktop-gen /usr/local/bin/
```



### Готовые бинарные файлы

Вы можете скачать скомпилированные версии в разделе [Releases](https://github.com/ArtexOS/DesktopEntryCreator/releases). Скачайте версию для вашей архитектуры, дайте права на выполнение и используйте.

## Использование

1. Запустите программу командой `desktop-gen`.
2. Введите **Name** (например, `My App`).
3. Укажите **Exec Path** (поддерживается `~/`).
4. Укажите **Icon Path**.
5. Выберите **Category**.

## Поиск и устранение неисправностей

Если ярлык не появился:

1. Проверьте наличие файла: `ls ~/.local/share/applications/`.
2. Обновите кэш меню: `update-desktop-database ~/.local/share/applications`.
3. Проверьте права доступа (рекомендуется `644`).

---

# [EN] Desktop Entry Generator

A lightweight CLI utility written in **Go** designed to quickly create application shortcuts (`.desktop` files) for Linux environments. Optimized for Debian/Ubuntu-based distributions (tested on Linux Mint).

## Features

* **Interactive Setup**: Step-by-step CLI prompts for application name, executable path, icon, and category.
* **Tilde Expansion**: Automatically resolves home directory shortcuts (e.g., converts `~/` to `/home/user/`).
* **Path Validation**: Verifies the existence of the executable and icon files before generating the entry to prevent broken shortcuts.
* **XDG Compliance**: Automatically ensures the target directory `~/.local/share/applications/` exists.

## Installation

### From Source

If you have the Go toolchain installed:

```bash
go run main.go
```

### Manual Binary Installation

To use the utility as a system-wide command:

1. Build the binary:
```bash
go build -o desktop-gen main.go
```


2. Move it to your local bin directory:
```bash
sudo mv desktop-gen /usr/local/bin/
```



### Download Pre-built Binaries

Compiled binaries are available in the [Releases](https://github.com/ArtexOS/DesktopEntryCreator/releases) section. Download the version corresponding to your architecture and grant it execution permissions.

## Usage

1. Launch the program by typing `desktop-gen`.
2. Enter the **Application Name**.
3. Provide the **Executable Path** (supports `~/`).
4. Provide the **Icon Path**.
5. Select an **Application Category**.

## Troubleshooting

If the shortcut does not appear in your application menu:

1. **Verification**: Confirm the file exists by running `ls ~/.local/share/applications/`.
2. **Cache**: Refresh the desktop database manually:
```bash
update-desktop-database ~/.local/share/applications
```


3. **Permissions**: Ensure the generated file has the correct read permissions (typically `644`).


---

## Acknowledgments / Благодарности

* **English:** This project was developed with the assistance of **Gemini** (a large language model by Google), which helped with Go syntax optimization, code refactoring, and documentation (README) composition.


* **Русский:** Этот проект был разработан при содействии **Gemini** (большой языковой модели от Google). Модель помогла с оптимизацией синтаксиса Go, рефакторингом кода и составлением документации (README).