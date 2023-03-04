package models

import "time"

type Song struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	Playing     bool          `json:"playing"`
}
