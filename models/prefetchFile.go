package models

import "time"

type PrefetchFile struct {
	Name           string
	LastModification time.Time
}