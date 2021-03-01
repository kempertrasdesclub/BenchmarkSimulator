package data

import "sync"

type DataCache struct {
	UserId         string `json:"user_id"`
	Status         string `json:"status"`
	Manual         bool   `json:"manual"`
	LastActivityAt int64  `json:"last_activity_at"`
	ActiveChannel  string `json:"-" db:"-"`
}

type Cache struct {
	Cache map[string]DataCache
	m     sync.RWMutex
}
