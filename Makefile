PACK            := kubernetes-the-hard-way
PROJECT         := github.com/unstoppablemango/pulumi-${PACK}

PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION_PATH    := provider/pkg/version.Version

WORKING_DIR     := $(shell pwd)
SCHEMA_FILE     := ${WORKING_DIR}/schema.yaml

PROVIDER_PKG    := $(shell find provider/pkg -type f)

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

.PHONY: default ensure
default: provider build_sdks
ensure: bin/pulumictl .make/provider_mod_download

# Binaries
.PHONY: codegen provider
codegen: bin/$(CODEGEN)
provider: bin/$(LOCAL_PROVIDER_FILENAME)

.PHONY: install_provider
install_provider: .make/install_provider

.PHONY: generate generate_java generate_nodejs generate_python generate_dotnet generate_go
generate: generate_java generate_nodejs generate_python generate_dotnet generate_go
generate_java: .make/generate_java
generate_nodejs: .make/generate_nodejs
generate_python: .make/generate_python
generate_dotnet: .make/generate_dotnet
#generate_go: .make/generate_go_local
generate_go: .make/generate_go

.PHONY: local_generate_code
local_generate_code: generate_java
local_generate_code: generate_nodejs
local_generate_code: generate_python
local_generate_code: generate_dotnet
local_generate_code: generate_go
local_generate: local_generate_code

.PHONY: build only_build build_sdks build_nodejs build_python build_dotnet build_java build_go
build: codegen local_generate provider build_sdks
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

.PHONY: clean
clean:
	rm -rf nuget
	rm -rf .make
	rm -rf bin
	rm -rf dist
	rm -rf sdk/dotnet/{bin,obj}
	rm -rf sdk/nodejs/bin
	rm -rf sdk/go
	rm -rf sdk/python/bin
	rm -rf sdk/java/{.gradle,build}
	if dotnet nuget list source | grep "$(WORKING_DIR)"; then \
		dotnet nuget remove source "$(WORKING_DIR)" \
	; fi

.PHONY: tidy
tidy:
	cd provider && go mod tidy

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
	cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN)

