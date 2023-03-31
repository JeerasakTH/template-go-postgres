migrate:
	migrate -source file://migration \
			-database postgres://root:secret@localhost:5432/root?sslmode=disable up

.PHONY: migrate