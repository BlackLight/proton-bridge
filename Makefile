export GO111MODULE=on

# By default, the target OS is the same as the host OS,
# but this can be overridden by setting TARGET_OS to "windows"/"darwin"/"linux".
GOOS:=$(shell go env GOOS)
TARGET_CMD?=Desktop-Bridge
TARGET_OS?=${GOOS}
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

## Build
.PHONY: build build-gui build-nogui build-launcher versioner hasher

# Keep version hardcoded so app build works also without Git repository.
BRIDGE_APP_VERSION?=2.3.0+git
APP_VERSION:=${BRIDGE_APP_VERSION}
APP_FULL_NAME:=Proton Bridge
APP_VENDOR:=Proton AG
SRC_ICO:=bridge.ico
SRC_ICNS:=Bridge.icns
SRC_SVG:=bridge.svg
EXE_NAME:=proton-bridge
REVISION:=$(shell git rev-parse --short=10 HEAD)
BUILD_TIME:=$(shell date +%FT%T%z)

BUILD_FLAGS:=-tags='${BUILD_TAGS}'
BUILD_FLAGS_LAUNCHER:=${BUILD_FLAGS}
BUILD_FLAGS_GUI:=-tags='${BUILD_TAGS} build_qt'
GO_LDFLAGS:=$(addprefix -X github.com/ProtonMail/proton-bridge/v2/internal/constants., Version=${APP_VERSION} Revision=${REVISION} BuildTime=${BUILD_TIME})
GO_LDFLAGS+=-X "github.com/ProtonMail/proton-bridge/v2/internal/constants.FullAppName=${APP_FULL_NAME}"
ifneq "${BUILD_LDFLAGS}" ""
    GO_LDFLAGS+=${BUILD_LDFLAGS}
endif
GO_LDFLAGS_LAUNCHER:=${GO_LDFLAGS}
ifeq "${TARGET_OS}" "windows"
    GO_LDFLAGS_LAUNCHER+=-H=windowsgui
endif

BUILD_FLAGS+=-ldflags '${GO_LDFLAGS}'
BUILD_FLAGS_GUI+=-ldflags "${GO_LDFLAGS}"
BUILD_FLAGS_LAUNCHER+=-ldflags '${GO_LDFLAGS_LAUNCHER}'

DEPLOY_DIR:=cmd/${TARGET_CMD}/deploy
DIRNAME:=$(shell basename ${CURDIR})
EXE:=${EXE_NAME}
EXE_GO:=bridge
EXE_GUI_NAME:=bridge-gui
EXE_GUI:=${EXE_GUI_NAME}
LAUNCHER_PATH:=./cmd/launcher/
LAUNCHER_EXE:=launcher-${EXE}

ifeq "${TARGET_OS}" "windows"
    EXE:=${EXE}.exe
    EXE_GUI:=${EXE_GUI}.exe
    EXE_GO:=${EXE_GO}.exe
    LAUNCHER_EXE:=${LAUNCHER_EXE}.exe
    RESOURCE_FILE:=resource.syso
endif
ifeq "${TARGET_OS}" "darwin"
    DARWINAPP_CONTENTS:=${DEPLOY_DIR}/darwin/${EXE}.app/Contents
    EXE:=${EXE}.app
    EXE_GUI:=${EXE_GUI_NAME}.app
    EXE_BINARY_DARWIN:=Contents/MacOS/${EXE_NAME}
    EXE_GO:=${EXE_GUI}/Contents/MacOS/${EXE_GO}
endif
EXE_TARGET:=${DEPLOY_DIR}/${TARGET_OS}/${EXE}
EXE_GUI_TARGET:=${DEPLOY_DIR}/${TARGET_OS}/${EXE_GUI}
EXE_GO_TARGET:=${DEPLOY_DIR}/${TARGET_OS}/${EXE_GO}

TGZ_TARGET:=bridge_${TARGET_OS}_${REVISION}.tgz

ifdef QT_API
    VENDOR_TARGET:=prepare-vendor update-qt-docs
else
    VENDOR_TARGET=update-vendor
endif

build: build-gui

build-gui: ${TGZ_TARGET}

build-nogui: ${EXE_NAME}

