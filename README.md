# Whipr

**Whipr** is a simple system tray application for Linux that shows the currently selected text in a popup or notification.  
Ideal for quick translation workflows or clipboard management.

---

## Project Status

This project is an **application in progress**. Currently, it supports Linux, and support for macOS and Windows is planned for future releases.

---

## Features

- System tray icon with menu
- Show selected text in a popup window or system notification
- Lightweight, and easy to use

---

## Requirements

- **Go** (v1.18+ recommended)
- **xclip** (for reading selected text)
- **zenity** (for popup windows)
- **libnotify-bin** (for system notifications)
- **libayatana-appindicator3-dev** (for tray icon support)
- **AppIndicator extension enabled** (for GNOME/Pop!_OS users)

### Install dependencies (Ubuntu/Pop!_OS/Debian):

```bash
sudo apt update
sudo apt install xclip zenity libnotify-bin libayatana-appindicator3-dev gnome-shell-extension-appindicator
```

---

## Build

```bash
git clone https://github.com/caioqf/whipr.git
cd whipr
go build -o whipr
```

---

## Usage

```bash
./whipr
```

- The tray icon should appear (make sure AppIndicator extension is enabled in GNOME).
- Select any text with your mouse.
- Click the tray icon and choose **"Translate selected text"**.
- The selected text will appear in a popup or notification, depending on your menu settings.

---

## Troubleshooting

- **No tray icon?**
  - Make sure the AppIndicator extension is enabled in GNOME Tweaks or via `gnome-extensions`
  - Restart GNOME Shell (`Alt+F2`, type `r`, press Enter).
- **No notification?**
  - Ensure `libnotify-bin` is installed.
- **Popup/notification not showing?**
  - Test manually:  
    `xclip -o -selection primary`  
    `zenity --info --text="test"`  
    `notify-send "Test" "Hello"`

---

## Customization

- You can configure whether the translation appears as a popup or notification via the tray menu checkboxes.

---

## License

MIT
