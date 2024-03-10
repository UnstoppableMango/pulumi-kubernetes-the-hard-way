PACK            := kubernetes-the-hard-way
PROJECT         := github.com/unstoppablemango/pulumi-${PACK}

PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION_PATH    := provider/pkg/version.Version

WORKING_DIR     := $(shell pwd)
SCHEMA_PATH     := ${WORKING_DIR}/schema.yaml

PROVIDER_PKG    := $(shell find provider/pkg -type f)

# TODO: Can we remove?
GOPATH          := $(shell go env GOPATH)

JAVA_GEN        := pulumi-java-gen

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOEXE ?= $(shell go env GOEXE)
LOCAL_PROVIDER_FILENAME := $(PROVIDER)$(GOEXE)
export GOWORK := off
# Only use explicitly installed plugins - this is to avoid using any ambient plugins from the PATH
export PULUMI_IGNORE_AMBIENT_PLUGINS = true

ifeq ($(CI)$(PROVIDER_VERSION),true)
$(error PROVIDER_VERSION must be set in CI environments)
endif

# Input during CI using `make [TARGET] PROVIDER_VERSION=""` or by setting a PROVIDER_VERSION environment variable
# Local builds will just used this fixed default version unless specified
PROVIDER_VERSION ?= 0.0.1-alpha.0+dev
# Ensure the leading "v" is removed - use this normalised version everywhere rather than the raw input to ensure consistency.
# These variables are lazy (no `:`) so they're not calculated until after the dependency is installed
VERSION_GENERIC = $(shell bin/pulumictl convert-version -l generic -v "$(PROVIDER_VERSION)")
VERSION_FLAGS   = -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION_GENERIC}"

# Ensure make directory exists
# For targets which either don't generate a single file output, or the file is committed, we use a "sentinel"
# file within `.make/` to track the staleness of the target and only rebuild when needed. At the end of each
# relevant target we run `@touch $@` to update the file which is the name of the target.
_ := $(shell mkdir -p .make)

generate:: gen_go_sdk gen_dotnet_sdk gen_nodejs_sdk gen_python_sdk

build:: build_provider build_dotnet_sdk build_nodejs_sdk build_python_sdk

install:: install_provider install_dotnet_sdk install_nodejs_sdk


# Provider

#build_provider::
#	rm -rf ${WORKING_DIR}/bin/${PROVIDER}
#	cd provider/cmd/${PROVIDER} && VERSION=${VERSION} SCHEMA=${SCHEMA_PATH} go generate main.go
#	cd provider/cmd/${PROVIDER} && go build -o ${WORKING_DIR}/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" .

install_provider:: build_provider
	cp ${WORKING_DIR}/bin/${PROVIDER} ${GOPATH}/bin


# Go SDK

gen_go_sdk::
	rm -rf sdk/go
	cd provider/cmd/${CODEGEN} && go run . go ../../../sdk/go ${SCHEMA_PATH}


# .NET SDK

gen_dotnet_sdk::
	rm -rf sdk/dotnet
	cd provider/cmd/${CODEGEN} && go run . dotnet ../../../sdk/dotnet ${SCHEMA_PATH}

build_dotnet_sdk:: DOTNET_VERSION := ${VERSION}
build_dotnet_sdk:: gen_dotnet_sdk
	cd sdk/dotnet/ && \
		echo "${DOTNET_VERSION}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

install_dotnet_sdk:: build_dotnet_sdk
	rm -rf ${WORKING_DIR}/nuget
	mkdir -p ${WORKING_DIR}/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;


# Node.js SDK

gen_nodejs_sdk::
	rm -rf sdk/nodejs
	cd provider/cmd/${CODEGEN} && go run . nodejs ../../../sdk/nodejs ${SCHEMA_PATH}

build_nodejs_sdk:: gen_nodejs_sdk
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc --version && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json && \
		rm ./bin/package.json.bak

install_nodejs_sdk:: build_nodejs_sdk
	yarn link --cwd ${WORKING_DIR}/sdk/nodejs/bin


# Python SDK

