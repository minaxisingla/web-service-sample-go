GOPKGNAME = minaxi/web-service-sample
SOURCE_DIR ?= $(shell pwd)
export GOPATH = $(SOURCE_DIR)/$(BUILD_DIR)

init:
	@echo "GO VERSION IS:" 'go version'

get:
	@mkdir -p $(BUILD_DIR)/src/$(GOPKGNAME)
	@rsync $(BUILD_DIR)/src/$(GOPKGNAME)
	cd $(BUILD_DIR)/src/$(GOPKGNAME)

build:	get
			@echo Installed web-service-sample $(BUILD_SHORTCOMMIT)/$(BUILD_VERSION) $(BUILD_DATE)
			go install -v $(GOPKGNAME)/...
			go run web.go

test:
			go test $(GOPATH)/webservice

publish: