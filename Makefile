CURRENT_DIR = $(shell pwd)

DB_URL := postgres://postgres:nocode@localhost:5432/postgres?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path migrations -database '${DB_URL}' -verbose up

mig-down:
	migrate -path migrations -database '${DB_URL}' -verbose down

mig-force:
	migrate -path migrations -database '${DB_URL}' -verbose force 1

mig-create:
	migrate create -ext sql -dir migrations -seq auth_service_table

swag-gen:
	~/go/bin/swag init -g api/router.go -o api/docs
#   rm -r db/migrations