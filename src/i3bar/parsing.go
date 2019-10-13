package i3bar

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ParseHeader parse the first bytes of i3bar protocol
// it decodes the header and the first token [ of the
// protocol, after we will have only blocks
func ParseHeader(input *json.Decoder) (*Header, json.Token, error) {
	var header *Header
	// First parse the header
	err := input.Decode(&header)
	if err != nil {
		return nil, nil, errors.Wrap(
			err,
			"Error during the header parsing, aborting input parsing",
		)
	}
	t, err := input.Token()
	if err != nil {
		return nil, nil, errors.Wrap(
			err,
			"Error during infinity array token parsing, aborting",
		)
	}
	return header, t, nil
}

// ParseBlocks accepts a json decoder and try to decode
// from the decoder the i3bar blocks, validate the header
// and outputs the blocks
func ParseBlocks(input *json.Decoder) ([]*Block, error) {
	var blocks []*Block
	err := input.Decode(&blocks)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"Error during the block parsing, aborting input parsing",
		)
	}

	return blocks, nil
}