go-build=go build $(1) -o $(2) $(3)
go-build-unify=${go-build}
ifeq "${GOOS}-$(shell uname -m)" "darwin-arm64"
	go-build-unify= \
		CGO_ENABLED=1 GOARCH=arm64 $(call go-build,$(1),$(2)_arm,$(3)) && \
		CGO_ENABLED=1 GOARCH=amd64 $(call go-build,$(1),$(2)_amd,$(3)) && \
		lipo -create -output $(2) $(2)_arm $(2)_amd && rm -f $(2)_arm $(2)_amd
endif

${EXE_NAME}: gofiles
	$(call go-build-unify,${BUILD_FLAGS},"${EXE_NAME}","cmd/${TARGET_CMD}/main.go")

build-launcher: ${RESOURCE_FILE}
ifeq "${GOOS}" "windows"
	powershell Copy-Item ${ROOT_DIR}/${RESOURCE_FILE} ${ROOT_DIR}/${LAUNCHER_PATH}${RESOURCE_FILE}
endif
	$(call go-build-unify,${BUILD_FLAGS_LAUNCHER},"${LAUNCHER_EXE}","${LAUNCHER_PATH}")
ifeq "${GOOS}" "windows"
	powershell Remove-Item ${ROOT_DIR}/${LAUNCHER_PATH}${RESOURCE_FILE} -Force
endif

versioner:
	go build ${BUILD_FLAGS} -o versioner utils/versioner/main.go

hasher:
	go build -o hasher utils/hasher/main.go

${TGZ_TARGET}: ${DEPLOY_DIR}/${TARGET_OS}
	rm -f $@
	tar -czvf $@ -C ${DEPLOY_DIR}/${TARGET_OS} .

${DEPLOY_DIR}/linux: ${EXE_TARGET} build-launcher
	cp -pf ./dist/${SRC_SVG} ${DEPLOY_DIR}/linux/logo.svg
	cp -pf ./LICENSE ${DEPLOY_DIR}/linux/
	cp -pf ./Changelog.md ${DEPLOY_DIR}/linux/
	cp -pf ./dist/${EXE_NAME}.desktop ${DEPLOY_DIR}/linux/
	cp -pf ${LAUNCHER_EXE} ${DEPLOY_DIR}/linux/

${DEPLOY_DIR}/darwin: ${EXE_TARGET} build-launcher
	mv ${EXE_TARGET}/Contents/MacOS/{${EXE_GUI_NAME},${EXE_NAME}}
	perl -i -pe"s/>${EXE_GUI_NAME}/>${LAUNCHER_EXE}/g" ${ROOT_DIR}/${EXE_TARGET}/Contents/Info.plist
	cp ./dist/${SRC_ICNS} ${DARWINAPP_CONTENTS}/Resources/${SRC_ICNS}
	cp LICENSE ${DARWINAPP_CONTENTS}/Resources/
	rm -rf "${DARWINAPP_CONTENTS}/Frameworks/QtWebEngine.framework"
	rm -rf "${DARWINAPP_CONTENTS}/Frameworks/QtWebView.framework"
	rm -rf "${DARWINAPP_CONTENTS}/Frameworks/QtWebEngineCore.framework"
	cp ${LAUNCHER_EXE}  ${EXE_TARGET}/Contents/MacOS/${LAUNCHER_EXE}
	./utils/remove_non_relative_links_darwin.sh "${EXE_TARGET}/${EXE_BINARY_DARWIN}"

