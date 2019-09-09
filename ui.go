package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type UIController interface {
	Render()
	OnResize(ui.Resize)
	UpdateData()
}

