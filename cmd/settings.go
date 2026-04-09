package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type AppSettings struct {
	mu             sync.RWMutex
	UsePopup       bool               `json:"use_popup"`
	UseNotification      bool               `json:"use_notification"`
	DetectLanguage bool               `json:"detect_language"`

	APIProvider      string            `json:"api_provider"`
	APIKey           string            `json:"api_key"`
	Model            string            `json:"model"`

	QuickLangs []string          `json:"quick_langs"`
	Shortcut string             `json:"shortcut"`
}


var settings = &AppSettings{
	UsePopup:  false,
	UseNotification: true,
	DetectLanguage: true,
	
	APIProvider:    "openai",
	APIKey:         "",
	Model:          "gpt-3.5-turbo",
	
	QuickLangs:     []string{"en", "pt", "es", "fr", "jp", "ch"},
	Shortcut:       "ctrl+shift+t",
}

func LoadSettings() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Failed to get home dir: %v", err)
		return
	}
	configPath := filepath.Join(home, ".config", "whipr", "settings.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			saveSettings()
			return
		}
		log.Printf("Failed to read settings: %v", err)
		return
	}
	var loaded AppSettings
	if err := json.Unmarshal(data, &loaded); err != nil {
		log.Printf("Failed to unmarshal settings: %v", err)
		return
	}
	settings.mu.Lock()
	settings.UsePopup = loaded.UsePopup
	settings.UseNotification = loaded.UseNotification
	settings.mu.Unlock()
}

func saveSettings() {
	settings.mu.RLock()
	toSave := AppSettings{
		UsePopup:  settings.UsePopup,
		UseNotification: settings.UseNotification,
	}
	settings.mu.RUnlock()
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Failed to get home dir: %v", err)
		return
	}
	configDir := filepath.Join(home, ".config", "whipr")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Printf("Failed to create config dir: %v", err)
		return
	}
	configPath := filepath.Join(configDir, "settings.json")
	data, err := json.Marshal(toSave)
	if err != nil {
		log.Printf("Failed to marshal settings: %v", err)
		return
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		log.Printf("Failed to write settings: %v", err)
	}
}

func (s *AppSettings) SetPopup(enabled bool) {
	s.mu.Lock()
	s.UsePopup = enabled
	s.UseNotification = !enabled
	s.mu.Unlock()
	saveSettings()
}

func (s *AppSettings) SetNotification(enabled bool) {
	s.mu.Lock()
	s.UseNotification = enabled
	s.UsePopup = !enabled
	s.mu.Unlock()
	saveSettings()
}

func (s *AppSettings) ShouldUsePopup() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.UsePopup
}

func (s *AppSettings) ShouldUseNotification() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.UseNotification
}

func ShouldUsePopup() bool          { return settings.ShouldUsePopup() }
func ShouldUseNotification() bool         { return settings.ShouldUseNotification() }
func SetPopupEnabled(enabled bool)  { settings.SetPopup(enabled) }
func SetNotificationEnabled(enabled bool) { settings.SetNotification(enabled) }
