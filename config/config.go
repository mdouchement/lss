package config

import (
	"bytes"
	"os/exec"
	"path/filepath"

	"github.com/caarlos0/env"
)

// Cfg variable holds all LSS configuration.
var Cfg *Config

// Config holds all configuration for our program.
type Config struct {
	Version string `env:"APP_VERSION"`

	Workspace       string `env:"LSS_WORKSPACE" envDefault:"./workspace"`
	UploadSizeLimit string `env:"LSS_UPLOAD_SIZE_LIMIT" envDefault:"8G"`
	RouterNamespace string `env:"LSS_ROUTER_NAMESPACE" envDefault:""`
}

func init() {
	Cfg = &Config{}
	env.Parse(Cfg)

	if Cfg.Version == "" {
		Cfg.Version = version()
	}
	Cfg.Workspace = filepath.Clean(Cfg.Workspace)
}

func version() string {
	var version bytes.Buffer
	tag, _ := exec.Command("sh", "-c", "git describe --abbrev=0 --tags 2> /dev/null").Output()
	branch, _ := exec.Command("git", "symbolic-ref", "-q", "--short", "HEAD").Output()
	commit, _ := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	version.Write(tag)
	version.WriteString("-")
	version.Write(branch)
	version.WriteString("-")
	version.Write(commit)
	return string(bytes.Replace(version.Bytes(), []byte("\n"), []byte{}, -1))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
