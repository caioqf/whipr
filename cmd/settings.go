package cmd

import (
	"log"
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

func (s *AppSettings) SetPopup(enabled bool) {
	log.Println("setting whipr to use popup")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.usePopup = enabled
	s.useNotify = !enabled
}

func (s *AppSettings) SetNotify(enabled bool) {
	log.Println("setting whipr to use notification")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.useNotify = enabled
	s.usePopup = !enabled
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
