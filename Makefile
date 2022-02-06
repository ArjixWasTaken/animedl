.PHONY: build run rundl



build:  # Builds animedl
	@rd /s /q dist && mkdir dist && cd dist && go build ../animedl/main.go

run:  # Runs animedl
	@cd dist && .\main.exe

rundl:  # Runs animedl
	@cd dist && .\main.exe dl "overlord iii"
