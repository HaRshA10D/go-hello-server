APP=go-hello-server
ALL_PACKAGES= $(shell go list ./...)

# TO INSTALL EXTERNAL DEPENDENCIES
setup:
	@go get -u golang.org/x/lint/golint

# TO RUN LINT
lint:
	@golint $(ALL_PACKAGES)

# TO COPY DEV CONFIG
copy-dev-config:
	@rm -f application.yml
	@cp configs/application.yml.dev application.yml

# TO BUILD THE APP
build: clean
	@echo "Building app"
	@go build -o "build/$(APP)" main.go

# TO CLEAN BUILD FOLDER
clean:
	@echo "Cleaning build folder"
	@rm -rf build && mkdir build

# TO RUN THE APP FROM BUILD FOLDER
run:
	@echo "Running app from build/$(APP)"
	@"build/$(APP)" $(CMD)

# TO BUILD AND RUN THE APP FOR DEV CONFIG
run-dev: copy-dev-config build
	@echo "Running app from build/$(APP)"
	@"build/$(APP)" $(CMD)
