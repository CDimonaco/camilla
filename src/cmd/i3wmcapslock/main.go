package main

import (
	"fmt"
	"os"

	"github.com/cdimonaco/i3wmcapslock/src/i3bar"
	"github.com/cdimonaco/i3wmcapslock/src/led"
)

func main() {
	i3Client, err := i3bar.NewClient(
		os.Stdin,
		os.Stdout,
	)
	if err != nil {
		fmt.Printf("Fatal error %s", err.Error())
	}
	blockChan := i3Client.Start()
	for blocks := range blockChan {
		mask, err := led.Status()
		// Check caps status
		capsLock := led.CapsLock()
		capsBlock := i3bar.Block{
			FullText: "Caps",
			Name:     "capslock",
			Color:    capsLock.Color().Inactive,
		}
		if mask&capsLock.Mask() != 0 {
			capsBlock = i3bar.Block{
				FullText: "Caps",
				Name:     "capslock",
				Color:    capsLock.Color().Active,
				Urgent:   true,
			}
		}
		if err != nil {
			panic(err.Error())
		}
		i3Client.AppendBlock(blocks, &capsBlock, true)
	}
}
