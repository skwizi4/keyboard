package keyboard

import (
	"bufio"
	"os"

	"strings"
)

func GetString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "error", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}
