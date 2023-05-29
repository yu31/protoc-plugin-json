# Copyright (C) 2021 Yu.

VERBOSE = no
CASE = ""
DIR = ""

.PHONY: help
help: ## help for command
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_%-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: format
format: ## Execute go fmt ./...
	@[ ${VERBOSE} = "yes" ] && set -x; go fmt ./...;

.PHONY: vet
vet: ## Execute go vet ./...
	@[ ${VERBOSE} = "yes" ] && set -x; go vet ./...;

.PHONY: lint
lint: ## Execute staticcheck ./...
	@[ ${VERBOSE} = "yes" ] && set -x; staticcheck ./...;

.PHONY: tidy
tidy: ## Execute go mod tidy
	@[ ${VERBOSE} = "yes" ] && set -x; go mod tidy;

.PHONY: check
check: ## Execute tidy format vet lint
check: tidy format vet lint

#.PHONY: benchmark
#benchmark:
#	@[[ ${VERBOSE} = "yes" ]] && set -x; go test -test.benchmark="." -test.run="Benchmark" -benchmem -count=1 ./...;

.PHONY: clean
clean: ## Delete build directory
	@[[ ${VERBOSE} = "yes" ]] && set -x; /bir/rm -fr ./build

.PHONY: generate-java
generate-java: ## Generate Java code by proto file.
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/generate_java.sh || bash scripts/generate_java.sh;

.PHONY: generate-go
generate-go: ## Generate Golang code by proto file.
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/generate_go.sh || bash scripts/generate_go.sh;

.PHONY: generate-py
generate-py: ## Generate Python code by proto file.
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/generate_py.sh || bash scripts/generate_py.sh;

.PHONY: generate
generate: ## Wrapper run generate-go generate-java generate-py.
generate: generate-go generate-java generate-py

.PHONY: compile
compile: ## Compile the cmd in local.
compile: generate-go
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/compile.sh || bash scripts/compile.sh;

.PHONY: install
install: ## Copy the cmd library to $GOPATH/bin
install: compile
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/install.sh || bash scripts/install.sh;

.PHONY: generate-test
generate-test: ## Generated the test code by test proto file.
generate-test: compile
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/generate_test_go.sh || bash scripts/generate_test_go.sh;

.PHONY: test
test: ## Run the test cases that in ./xgo/tests/cases
test: generate-test tidy format vet
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/run_test_go.sh "${CASE}" "${DIR}" || bash scripts/run_test_go.sh "${CASE}" "${DIR}"
	@#[[ ${VERBOSE} = "yes" ]] && set -x; go test -v -test.count=1 -failfast -test.run="${CASE}" ./xgo/tests/cases/...;

.PHONY: test-only
test-only:  ## Run test cases without generated codes.
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/run_test_go.sh "${CASE}" "${DIR}" || bash scripts/run_test_go.sh "${CASE}" "${DIR}"
#	@[[ ${VERBOSE} = "yes" ]] && set -x; go test -v -test.count=1 -failfast -test.run="${CASE}" ./xgo/tests/cases/...;

.PHONY: bench
bench: generate-test check
	@[[ ${VERBOSE} = "yes" ]] && set -x; go test -test.bench="." -test.run="Benchmark" -benchmem -count=1 -test.benchtime=1s -test.parallel=8 ./xgo/tests/benchmark/...;

.PHONY: bench-only
bench-only:
	@[[ ${VERBOSE} = "yes" ]] && set -x; go test -test.bench="." -test.run="Benchmark" -benchmem -count=1 -test.benchtime=1s -test.parallel=8 ./xgo/tests/benchmark/...;

.PHONY: test-error-go
test-error-go:   ## Run test cases for invalid parameters in tags for golang.
test-error-go: compile
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/test_error_go.sh "${CASE}" || bash scripts/test_error_go.sh "${CASE}"

.PHONY: test-error
test-error:  ## Run test cases for invalid parameters in tags.
test-error: test-error-go

# publishing java jar to central repository
.PHONY: java-release
java-release:
	@[[ ${VERBOSE} = "yes" ]] && set -x; cd xjava; mvn clean deploy -P release

.DEFAULT_GOAL = help

# Target name % means that it is a rule that matches anything, @: is a recipe;
# the : means do nothing
%:
	@:

