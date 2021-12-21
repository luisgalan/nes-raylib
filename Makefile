.PHONY: build clean

build:
	go build -v -o build/app
	cp -r assets build/assets

clean:
	@echo "Cleaning up build directory..."
	rm -rf build/*

