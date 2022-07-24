package main

import (
	"io"
	"time"
)

type Artifact interface {
	Title() string
	Creator() string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Streamer
}

type Video interface {
	Streamer
	Resolution() (x, y int)
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}
