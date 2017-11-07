package uigo

import "fmt"

const (
	// Reset will reset any modifies
	Reset = "0m"

	// BrightBlue is a bright, obvious blue
	BrightBlue = "38;5;45m"

	// BrightGreen is a bright, obvious green
	BrightGreen = "38;5;156m"

	// Green is a obvious green
	Green = "38;5;156m"

	// BrightOrange is a bright, obvious Orange
	BrightOrange = "38;5;214m"

	// Orange is a bright, obvious Orange
	Orange = "38;5;208m"

	// Pink is a bright, obvious Pink
	Pink = "38;5;198m"

	// Purple is a bright, obvious Purple
	Purple = "38;5;135m"

	// Grey is a light, but not white
	Grey = "38;5;249m"
)

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

var (
	// DefaultSep is the default seperator for the Wrap methods
	DefaultSep = "[]"
	// DefaultSepColor is the default color for seperators
	DefaultSepColor = Grey
	// DefaultPayloadColor is the default payload color
	DefaultPayloadColor = BrightGreen
)

// UI creates the ui object, which is just a string
type UI struct {
	Header  string
	Payload string
	Output  string
}

// SetSeperator will change the seperator characters
func SetSeperator(sep string) {
	DefaultSep = sep
}

// New instantiates a new UI object
func New(header string) *UI {
	return &UI{
		Header: header,
	}
}

// Print the message
func (ui *UI) Print(message string) {
	fmt.Printf("%s %s", ui.Output, message)
}

// Wrap will wrapp the ui.Payload with seperator
func (ui *UI) Wrap() *UI {
	ui.Output = fmt.Sprintf("%s%s%s", string(DefaultSep[0]), ui.Header, string(DefaultSep[1]))
	return ui
}

// WrapWith will wrap your Payload with the provided seperator
func (ui *UI) WrapWith(sep string) *UI {
	if len(sep) != 2 {
		sep = DefaultSep
	}
	ui.Output = fmt.Sprintf("%s%s%s", string(sep[0]), ui.Header, string(sep[1]))
	return ui
}
