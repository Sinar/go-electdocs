run:
	@cd cmd/go-electdocs && go run main.go

test:
	@cd ./internal/pdf && gotest ./...

work:
	@go work use -r .
	@go list -m -json all
