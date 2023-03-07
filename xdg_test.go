package xdg_test

import (
	"os"
	"testing"

	"github.com/lucasepe/xdg"
	"github.com/stretchr/testify/require"
)

type envSample struct {
	name     string
	value    string
	expected string
	actual   func() string
}

func testDirs(t *testing.T, samples ...*envSample) {
	t.Logf("Home: %s", xdg.HomeDir())

	for _, sample := range samples {
		require.NoError(t, os.Setenv(sample.name, sample.value))
		t.Logf("%s: %v", sample.name, os.Getenv(sample.name))
		require.Equal(t, sample.expected, sample.actual())
		t.Logf("%s: %v", sample.name, sample.actual())
	}
}
