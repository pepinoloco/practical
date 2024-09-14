package config

import (
	"practical/appendix-b/pkgcli/telemetry"
	"github.com/rs/zerolog"
)

type PkgCliConfig struct {
	Logger  zerolog.Logger
	Metrics telemetry.MetricReporter
	Tracer  telemetry.TraceReporter
	AuthConfig
}

type AuthConfig struct {
	Token string `yaml:"auth_token"`
}
