package i3bar

import (
	"encoding/json"
	"fmt"
	"os"
)

// Client represents a client for interacting with i3bar
// through the i3bar protocol, uses the function
// exported in i3 package and combine them in order to obtain
// an easy interaction with i3bar
type Client struct {
	input         *os.File
	output        *os.File
	inputDecoder  *json.Decoder
	outputEncoder *json.Encoder
	blocks        chan []*Block
}

// Start configure the client, initialize the
// output channel and start the main loop
func (c *Client) Start() chan []*Block {
	err := c.initializeCom()
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Error during the initialize of communication protocol - %s",
			err.Error(),
		)
		os.Exit(1)
	}
	// The main loop
	go func() {
		for c.inputDecoder.More() {
			blocks, err := ParseBlocks(c.inputDecoder)
			if err != nil {
				fmt.Fprintf(
					os.Stderr,
					"Error during block parsing - %s",
					err.Error(),
				)
			}
			c.blocks <- blocks
		}
		close(c.blocks)
	}()
	return c.blocks
}

// AppendBlock append a new block to the block passed
// and outputs them in the encoder of i3bar protocol
func (c *Client) AppendBlock(currentBlocks []*Block, newBlock *Block) {
	AppendBlock(c.outputEncoder, currentBlocks, newBlock)
}

// initializeCom perform the first one time  operations in order
// to initiate a communication with i3bar protocol
func (c *Client) initializeCom() error {
	c.blocks = make(chan []*Block)
	header, token, err := ParseHeader(c.inputDecoder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during header parsing %s", err.Error())
		os.Exit(1)
	}
	tokenString := token.(json.Delim).String()
	// Validate token, we expect a [ delimiter
	if !validToken(tokenString) {
		fmt.Fprintf(os.Stderr, "Unexpected token found %s", token)
		os.Exit(1)
	}
	// Encode the header
	err = c.outputEncoder.Encode(header)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during header encoding %s", err.Error())
		os.Exit(1)
	}
	// Write on ouput the token
	c.output.WriteString(tokenString + "\n")
	return nil
}

func validToken(token string) bool {
	return token == "["
}

func NewClient(input *os.File, output *os.File) *Client {
	inputDecoder := json.NewDecoder(input)
	outputEncoder := json.NewEncoder(output)
	return &Client{
		inputDecoder:  inputDecoder,
		outputEncoder: outputEncoder,
		output:        output,
		input:         input,
	}
}
