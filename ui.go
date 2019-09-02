package main

import (
	"sync"

	uiv3 "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	ui "gopkg.in/gizak/termui.v1"
)

const (
	gMaxPointNum = 51
)

type hotTableWidget struct {
	title  string
	widget *widgets.SparklineGroup
	x, y   int // top-left
	data   []float64

	mu sync.RWMutex
}

func newHotTableWidget(title string, x, y int) *hotTableWidget {
	ret := &hotTableWidget{
		title: title,
		x:     x,
		y:     y,
	}
	sl := widgets.NewSparkline()
	sl.Data = ret.data
	sl.LineColor = uiv3.ColorGreen

	slg := widgets.NewSparklineGroup(sl)
	slg.Title = ret.title

	ret.widget = slg
	return ret
}

func (w *hotTableWidget) pushDataPoint(d float64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.data = append(w.data, d)
	if len(w.data) > gMaxPointNum {
		w.data = w.data[1:]
	}
	w.widget.Sparklines[0].Data = w.data
}

func (w *hotTableWidget) render() {
	w.mu.RLock()
	defer w.mu.RUnlock()

	w.widget.SetRect(w.x, w.y, w.x+ui.TermWidth()/5, 5)
	uiv3.Render(w.widget)
}
