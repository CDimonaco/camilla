package led

// Mask represet the bitmask of led
type Mask int8

// Color represents the colors
// of an led active and inactive
type Color struct {
	Active   string
	Inactive string
}

// Led is the interface for leds
type Led interface {
	Mask() Mask
	Label() string
	Color() Color
}
