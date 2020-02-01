package ui

import "github.com/gdamore/tcell/views"

type Ui struct {
	view  views.View
	panel views.Widget
	views.WidgetWatchers
}
