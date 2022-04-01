package main

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln("could not load env", err)
	}
}
