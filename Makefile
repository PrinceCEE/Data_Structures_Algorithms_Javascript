test:
	@clear
	@echo "Running ${test_name}"
	go test -run ${test_name} ./...