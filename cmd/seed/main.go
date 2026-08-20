package main

import (
	"flag"
	"log"
	"time"
)

// cmd/seed is a dev/CI-only tool that seeds two e2e users (operator + normal) and writes their minted,
// short-lived JWTs to a file as env vars. It is a separate binary from cmd/api on purpose: no build
// target references it, so it never ships in the production image.
func main() {
	cfg := flag.String("config", "", "path to env configuration file")
	opEmail := flag.String("operator-email", "", "operator email (auto-registered if missing)")
	usrEmail := flag.String("user-email", "", "normal user email (auto-registered if missing)")
	out := flag.String("out", "e2e-tokens.env", "output file path for the generated env vars")
	ttl := flag.Duration("ttl", 24*time.Hour, "token validity duration (short-lived; regenerated each CI run)")
	flag.Parse()

	if *opEmail == "" || *usrEmail == "" {
		log.Fatal("both --operator-email and --user-email are required")
	}

	if err := seed(*cfg, *opEmail, *usrEmail, *out, *ttl); err != nil {
		log.Fatalf("seed: %v", err)
	}
}