gen_python_sdk::
	rm -rf sdk/python
	cd provider/cmd/${CODEGEN} && go run . python ../../../sdk/python ${SCHEMA_PATH}
	cp ${WORKING_DIR}/README.md sdk/python

build_python_sdk:: PYPI_VERSION := ${VERSION}
build_python_sdk:: gen_python_sdk
	cd sdk/python/ && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
#		sed -i.bak -e 's/^  version = .*/  version = "$(PYPI_VERSION)"/g' ./bin/pyproject.toml && \
#		rm ./bin/pyproject.toml.bak && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build && \
		cd ./bin && \
		../venv/bin/python -m build .

## Empty build target for Go
build_go_sdk::

.PHONY: upgrade_tools upgrade_java upgrade_pulumi upgrade_pulumictl upgrade_schematools
upgrade_tools: upgrade_java upgrade_pulumi upgrade_pulumictl upgrade_schematools
upgrade_java:
	gh release list --repo pulumi/pulumi-java --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 > .pulumi-java-gen.version
upgrade_pulumi:
	gh release list --repo pulumi/pulumi --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .pulumi.version
upgrade_pulumictl:
	gh release list --repo pulumi/pulumictl --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .pulumictl.version
upgrade_schematools:
	gh release list --repo pulumi/schema-tools --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .schema-tools.version


# --------- File-based targets --------- #

.pulumi/bin/pulumi: PULUMI_VERSION := $(shell cat .pulumi.version)
.pulumi/bin/pulumi: HOME := $(WORKING_DIR)
.pulumi/bin/pulumi: .pulumi.version
	curl -fsSL https://get.pulumi.com | sh -s -- --version "$(PULUMI_VERSION)"

# Download local copy of pulumictl based on the version in .pulumictl.version
# Anywhere which uses VERSION_GENERIC or VERSION_FLAGS should depend on bin/pulumictl
bin/pulumictl: PULUMICTL_VERSION := $(shell cat .pulumictl.version)
bin/pulumictl: PLAT := $(shell go version | sed -En "s/go version go.* (.*)\/(.*)/\1-\2/p")
bin/pulumictl: PULUMICTL_URL := "https://github.com/pulumi/pulumictl/releases/download/v$(PULUMICTL_VERSION)/pulumictl-v$(PULUMICTL_VERSION)-$(PLAT).tar.gz"
bin/pulumictl: .pulumictl.version
	@echo "Installing pulumictl"
	@mkdir -p bin
	wget -q -O - "$(PULUMICTL_URL)" | tar -xzf - -C $(WORKING_DIR)/bin pulumictl
	@touch bin/pulumictl
	@echo "pulumictl" $$(./bin/pulumictl version)

# Download local copy of schema-tools based on the version in .schema-tools.version
bin/schema-tools: SCHEMA_TOOLS_VERSION := $(shell cat .schema-tools.version)
bin/schema-tools: PLAT := $(shell go version | sed -En "s/go version go.* (.*)\/(.*)/\1-\2/p")
bin/schema-tools: SCHEMA_TOOLS_URL := "https://github.com/pulumi/schema-tools/releases/download/v$(SCHEMA_TOOLS_VERSION)/schema-tools-v$(SCHEMA_TOOLS_VERSION)-$(PLAT).tar.gz"
bin/schema-tools: .schema-tools.version
	@echo "Installing schema-tools"
	@mkdir -p bin
	wget -q -O - "$(SCHEMA_TOOLS_URL)" | tar -xzf - -C $(WORKING_DIR)/bin schema-tools
	@touch bin/schema-tools
	@echo "schema-tools" $$(./bin/schema-tools version)

bin/$(CODEGEN): bin/pulumictl .make/provider_mod_download provider/cmd/$(CODEGEN)/* $(PROVIDER_PKG)
	# TODO: Why is this weird
	#cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN)
	cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) ./cmd/$(CODEGEN)

# --------- Sentinel targets --------- #

.make/provider_mod_download: provider/go.mod provider/go.sum
	cd provider && go mod download
	@touch $@
