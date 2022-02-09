.PHONY: build run rundl

# NOTE: This is a windows-only Makefile, this won't work on any other platform.

build:  # Builds animedl
	@rd /s /q dist && mkdir dist && cd dist && go build ../animedl/main.go

run:    # Runs animedl
	@cd dist && .\main.exe

rundl:  # Runs animedl
	@cd dist && .\main.exe dl "overlord"
