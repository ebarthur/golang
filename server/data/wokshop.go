package data

import "time"

type Workshop struct {
	Course Course
	Date   time.Time
}

// Embedding
type WorkshopEmbed struct {
	Course
	Date time.Time
}
