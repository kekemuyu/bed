package core

// Event represents the event emitted by UI.
type Event int

// Events
const (
	Nop = iota
	Quit
	CursorUp
	CursorDown
	CursorLeft
	CursorRight
	CursorPrev
	CursorNext
	CursorHead
	CursorEnd
	ScrollUp
	ScrollDown
	PageUp
	PageDown
	PageTop
	PageEnd
)
