package main

import (
	"fmt"
	"os"

	"github.com/cdimonaco/i3wmcapslock/src/i3bar"
)

func main() {
	i3Client := i3bar.NewClient(
		os.Stdin,
		os.Stdout,
	)
	blockChan := i3Client.Start()
	for blocks := range blockChan {
		//Blocks represents the last protocol block
		lastBlock := blocks[len(blocks)-1]
		i3Client.AppendBlock(blocks, lastBlock)
		fmt.Print(",")
	}
}
