# create migration file
#	migrate create -ext sql -dir db/migrations -seq create_roles_table

include .env
export

run:
	go run cmd/main.go
