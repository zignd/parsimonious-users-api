package utils

import (
	"fmt"
	"os"

	"github.com/zignd/parsimonious-users-api/envs"
)

func ValidateEnvs() error {
	if os.Getenv(envs.PostgresAddr) == "" {
		return fmt.Errorf("%s is required", envs.PostgresAddr)
	}
	if os.Getenv(envs.PostgresUser) == "" {
		return fmt.Errorf("%s is required", envs.PostgresUser)
	}
	if os.Getenv(envs.PostgresPassword) == "" {
		return fmt.Errorf("%s is required", envs.PostgresPassword)
	}
	if os.Getenv(envs.PostgresPassword) == "" {
		return fmt.Errorf("%s is required", envs.PostgresPassword)
	}
	if os.Getenv(envs.HTTPServerPort) == "" {
		return fmt.Errorf("%s is required", envs.HTTPServerPort)
	}

	return nil
}
