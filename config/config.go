package config

import (
    "os"
)

type Config struct {
    SMPPHost     string
    SMPPPort     string
    SMPPUser     string
    SMPPPassword string
    Sender       string
}

func LoadConfig() *Config {
    return &Config{
        SMPPHost:     os.Getenv("SMPP_HOST"),
        SMPPPort:     os.Getenv("SMPP_PORT"),
        SMPPUser:     os.Getenv("SMPP_USER"),
        SMPPPassword: os.Getenv("SMPP_PASS"),
        Sender:       os.Getenv("SMPP_SENDER"),
    }
}
