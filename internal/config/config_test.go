package config

import "testing"

func TestLoad(t *testing.T) {
	t.Setenv("METRICS_WORKERS", "")
	t.Setenv("METRICS_BUCKET_SIZE", "")
	c := Load()
	if c.Workers != 2 || c.BucketSize != 2 {
		t.Fatalf("%+v", c)
	}
}
