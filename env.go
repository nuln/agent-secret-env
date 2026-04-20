// Package env implements a SecretProvider that reads secrets from environment variables.
package env

import (
	"context"
	"fmt"
	"os"

	agent "github.com/nuln/agent-core"
)

func init() {
	agent.RegisterPluginConfigSpec(agent.PluginConfigSpec{
		PluginName:  "env",
		PluginType:  "secret",
		Description: "Reads secrets from environment variables",
		Fields:      nil,
	})

	agent.RegisterSecretProvider("env", func(opts map[string]any) (agent.SecretProvider, error) {
		return &EnvSecretProvider{}, nil
	})
}

// EnvSecretProvider implements agent.SecretProvider using os.Getenv.
type EnvSecretProvider struct{}

// Get returns the value of the environment variable named key.
func (p *EnvSecretProvider) Get(_ context.Context, key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("env secret: %q not set", key)
	}
	return val, nil
}
