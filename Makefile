# * Makefile
.SILENT: run signup
# TO_NULL = > /dev/null 2>&1 
run:
	@cd api && go mod init github.com/Stei-ITstudents/go-auth/api; go get -u ./... || true
	@cd api && go mod tidy || true
	@cd api && go run main.go || true
signup:
	@curl -X POST http://localhost:8080/signup -d '{"username":"NewUser", "password":"password"}'