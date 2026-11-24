-include .env
export

NAME := $(filter-out migrate-create,$(MAKECMDGOALS))

.PHONY: migrate-up migrate-down migrate-status migrate-create

migrate-up:
	goose -dir migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DATABASE_URL)" down

migrate-status:
	goose -dir migrations postgres "$(DATABASE_URL)" status

migrate-create:
	@test -n "$(NAME)" || (echo "usage: make migrate-create <name>"; exit 1)
	goose -dir migrations -s create $(NAME) sql

# swallow the extra word so make doesn't treat it as a missing target
%:
	@: