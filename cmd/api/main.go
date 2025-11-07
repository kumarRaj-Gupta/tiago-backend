package main

import (
	"log"
	"tiago/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8880"),
	}

	app := application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
