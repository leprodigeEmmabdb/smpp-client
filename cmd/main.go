package main

import (
    "log"
    "github.com/tonpseudo/smpp-client/internal/smppclient"
)

func main() {
    if err := smppclient.Connect(); err != nil {
        log.Fatalf("Erreur SMPP: %v", err)
    }
}
