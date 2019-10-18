package i3bar

import (
	"encoding/json"
	"os"
	"syscall"
)

// AppendBlock will append a valid i3 protocol block
// to the json encoder output
func AppendBlock(output *json.Encoder, currentBlocks []*Block, newBlock *Block) error {
	// append the new block to the current blocks
	newBlocks := append(currentBlocks, newBlock)
	// Encode the new blocks on the ouput encoder
	return output.Encode(newBlocks)
}

// ImmediateUpdate sends a signal to i3status process
// in order to update immediately the bar regardless the
// protocol time
func ImmediateUpdate(process *os.Process) error {
	return process.Signal(syscall.SIGUSR1)
}
