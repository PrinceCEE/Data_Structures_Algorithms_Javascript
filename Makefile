test:
	@clear
	@echo "Running ${test_name}"
	go test -v -run ${test_name} ./...