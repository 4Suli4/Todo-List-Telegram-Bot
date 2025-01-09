package main

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func Init(t *testing.T) {
	errGotEnv := godotenv.Load()
	if errGotEnv != nil {
		t.Fatalf("Error loading .env file: %v", errGotEnv)
	}
}

func TestLoadEnv(t *testing.T) {

	t.Run("Init", Init)

	testKey := os.Getenv("TEST_KEY")
	if testKey == "" {
		t.Errorf("TEST_KEY env variable not set %s", testKey)
	}
}

func TestLoadEnvFailed(t *testing.T) {

	t.Run("Init", Init)

	testKey := os.Getenv("MISSING_KEY")
	if testKey != "" {
		t.Fatalf("TEST_KEY env variable not set %s", testKey)
	}
}
