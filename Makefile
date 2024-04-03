PACK            := kubernetes-the-hard-way
PROJECT         := github.com/unstoppablemango/pulumi-${PACK}

PROVIDER        := pulumi-resource-${PACK}

WORKING_DIR     := $(shell pwd)
SCHEMA_FILE     := ${WORKING_DIR}/schema.yaml
PROVIDER_PKG    := ${WORKING_DIR}/provider/cmd/${PROVIDER}/package.json

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
PROVIDER_VERSION ?= 0.0.1
# Ensure the leading "v" is removed - use this normalised version everywhere rather than the raw input to ensure consistency.
# These variables are lazy (no `:`) so they're not calculated until after the dependency is installed
VERSION_GENERIC = $(shell bin/pulumictl convert-version -l generic -v "$(PROVIDER_VERSION)")
VERSION_FLAGS   = -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION_GENERIC}"

# Ensure make directory exists
# For targets which either don't generate a single file output, or the file is committed, we use a "sentinel"
# file within `.make/` to track the staleness of the target and only rebuild when needed. At the end of each
# relevant target we run `@touch $@` to update the file which is the name of the target.
_ := $(shell mkdir -p .make)

.PHONY: default ensure
default: provider build_sdks
ensure: bin/pulumictl .pulumi/bin/pulumi

# Binaries
.PHONY: provider
provider: bin/$(LOCAL_PROVIDER_FILENAME)

.PHONY: test test_dotnet test_python test_go test_nodejs
# Set TEST_TAGS to override -tags for tests
TEST_TAGS ?= all
# Set TEST_NAME to filter tests by name
TEST_NAME ?=
TEST_RUN =
ifneq ($(TEST_NAME),)
TEST_RUN = -run ^$(TEST_NAME)$$
endif
export PULUMI_LOCAL_NUGET=$(WORKING_DIR)/nuget
test: provider build_dotnet build_python install_sdks bin/gotestfmt
	cd examples && go test -json -v -tags=$(TEST_TAGS) -timeout 2h $(TEST_RUN) | tee /tmp/gotest.log | gotestfmt
test_dotnet: provider build_dotnet install_dotnet_sdk
	cd examples && go test -v -tags=dotnet -timeout 2h $(TEST_RUN)
test_python: provider build_python
	cd examples && go test -v -tags=python -timeout 2h $(TEST_RUN)
test_go: provider
	cd examples && go test -v -tags=go -timeout 2h $(TEST_RUN)
test_nodejs: provider install_nodejs_sdk
	cd examples && go test -v -tags=nodejs -timeout 2h $(TEST_RUN)

.PHONY: install_provider
install_provider: .make/install_provider

.PHONY: generate generate_java generate_nodejs generate_python generate_dotnet generate_go generate_types
generate: generate_java generate_nodejs generate_python generate_dotnet generate_go generate_types
generate_java: .make/generate_java
generate_nodejs: .make/generate_nodejs
generate_python: .make/generate_python
generate_dotnet: .make/generate_dotnet
generate_go: .make/generate_go
generate_types: .make/generate_types

.PHONY: local_generate_code
local_generate_code: generate_java
local_generate_code: generate_nodejs
local_generate_code: generate_python
local_generate_code: generate_dotnet
local_generate_code: generate_go
local_generate_code: generate_types
local_generate: local_generate_code

.PHONY: build only_build build_sdks build_nodejs build_python build_dotnet build_java build_go
build: local_generate provider build_sdks
# Required for the codegen action that runs in pulumi/pulumi
only_build: build
build_sdks: build_nodejs build_dotnet build_python build_go build_java
build_nodejs: .make/build_nodejs
build_python: .make/build_python
build_dotnet: .make/build_dotnet
build_java: .make/build_java
build_go: .make/build_go

