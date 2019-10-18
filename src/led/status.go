package led

import (
	"os/exec"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

// Status returns led status
// it simply return the bitmask of leds
// is the output of the command
// xset q | grep "LED mask" with a regex
// applied on the output
func Status() (Mask, error) {
	ledMaskRegexp := regexp.MustCompile("[0-9]{8}")
	output, err := exec.Command("bash", "-c", "xset -q | grep \"LED mask\"").Output()
	if err != nil {
		return 0, errors.Wrap(err, "Error during led mask decoding")
	}
	outputString := string(output)
	stringMask := ledMaskRegexp.FindString(outputString)
	mask, err := strconv.ParseInt(stringMask, 2, 8)
	if err != nil {
		return 0, errors.Wrap(err, "Error during int parsing of string ledmask")
	}
	return Mask(mask), nil
}
