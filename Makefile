TEST?=./...

include ./.env.development.local

provider:
	@echo "--- Running Provider Pact tests "
	go test -count=1 -tags=integration ./test/contract/api -run "TestPactProvider"

.PHONY: provider