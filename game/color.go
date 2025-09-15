package game

type Color int

const (
	Empty Color = iota
	Black
	White
	Wall
	None
)

func ColorToStr(c Color) string {
	switch c {
	case Black:
		return "○"
	case White:
		return "◉"
	case Empty:
		return " "
	default:
		panic("unhandled default case")
	}
	return ""
}

func OpponentColor(me Color) Color {
	switch me {
	case Black:
		return White
	case White:
		return Black
	default:
		panic("unhandled default case")
	}
	panic("invalid state")
}