bin/$(LOCAL_PROVIDER_FILENAME): bin/pulumictl .make/provider_mod_download provider/cmd/$(PROVIDER)/*.go $(PROVIDER_PKG)
	cd provider/cmd/$(PROVIDER) && VERSION=${VERSION_GENERIC} SCHEMA=${SCHEMA_FILE} go generate main.go
	cd provider && \
		CGO_ENABLED=0 go build -o $(WORKING_DIR)/bin/$(LOCAL_PROVIDER_FILENAME) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER)

bin/linux-amd64/$(PROVIDER): TARGET := linux-amd64
bin/linux-arm64/$(PROVIDER): TARGET := linux-arm64
bin/darwin-amd64/$(PROVIDER): TARGET := darwin-amd64
bin/darwin-arm64/$(PROVIDER): TARGET := darwin-arm64
bin/windows-amd64/$(PROVIDER).exe: TARGET := windows-amd64
bin/%/$(PROVIDER) bin/%/$(PROVIDER).exe: bin/pulumictl .make/provider_mod_download provider/cmd/$(PROVIDER)/*.go $(PROVIDER_PKG)
	@# check the TARGET is set
	test $(TARGET)
	cd provider && \
		export GOOS=$$(echo "$(TARGET)" | cut -d "-" -f 1) && \
		export GOARCH=$$(echo "$(TARGET)" | cut -d "-" -f 2) && \
		CGO_ENABLED=0 go build -o ${WORKING_DIR}/$@ $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER)

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

# --------- Sentinel targets --------- #

.make/provider_mod_download: provider/go.mod provider/go.sum
	cd provider && go mod download
	@touch $@

.make/generate_java: bin/pulumictl .pulumi/bin/pulumi schema.yaml
	rm -rf sdk/java
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language java
	@touch $@

.make/generate_nodejs: bin/pulumictl .pulumi/bin/pulumi schema.yaml
	rm -rf sdk/nodejs
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language nodejs
	sed -i.bak -e "s/sourceMap/inlineSourceMap/g" sdk/nodejs/tsconfig.json
	rm sdk/nodejs/tsconfig.json.bak
	@touch $@

.make/generate_python: bin/pulumictl .pulumi/bin/pulumi schema.yaml
	rm -rf sdk/python
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language python
	cp README.md sdk/python
	@touch $@

.make/generate_dotnet: bin/pulumictl .pulumi/bin/pulumi schema.yaml
	rm -rf sdk/dotnet
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language dotnet
#	sed -i.bak -e "s/<\/Nullable>/<\/Nullable>\n    <UseSharedCompilation>false<\/UseSharedCompilation>/g" sdk/dotnet/UnMango.KubernetesTheHardWay.csproj
#	rm sdk/dotnet/UnMango.KubernetesTheHardWay.csproj.bak
	@touch $@

.make/generate_go: bin/pulumictl .pulumi/bin/pulumi schema.yaml
	rm -rf sdk/go
	.pulumi/bin/pulumi package gen-sdk $(SCHEMA_FILE) --language go
	@touch $@

# TODO: Fix or remove if not needed
#.make/generate_go_local: bin/pulumictl bin/$(CODEGEN)
#	@mkdir -p sdk/pulumi-${PACK}
#	@# Unmark this is as an up-to-date local build
#	rm -f .make/prepublish_go
#	rm -rf $$(find sdk/pulumi-${PACK} -mindepth 1 -maxdepth 1 ! -name ".git")
#	bin/$(CODEGEN) go $(VERSION_GENERIC) ${WORKING_DIR}/sdk/pulumi-${PACK} $(SCHEMA_FILE)
#	@# Tidy up all go.mod files
#	find sdk/pulumi-${PACK} -type d -maxdepth 1 -exec sh -c "cd \"{}\" && go mod tidy" \;
#	@touch $@

#.make/prepublish_go:
#	@# Unmark this is as an up-to-date local build
#	rm -f .make/generate_go_local
#	@# Remove go module replacements which are added for local testing
#	@# Note: must use `sed -i -e` to be portable - but leaves go.mod-e behind on macos
#	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.mod -exec sed -i -e '/replace github\.com\/pulumi\/pulumi-azure-native-sdk/d' {} \;
#	@# Remove sed backup files if using older sed versions
#	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.mod-e -delete
#	@# Delete go.sum files as these are not used at the point of publishing.
#	@# This is because we depend on the root package which will come from the same release commit, that doesn't yet exist.
#	find sdk/pulumi-azure-native-sdk -maxdepth 2 -type f -name go.sum -delete
#	cp README.md LICENSE sdk/pulumi-azure-native-sdk/
#	@touch $@

.make/nodejs_yarn_install: .make/generate_nodejs sdk/nodejs/package.json
	yarn install --cwd sdk/nodejs
	@touch $@

.make/build_nodejs: VERSION_JS = $(shell bin/pulumictl convert-version -l javascript -v "$(VERSION_GENERIC)")
.make/build_nodejs: bin/pulumictl .make/nodejs_yarn_install
	cd sdk/nodejs/ && \
		NODE_OPTIONS=--max-old-space-size=12288 yarn run tsc --diagnostics --incremental && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		mkdir -p bin/scripts && \
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
		echo "${PACK}\n$(VERSION_DOTNET)" >version.txt && \
		dotnet build /p:Version=$(VERSION_DOTNET)
	@touch $@

.make/build_java: bin/pulumictl .make/generate_java
	cd sdk/java/ && \
		gradle --console=plain -Pversion=$(VERSION_GENERIC) build
	@touch $@

# TODO: Fix
.make/build_go: .make/generate_go
	find sdk/go -maxdepth 1 -type d -exec sh -c "cd \"{}\" && go build" \;
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

.make/install_provider: bin/pulumictl .make/provider_mod_download provider/cmd/$(PROVIDER)/* $(PROVIDER_PKG)
	cd provider && go install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER)
	@touch $@
