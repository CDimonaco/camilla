package i3bar

// Block type represent a block in the i3bar json protocol
// more on https://i3wm.org/docs/i3bar-protocol.html
type Block struct {
	FullText            string `json:"full_text,omitempty"`
	ShortText           string `json:"short_text,omitempty"`
	Color               string `json:"color,omitempty"`
	Background          string `json:"background,omitempty"`
	Border              string `json:"border,omitempty"`
	BorderTop           string `json:"border_top,omitempty"`
	BorderRight         string `json:"border_right,omitempty"`
	BorderBottom        string `json:"border_bottom,omitempty"`
	BorderLeft          string `json:"border_left,omitempty"`
	MinWidth            int    `json:"min_width,omitempty"`
	Align               string `json:"align,omitempty"`
	Name                string `json:"name,omitempty"`
	Instance            string `json:"instance,omitempty"`
	Urgent              bool   `json:"urgent,omitempty"`
	Separator           bool   `json:"separator,omitempty"`
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"`
	Markup              string `json:"markup,omitempty"`
}

// Header type represents the header of i3 protocol
// is the first part of the protocol before the blocks
type Header struct {
	Version     int  `json:"version,omitempty"`
	StopSignal  int  `json:"stop_signal,omitempty"`
	ContSignal  int  `json:"cont_signal,omitempty"`
	ClickEvents bool `json:"click_events,omitempty"`
}
