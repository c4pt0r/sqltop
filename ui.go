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

type TextGrid struct {
	grid *ui.Grid
	par  *widgets.Paragraph
	x, y int // top, left
}

func newTextGrid(x, y int) *TextGrid {
	termWidth, termHeight := ui.TerminalDimensions()

	par := widgets.NewParagraph()
	par.Border = false

	grid := ui.NewGrid()
	grid.SetRect(x, y, termWidth, termHeight)
	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0, par),
		),
	)
	return &TextGrid{
		grid: grid,
		par:  par,
		x:    x,
		y:    y,
	}
}

// it's caller's duty to be threaded safe
func (g *TextGrid) SetText(str string) {
	g.par.Text = str
}

func (g *TextGrid) OnResize(payload ui.Resize) {
	g.grid.SetRect(g.x, g.y, payload.Width, payload.Height-15)
	ui.Render(g.grid)
}

func (g *TextGrid) Render() {
	ui.Render(g.grid)
}
