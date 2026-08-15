package main

import (
	"fmt"

	"metrics/internal/config"
	"metrics/internal/service"
	"metrics/internal/store"
)

func main() {
	cfg := config.Load()
	st := store.New()
	svc := service.New(st, cfg)
	_ = svc
	fmt.Println("metrics ready")
}
