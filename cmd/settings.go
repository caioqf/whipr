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
	usePopup  bool
	useNotify bool
}

var settings = &AppSettings{
	useNotify: true,
	usePopup:  false,
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
		if !os.IsNotExist(err) {
			log.Printf("Failed to read settings: %v", err)
		}
		return
	}
	var loaded AppSettings
	if err := json.Unmarshal(data, &loaded); err != nil {
		log.Printf("Failed to unmarshal settings: %v", err)
		return
	}
	settings.mu.Lock()
	settings.usePopup = loaded.usePopup
	settings.useNotify = loaded.useNotify
	settings.mu.Unlock()
}

func saveSettings() {
	settings.mu.RLock()
	toSave := AppSettings{
		usePopup:  settings.usePopup,
		useNotify: settings.useNotify,
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
	log.Println("setting whipr to use popup")
	s.mu.Lock()
	s.usePopup = enabled
	s.useNotify = !enabled
	s.mu.Unlock()
	saveSettings()
}

func (s *AppSettings) SetNotify(enabled bool) {
	log.Println("setting whipr to use notification")
	s.mu.Lock()
	s.useNotify = enabled
	s.usePopup = !enabled
	s.mu.Unlock()
	saveSettings()
}

func (s *AppSettings) ShouldUsePopup() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.usePopup
}

func (s *AppSettings) ShouldUseNotify() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.useNotify
}

func ShouldUsePopup() bool          { return settings.ShouldUsePopup() }
func ShouldUseNotify() bool         { return settings.ShouldUseNotify() }
func SetPopupEnabled(enabled bool)  { settings.SetPopup(enabled) }
func SetNotifyEnabled(enabled bool) { settings.SetNotify(enabled) }
