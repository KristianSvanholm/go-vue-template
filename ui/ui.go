package ui

import "embed"

// The following is an important go directive. Do not remove
//go:embed all:dist
var StaticFiles embed.FS
