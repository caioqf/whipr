package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type AppSettings struct {
	mu        sync.RWMutex
	UsePopup  bool
	UseNotify bool
}

var settings = &AppSettings{
	UseNotify: false,
	UsePopup:  true,
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
	settings.UseNotify = loaded.UseNotify
	settings.mu.Unlock()
}

func saveSettings() {
	settings.mu.RLock()
	toSave := AppSettings{
		UsePopup:  settings.UsePopup,
		UseNotify: settings.UseNotify,
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
	s.UseNotify = !enabled
	s.mu.Unlock()
	saveSettings()
}

func (s *AppSettings) SetNotify(enabled bool) {
	s.mu.Lock()
	s.UseNotify = enabled
	s.UsePopup = !enabled
	s.mu.Unlock()
	saveSettings()
}

func (s *AppSettings) ShouldUsePopup() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.UsePopup
}

func (s *AppSettings) ShouldUseNotify() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.UseNotify
}

func ShouldUsePopup() bool          { return settings.ShouldUsePopup() }
func ShouldUseNotify() bool         { return settings.ShouldUseNotify() }
func SetPopupEnabled(enabled bool)  { settings.SetPopup(enabled) }
func SetNotifyEnabled(enabled bool) { settings.SetNotify(enabled) }
