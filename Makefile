APP=go-hello-server

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
	@"build/$(APP)"
