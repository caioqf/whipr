package cmd

import "sync"

type AppSettings struct {
	mu        sync.RWMutex
	usePopup  bool
	useNotify bool
}

var settings = &AppSettings{
	useNotify: true,
}

func (s *AppSettings) SetPopup(enabled bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.usePopup = enabled
	s.useNotify = !enabled
}

func (s *AppSettings) SetNotify(enabled bool) {
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

func ShouldUsePopup() bool  { return settings.ShouldUsePopup() }
func ShouldUseNotify() bool { return settings.ShouldUseNotify() }
