# Load .env file as Makefile environment variables
include .env
export $(shell sed 's/=.*//' .env)

run:
	go run ./cmd/api