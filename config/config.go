package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	AWS_REGION string
	AWS_BUCKET string
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "..")
	godotenv.Load(filepath.Join(dir, ".env"))
}
func GetEnv() *Config {
	return &Config{
		AWS_REGION: os.Getenv("AWS_REGION"),
		AWS_BUCKET: os.Getenv("AWS_BUCKET"),
	}
}
