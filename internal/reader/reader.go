package reader

import (
	"fmt"
	"os"
)

func LoadOpenApiSpec(p string) (string, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		return "", fmt.Errorf("cannot read openapi: %w", err)
	}

	return string(b), nil
}