${DEPLOY_DIR}/windows: ${EXE_TARGET} build-launcher
	cp ./dist/${SRC_ICO} ${DEPLOY_DIR}/windows/logo.ico
	cp LICENSE ${DEPLOY_DIR}/windows/
	cp ${LAUNCHER_EXE} ${DEPLOY_DIR}/windows/$(notdir ${LAUNCHER_EXE})
	# plugins are installed in a plugins folder while needs to be near the exe
	mv ${DEPLOY_DIR}/windows/plugins/* ${DEPLOY_DIR}/windows/.
	rm -rf ${DEPLOY_DIR}/windows/plugins

${EXE_TARGET}: check-build-essentials ${EXE_NAME}
	# TODO: resource.syso for windows
	cd internal/frontend/bridge-gui/bridge-gui && \
		BRIDGE_APP_FULL_NAME="${APP_FULL_NAME}" \
		BRIDGE_VENDOR="${APP_VENDOR}" \
		BRIDGE_APP_VERSION=${APP_VERSION} \
		BRIDGE_REVISION=${REVISION} \
		BRIDGE_BUILD_TIME=${BUILD_TIME} \
		BRIDGE_GUI_BUILD_CONFIG=Release \
		BRIDGE_INSTALL_PATH=${ROOT_DIR}/${DEPLOY_DIR}/${GOOS} \
		./build.sh install
	cp -pf "${ROOT_DIR}/${EXE_NAME}" "$(ROOT_DIR)/${EXE_GO_TARGET}"
	mv "${ROOT_DIR}/${EXE_GUI_TARGET}" "$(ROOT_DIR)/${EXE_TARGET}"

WINDRES_YEAR:=$(shell date +%Y)
APP_VERSION_COMMA:=$(shell echo "${APP_VERSION}" | sed -e 's/[^0-9,.]*//g' -e 's/\./,/g')
${RESOURCE_FILE}: ./dist/info.rc ./dist/${SRC_ICO} .FORCE
	rm -f ./*.syso
	windres --target=pe-x86-64 \
		-I ./internal/frontend/share/ \
		-D ICO_FILE=${SRC_ICO} \
		-D EXE_NAME="${EXE_NAME}" \
		-D FILE_VERSION="${APP_VERSION}" \
		-D ORIGINAL_FILE_NAME="${EXE}" \
		-D PRODUCT_VERSION="${APP_VERSION}" \
		-D FILE_VERSION_COMMA=${APP_VERSION_COMMA} \
		-D YEAR=${WINDRES_YEAR} \
		-o ./${RESOURCE_FILE} $<

## Dev dependencies
.PHONY: install-devel-tools install-linter install-go-mod-outdated install-git-hooks
LINTVER:="v1.47.2"
LINTSRC:="https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"

install-dev-dependencies: install-devel-tools install-linter install-go-mod-outdated

install-devel-tools: check-has-go
	go get -v github.com/golang/mock/gomock
	go get -v github.com/golang/mock/mockgen
	go get -v github.com/go-delve/delve

install-linter: check-has-go
	curl -sfL $(LINTSRC) | sh -s -- -b $(shell go env GOPATH)/bin $(LINTVER)

install-go-mod-outdated:
	which go-mod-outdated || go install github.com/psampaz/go-mod-outdated@latest

install-git-hooks:
	cp utils/githooks/* .git/hooks/
	chmod +x .git/hooks/*

## Checks, mocks and docs
.PHONY: check-has-go check-build-essentials add-license change-copyright-year test bench coverage mocks lint-license lint-golang lint updates doc release-notes
check-has-go:
	@which go || (echo "Install Go-lang!" && exit 1)
	go version


check_is_installed=if ! which $(1) > /dev/null; then echo "Please install $(1)"; exit 1; fi
check-build-essentials: check-qt-dir
	@$(call check_is_installed,zip)
	@$(call check_is_installed,unzip)
	@$(call check_is_installed,tar)
	@$(call check_is_installed,curl)
ifneq "${GOOS}" "windows"
	@$(call check_is_installed,cmake)
	@$(call check_is_installed,ninja)
endif

check-qt-dir:
	@if ! ls "${QT6DIR}/bin/qt.conf" > /dev/null; then echo "Please set QT6DIR"; exit 1; fi

add-license:
	./utils/missing_license.sh add

change-copyright-year:
	./utils/missing_license.sh change-year

test: gofiles
	go test -coverprofile=/tmp/coverage.out -run=${TESTRUN} \
		./internal/...\
		./pkg/...

bench:
	go test -run '^$$' -bench=. -memprofile bench_mem.pprof -cpuprofile bench_cpu.pprof ./internal/store
	go tool pprof -png -output bench_mem.png bench_mem.pprof
	go tool pprof -png -output bench_cpu.png bench_cpu.pprof

coverage: test
	go tool cover -html=/tmp/coverage.out -o=coverage.html

integration-test-bridge:
	${MAKE} -C test test-bridge

mocks:
	mockgen --package mocks github.com/ProtonMail/proton-bridge/v2/internal/users Locator,PanicHandler,CredentialsStorer,StoreMaker > internal/users/mocks/mocks.go
	mockgen --package mocks github.com/ProtonMail/proton-bridge/v2/pkg/listener Listener > internal/users/mocks/listener_mocks.go
	mockgen --package mocks github.com/ProtonMail/proton-bridge/v2/internal/store PanicHandler,BridgeUser,ChangeNotifier,Storer > internal/store/mocks/mocks.go
	mockgen --package mocks github.com/ProtonMail/proton-bridge/v2/pkg/listener Listener > internal/store/mocks/utils_mocks.go
	mockgen --package mocks github.com/ProtonMail/proton-bridge/v2/pkg/pmapi Client,Manager > pkg/pmapi/mocks/mocks.go
	mockgen --package mocks github.com/ProtonMail/proton-bridge/v2/pkg/message Fetcher > pkg/message/mocks/mocks.go

lint: gofiles lint-golang lint-license lint-dependencies lint-changelog

lint-license:
	./utils/missing_license.sh check

lint-dependencies:
	./utils/dependency_license.sh check

lint-changelog:
	./utils/changelog_linter.sh Changelog.md

lint-golang:
	which golangci-lint || $(MAKE) install-linter
	$(info linting with GOMAXPROCS=${GOMAXPROCS})
	golangci-lint run ./...

updates: install-go-mod-outdated
	# Uncomment the "-ci" to fail the job if something can be updated.
	go list -u -m -json all | go-mod-outdated -update -direct #-ci

doc:
	godoc -http=:6060

release-notes: release-notes/bridge_stable.html release-notes/bridge_early.html

release-notes/%.html: release-notes/%.md
	./utils/release_notes.sh $^

.PHONY: gofiles
# Following files are for the whole app so it makes sense to have them in bridge package.
# (Options like cmd or internal were considered and bridge package is the best place for them.)
gofiles: ./internal/bridge/credits.go
./internal/bridge/credits.go: ./utils/credits.sh go.mod
	cd ./utils/ && ./credits.sh bridge

## Run and debug
.PHONY: run run-qt run-qt-cli run-nogui run-nogui-cli run-debug run-qml-preview clean-vendor clean-frontend-qt clean-frontend-qt-common clean

LOG?=debug
LOG_IMAP?=client # client/server/all, or empty to turn it off
LOG_SMTP?=--log-smtp # empty to turn it off
RUN_FLAGS?=-m -l=${LOG} --log-imap=${LOG_IMAP} ${LOG_SMTP}

run: run-nogui-cli

run-qt: ${EXE_TARGET}
	PROTONMAIL_ENV=dev ./$< ${RUN_FLAGS} 2>&1 | tee last.log
run-qt-cli: ${EXE_TARGET}
	PROTONMAIL_ENV=dev ./$< ${RUN_FLAGS} -c

run-nogui: clean-vendor gofiles
	PROTONMAIL_ENV=dev go run ${BUILD_FLAGS} cmd/${TARGET_CMD}/main.go ${RUN_FLAGS} | tee last.log
run-nogui-cli: clean-vendor gofiles
	PROTONMAIL_ENV=dev go run ${BUILD_FLAGS} cmd/${TARGET_CMD}/main.go ${RUN_FLAGS} -c

run-debug:
	PROTONMAIL_ENV=dev dlv debug --build-flags "${BUILD_FLAGS}" cmd/${TARGET_CMD}/main.go -- ${RUN_FLAGS} --noninteractive

run-qml-preview:
	find internal/frontend/qml/ -iname '*qmlc' | xargs rm -f
	bridge_preview internal/frontend/qml/Bridge_test.qml


clean-frontend-qt:
	$(MAKE) -C internal/frontend -f Makefile.local clean

clean-vendor: clean-frontend-qt clean-frontend-qt-common
	rm -rf ./vendor

clean-gui:
	cd internal/frontend/bridge-gui/ && \
		rm -f Version.h && \
		rm -rf cmake-build-*/

clean-vcpkg:
	git submodule deinit -f ./extern/vcpkg
	rm -rf ./.git/submodule/vcpkg
	rm -rf ./extern/vcpkg
	git checkout -- extern/vcpkg

clean: clean-vendor clean-gui clean-vcpkg
	rm -rf vendor-cache
	rm -rf cmd/Desktop-Bridge/deploy
	rm -rf cmd/Import-Export/deploy
	rm -f build last.log mem.pprof main.go
	rm -f ${RESOURCE_FILE}
	rm -f release-notes/bridge.html
	rm -f release-notes/import-export.html


.PHONY: generate
generate:
	go generate ./...
	$(MAKE) add-license

.FORCE:
