SOURCE = ./cmd/hanviet
TARGET = ./build/hanviet
BINARY = hanviet
GOARCH = amd64

CC = go
go = $(shell which go 2> /dev/null)

ifeq (, $(go))
	@printf "\e[91mGo not found!"
endif

$(BINARY): clean $(SOURCE)
	@printf "\e[33mBuilding\e[90m %s\e[0m\n" $@
	@go build -o $(TARGET) $(SOURCE)
	@printf "\e[34mDone!\e[0m\n"

test: clean
	@printf "\e[33mTesting...\e[0m\n"
	go test $(SOURCE)
	@printf "\e[34mDone!\e[0m\n"

clean:
	@rm -f $(OUT)
	@printf "\e[34mAll clear!\e[0m\n"

install: $(OUT)
	@printf "\e[33mInstalling\e[90m %s\e[0m\n" $(OUT)
	sudo rm -f /usr/local/bin/$(OUT)
	sudo ln -s $(PWD)/$(TARGET) /usr/local/bin/$(OUT)
	@printf "\e[34mDone!\e[0m\n"

uninstall:
	@printf "\e[33mRemoving\e[90m %s\e[0m\n" $(OUT)
	sudo rm -f /usr/local/bin/$(OUT)
	@printf "\e[34mDone!\e[0m\n"