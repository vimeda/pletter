OSFLAG:=
    UNAME_S := $(shell uname -s)
    OSFLAG = osx
    ifeq ($(UNAME_S),Linux)
        OSFLAG = linux
    endif

    UNAME_P := $(shell uname -p)
    ifeq ($(UNAME_P),x86_64)
        OSFLAG := $(OSFLAG)-x86_64
    endif
    ifneq ($(filter %86,$(UNAME_P)),)
        OSFLAG := $(OSFLAG)-x86_32
    endif

PROTOC_ZIP="protoc-3.8.0-$(OSFLAG).zip"

.PHONY: compile

proto: tools
	@echo "Compile the pb.go files"
	@protoc --go_out=pb *.proto

#---------------
#-- tools
#---------------

.PHONY: tools tools.protoc-gen-go tools.protoc

tools: tools.protoc-gen-go tools.protoc 

tools.protoc-gen-go:
	@command -v protoc-gen-go >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing protoc-gen-go"; \
		go get github.com/golang/protobuf/protoc-gen-go; \
	fi

tools.protoc:
	@command -v protoc >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing protoc"; \
			curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/$(PROTOC_ZIP); \
			unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc; \
			unzip -o $(PROTOC_ZIP) -d /usr/local include/*; \
			rm -f $(PROTOC_ZIP); \
	fi