.PHONY: install_dotnet_sdk install_python_sdk install_go_sdk install_java_sdk install_nodejs_sdk install_sdks
# Required by CI steps - some can be skipped
install_dotnet_sdk: .make/install_dotnet_sdk
install_python_sdk:
install_go_sdk:
install_java_sdk:
install_nodejs_sdk: .make/install_nodejs_sdk
install_sdks: install_dotnet_sdk install_nodejs_sdk

.PHONY: link link_examples
link: link_examples
link_examples: install_nodejs_sdk
	cd examples/remote-ts && yarn link '@unmango/pulumi-kubernetes-the-hard-way'
	cd examples/simple-ts && yarn link '@unmango/pulumi-kubernetes-the-hard-way'

.PHONY: clean
clean:
	rm -rf nuget
	rm -rf .make
	rm -rf bin
	rm -rf dist
	rm -rf provider/scripts/vendor
	rm -rf sdk/dotnet/{bin,obj}
	rm -rf sdk/nodejs/bin
	rm -rf sdk/python/bin
	rm -rf sdk/java/{.gradle,build}
	@if dotnet nuget list source | grep "$(WORKING_DIR)"; then \
		dotnet nuget remove source "$(WORKING_DIR)" \
	; fi

vendor: provider/scripts/vendor/generate-provider-types.ts provider/scripts/vendor/pulumi-schema.d.ts

.PHONY: upgrade_tools upgrade_java upgrade_pulumi upgrade_pulumictl upgrade_schematools upgrade_yq
upgrade_tools: upgrade_java upgrade_pulumi upgrade_pulumictl upgrade_schematools upgrade_yq
upgrade_java:
	gh release list --repo pulumi/pulumi-java --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 > .pulumi-java-gen.version
upgrade_pulumi:
	gh release list --repo pulumi/pulumi --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .pulumi.version
upgrade_pulumictl:
	gh release list --repo pulumi/pulumictl --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .pulumictl.version
upgrade_schematools:
	gh release list --repo pulumi/schema-tools --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .schema-tools.version
upgrade_yq:
	gh release list --repo mikefarah/yq --exclude-drafts --exclude-pre-releases --limit 1 | cut -f1 | sed 's/^v//' > .yq.version

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

bin/gotestfmt:
	@mkdir -p bin
	GOBIN="${WORKING_DIR}/bin" go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@v2.5.0

bin/yq: YQ_VERSION := $(shell cat .yq.version)
bin/yq: PLAT := $(shell go version | sed -En "s/go version go.* (.*)\/(.*)/\1_\2/p")
bin/yq: BINARY := yq_$(PLAT)
bin/yq: .yq.version
	@mkdir -p bin
	wget https://github.com/mikefarah/yq/releases/download/v$(YQ_VERSION)/${BINARY}.tar.gz -O - |\
		tar xz ./${BINARY} && mv ${BINARY} $(WORKING_DIR)/bin/yq
	@touch bin/yq
	@./bin/yq --version

