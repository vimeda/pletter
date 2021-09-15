OSFLAG:=
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        OSFLAG = linux
		UNAME_P := $(shell uname -p)
		ifeq ($(UNAME_P),x86_64)
			OSFLAG := $(OSFLAG)-x86_64
		endif
		ifneq ($(filter %86,$(UNAME_P)),)
			OSFLAG := $(OSFLAG)-x86_32
		endif
    else ifeq ($(UNAME_S),Darwin)
		UNAME_M := $(shell uname -m)
		OSFLAG = osx-$(UNAME_M)
	else
		$(error Unsupported OS)
    endif

PROTOC_VERSION="3.12.1"
PROTOC_ZIP="protoc-$(PROTOC_VERSION)-$(OSFLAG).zip"

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

.PHONY: proto

proto: tools
	@echo "$(OK_COLOR)==> Compile the pb.go files$(NO_COLOR)"
	@protoc --go_out=pb --go_opt=paths=source_relative *.proto

#---------------
#-- tools
#---------------
.PHONY: tools tools.protoc-gen-go tools.protoc

tools: tools.protoc-gen-go tools.protoc 

tools.protoc-gen-go:
	@command -v protoc-gen-go >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "$(OK_COLOR)==> Installing protoc-gen-go$(NO_COLOR)"; \
		go get -u google.golang.org/protobuf/cmd/protoc-gen-go; \
	fi

tools.protoc:
	@command -v protoc >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "$(OK_COLOR)==> Installing fresh protoc $(PROTOC_VERSION)$(NO_COLOR)"; \
		curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_ZIP); \
		unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc; \
		unzip -o $(PROTOC_ZIP) -d /usr/local include/*; \
		rm -f $(PROTOC_ZIP); \
	elif ! protoc --version | grep -q "libprotoc $(PROTOC_VERSION)"; then \
		echo "$(OK_COLOR)==> Installing protoc $(PROTOC_VERSION)$(NO_COLOR)"; \
		rm -f /usr/local/bin/protoc; \
		rm -rf /usr/local/include/google; \
		curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_ZIP); \
		unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc; \
		unzip -o $(PROTOC_ZIP) -d /usr/local include/*; \
		rm -f $(PROTOC_ZIP); \
	else \
		echo "$(OK_COLOR)==> protoc up-to-date ($(PROTOC_VERSION))$(NO_COLOR)"; \
	fi \
