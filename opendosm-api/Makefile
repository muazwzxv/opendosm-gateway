include .env

migrations.new:
	goose create $(name) sql

up:
	goose ${GOOSE_DRIVER} ${GOOSE_DBSTRING} up

down:
	goose ${GOOSE_DRIVER} ${GOOSE_DBSTRING} down