bin/$(LOCAL_PROVIDER_FILENAME): bin/pulumictl .make/provider_mod_download provider/cmd/${PROVIDER}/schema.json provider/cmd/$(PROVIDER)/*.ts $(PROVIDER_PKG)
	cd provider/cmd/${PROVIDER}/ && \
		yarn tsc && \
		cp package.json schema.json ./bin && \
		sed -i.bak -e "s/\$${VERSION}/$(PROVIDER_VERSION)/g" bin/package.json && \
		yarn run pkg . ${PKG_ARGS} --target node16 --output $(WORKING_DIR)/$@

bin/linux-amd64/$(PROVIDER): TARGET := linuxstatic-x64
bin/linux-arm64/$(PROVIDER): TARGET := linuxstatic-arm64
bin/darwin-amd64/$(PROVIDER): TARGET := macos-x64
bin/darwin-arm64/$(PROVIDER): TARGET := macos-arm64
bin/windows-amd64/$(PROVIDER).exe: TARGET := win-x64
bin/%/$(PROVIDER) bin/%/$(PROVIDER).exe: bin/pulumictl .make/provider_mod_download provider/cmd/${PROVIDER}/schema.json provider/cmd/$(PROVIDER)/*.ts $(PROVIDER_PKG)
	@# check the TARGET is set
	test $(TARGET)
	cd provider/cmd/${PROVIDER}/ && \
		yarn tsc && \
		cp package.json schema.json ./bin && \
		sed -i.bak -e "s/\$${VERSION}/$(PROVIDER_VERSION)/g" bin/package.json && \
		yarn run pkg . ${PKG_ARGS} --target node16-$(TARGET) --output $(WORKING_DIR)/$@

dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-amd64.tar.gz: bin/linux-amd64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-arm64.tar.gz: bin/linux-arm64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-amd64.tar.gz: bin/darwin-amd64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-arm64.tar.gz: bin/darwin-arm64/$(PROVIDER)
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-windows-amd64.tar.gz: bin/windows-amd64/$(PROVIDER).exe
dist/$(PROVIDER)-v$(PROVIDER_VERSION)-%.tar.gz:
	@mkdir -p dist
	@# $< is the last dependency (the binary path from above)
	tar --gzip -cf $@ README.md LICENSE -C $$(dirname $<) .

dist: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-amd64.tar.gz
dist: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-linux-arm64.tar.gz
dist: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-amd64.tar.gz
dist: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-darwin-arm64.tar.gz
dist: dist/$(PROVIDER)-v$(PROVIDER_VERSION)-windows-amd64.tar.gz

provider/cmd/${PROVIDER}/schema.json: bin/yq $(SCHEMA_FILE)
	bin/yq -o json '.' $(SCHEMA_FILE) > provider/cmd/${PROVIDER}/schema.json

provider/scripts/vendor/generate-provider-types.ts: AWSX_VERSION := $(shell cat .awsx.version)
provider/scripts/vendor/generate-provider-types.ts: .awsx.version
	@mkdir -p provider/scripts/vendor
	cd provider/scripts && \
		curl -sSL 'https://raw.githubusercontent.com/pulumi/pulumi-awsx/v$(AWSX_VERSION)/awsx/scripts/generate-provider-types.ts' > vendor/generate-provider-types.ts && \
		patch vendor/generate-provider-types.ts patches/0001-fixes.patch && \
		patch vendor/generate-provider-types.ts patches/0002-undefinedResource.patch

provider/scripts/vendor/pulumi-schema.d.ts: AWSX_VERSION := $(shell cat .awsx.version)
provider/scripts/vendor/pulumi-schema.d.ts: .awsx.version
	@mkdir -p provider/scripts/vendor
	curl -sSL 'https://raw.githubusercontent.com/pulumi/pulumi-awsx/v$(AWSX_VERSION)/awsx/scripts/pulumi-schema.d.ts' > provider/scripts/vendor/pulumi-schema.d.ts

# --------- Sentinel targets --------- #

.make/provider_mod_download: provider/cmd/${PROVIDER}/package.json
	cd provider/cmd/${PROVIDER}/ && \
		yarn install
	@touch $@

.make/generate_java: bin/pulumictl .pulumi/bin/pulumi $(SCHEMA_FILE)
	rm -rf sdk/java
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language java
	@touch $@

.make/generate_nodejs: bin/pulumictl .pulumi/bin/pulumi $(SCHEMA_FILE)
	rm -rf sdk/nodejs
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language nodejs
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" sdk/nodejs/tsconfig.json
	sed -i.bak -e 's/"remote"/".\/remote"/g' sdk/nodejs/*.ts
	sed -i.bak -e 's/"tls"/".\/tls"/g' sdk/nodejs/*.ts
	sed -i.bak -e 's/"tools"/".\/tools"/g' sdk/nodejs/*.ts
	rm sdk/nodejs/*.bak
	@touch $@

.make/generate_python: bin/pulumictl .pulumi/bin/pulumi $(SCHEMA_FILE)
	rm -rf sdk/python
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language python
	cp README.md sdk/python
	@touch $@

.make/generate_dotnet: bin/pulumictl .pulumi/bin/pulumi $(SCHEMA_FILE)
	rm -rf sdk/dotnet
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language dotnet
	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" sdk/dotnet/UnMango.KubernetesTheHardWay.csproj
	sed -i.bak -e "s/UnMango.KubernetesTheHardWay.Config.Kubeconfig/UnMango.KubernetesTheHardWay.Config.Outputs.Kubeconfig/g" sdk/dotnet/Tls/ClusterPki.cs
	rm sdk/dotnet/*.bak sdk/dotnet/**/*.bak
	@touch $@

