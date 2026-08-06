package config

import (
	"log/slog"
	"testing"

	"github.com/abssh/auth-service/internal/logger"
	"github.com/abssh/auth-service/internal/testutil/assert"
)

type Env struct {
	EnvName  string
	EnvValue string
}

func TestLoadConfigs(t *testing.T) {
	setValidEnv := func() {
		t.Setenv("HTTP_PORT", "7000")
		t.Setenv("HTTP_HOST", "localhost")
		t.Setenv("GRPC_PORT", "7001")
		t.Setenv("GRPC_HOST", "localhost2")
		t.Setenv("LOG_LEVEL", "Warn")
	}

	t.Run("load valid configs", func(t *testing.T) {
		setValidEnv()

		cfg := &Config{}
		assert.AssertError(t, cfg.Load()).NoError()

		assert.Assert(t, cfg.HttpPort).IsEqualTo(7000)
		assert.Assert(t, cfg.HttpHost).IsEqualTo("localhost")
		assert.Assert(t, cfg.GrpcPort).IsEqualTo(7001)
		assert.Assert(t, cfg.GrpcHost).IsEqualTo("localhost2")
		assert.Assert(t, cfg.LogLevel).IsEqualTo(logger.LogLevel(slog.LevelWarn))
	})

	t.Run("load invalid configs", func(t *testing.T) {
		BadEnvValues := map[string]Env{
			"invalid http port value": {EnvName: "HTTP_PORT", EnvValue: "abc"},
			"invalid grpc port value": {EnvName: "GRPC_PORT", EnvValue: "def"},
			"invalid log level":  {EnvName: "LOG_LEVEL", EnvValue: "the bug"},
		}

		for msg, env := range BadEnvValues {
			t.Run(msg, func(t *testing.T) {
				setValidEnv()

				t.Setenv(env.EnvName, env.EnvValue)

				cfg := &Config{}
				err := cfg.Load()
				assert.AssertError(t, err).Error()
			})
		}
	})

}

func TestLoadConfigDefaults(t *testing.T) {
	cfg := &Config{}
	assert.AssertError(t, cfg.Load()).NoError()

	assert.Assert(t, cfg.HttpPort).IsEqualTo(8080)
	assert.Assert(t, cfg.HttpHost).IsEqualTo("")
	assert.Assert(t, cfg.GrpcPort).IsEqualTo(9090)
	assert.Assert(t, cfg.GrpcHost).IsEqualTo("")
	assert.Assert(t, cfg.LogLevel).IsEqualTo(logger.LogLevel(slog.LevelInfo))

}
