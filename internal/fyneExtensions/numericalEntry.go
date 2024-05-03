package fyneExtensions

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type NumericalEntry struct {
	widget.Entry
	iData binding.Int
}

func NewNumericalEntryWithData(data binding.Int) *NumericalEntry {
	n, err := data.Get()
	if err != nil {
		panic(err)
	}

	entry := &NumericalEntry{
		iData: data,
		Entry: widget.Entry{
			Text: fmt.Sprint(n),
		},
	}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *NumericalEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)
		e.applyData()
	}
}

func (e *NumericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
	e.applyData()
}

func (e *NumericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

func (e *NumericalEntry) applyData() {
	i, err := strconv.Atoi(e.Entry.Text)
	if err != nil {
		panic(err)
	}

	e.iData.Set(i)
}
