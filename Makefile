build:
	CGO_ENABLED=0 GOOS=linux go build \
		-a -tags netgo \
		-o migrate-vault-file-backend-to-consul
