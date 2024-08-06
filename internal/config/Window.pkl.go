// Code generated from Pkl module `remychantenay.pkl_playground.internal.config`. DO NOT EDIT.
package config

import "github.com/remychantenay/pkl-playground/internal/config/kind"

type Window struct {
	// Width is the width of the entire window.
	Width float64 `pkl:"width"`

	// Height is the eight of the entire window.
	Height float64 `pkl:"height"`

	// Elevation is the elevation of the window.
	// Nullable.
	Elevation *int `pkl:"elevation"`

	// Kind is the kind of window.
	Kind kind.Kind `pkl:"kind"`
}
