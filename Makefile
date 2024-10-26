.PHONY: test
test: ; $(info $(M) start unit testing...) @
	@go test $$(go list ./... | grep -v /mocks/) -gcflags=all=-l --race -v -short -coverprofile=profile.cov
	@echo "\n*****************************"
	@totalCoverage=$$(go tool cover -func profile.cov | awk '/total:/ {print substr($$3, 1, length($$3)-1)}'); \
	echo "**  TOTAL COVERAGE: $$totalCoverage%  **"; \
	echo "*****************************\n";
