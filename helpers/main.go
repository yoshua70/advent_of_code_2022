package helpers

import (
	"errors"
	"fmt"
	"os"
)

func GetInput(file_name string) (*os.File, error) {
	input, err := os.Open(fmt.Sprintf("../resources/%s", file_name))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while opening file %s", file_name))
	}

	return input, nil
}