.make/generate_go: bin/pulumictl .pulumi/bin/pulumi $(SCHEMA_FILE)
	rm -rf sdk/go
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language go
	@touch $@

.make/generate_types: vendor provider/cmd/${PROVIDER}/schema.json
	cd provider/scripts && yarn install && yarn gen-types
	cd provider/cmd/${PROVIDER} && sed -i.bak 's/input.remote.Connection/input.remote.ConnectionArgs/g' schema-types.ts && rm schema-types.ts.bak
	@touch $@

.make/nodejs_yarn_install: .make/generate_nodejs sdk/nodejs/package.json
	yarn install --cwd sdk/nodejs
	@touch $@

.make/build_nodejs: VERSION_JS = $(shell bin/pulumictl convert-version -l javascript -v "$(VERSION_GENERIC)")
.make/build_nodejs: bin/pulumictl .make/nodejs_yarn_install
	cd sdk/nodejs/ && \
		NODE_OPTIONS=--max-old-space-size=12288 yarn run tsc --diagnostics --incremental && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION_JS)/g" ./bin/package.json
	@touch $@

.make/build_python: VERSION_PYTHON = $(shell bin/pulumictl convert-version -l python -v "$(VERSION_GENERIC)")
.make/build_python: bin/pulumictl .make/generate_python
	cd sdk/python && \
		git clean -fxd && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^  version = .*/  version = "$(VERSION_PYTHON)"/g' ./bin/pyproject.toml && \
		rm ./bin/pyproject.toml.bak && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build && \
		cd ./bin && \
		../venv/bin/python -m build .
	@touch $@

.make/build_dotnet: VERSION_DOTNET = $(shell bin/pulumictl convert-version -l dotnet -v "$(PROVIDER_VERSION)")
.make/build_dotnet: bin/pulumictl .make/generate_dotnet
	cd sdk/dotnet && \
		echo "$(VERSION_DOTNET)" >version.txt && \
		dotnet build /p:Version=$(VERSION_DOTNET)
	@touch $@

.make/build_java: bin/pulumictl .make/generate_java
	cd sdk/java/ && \
		gradle --console=plain -Pversion=$(VERSION_GENERIC) build
	@touch $@

.make/build_go: .make/generate_go
	cd sdk/go/kubernetes-the-hard-way && go build
	@touch $@

.make/install_nodejs_sdk: .make/build_nodejs
	yarn link --cwd sdk/nodejs/bin
	@touch $@

.make/install_dotnet_sdk: .make/build_dotnet
	@mkdir -p nuget
	find sdk/dotnet/bin -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;
	if ! dotnet nuget list source | grep ${WORKING_DIR}; then \
		dotnet nuget add source ${WORKING_DIR}/nuget --name ${WORKING_DIR} \
	; fi
	@touch $@

.make/install_provider: bin/pulumictl provider/cmd/$(PROVIDER)/* $(PROVIDER_PKG)
	cd provider/cmd/${PROVIDER}/ && \
		yarn run pkg . ${PKG_ARGS} --target node16 --output ${WORKING_DIR}/bin/${PROVIDER}
	@touch $@
