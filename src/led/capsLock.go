package led

type capsLock struct {
	mask Mask
}

func (c capsLock) Mask() Mask {
	return c.mask
}

func (c capsLock) Label() string {
	return "Caps"
}

func (c capsLock) Color() Color {
	return Color{
		Active:   "#ff0000",
		Inactive: "#00",
	}
}

// CapsLock return the definition of CapsLock led
func CapsLock() Led {
	return capsLock{
		mask: 0x1,
	}
}
