package utils

import (
	"fmt"
	"log"
	"strings"
)

func MustNotError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func MergeErrors(errors []error) error {
	var messages []string = make([]string, 0)
	for _, err := range errors {
		if err != nil {
			messages = append(messages, err.Error())
		}
	}

	if len(messages) > 0 {
		return fmt.Errorf("%s", strings.Join(messages, ", "))
	}

	return nil
}
