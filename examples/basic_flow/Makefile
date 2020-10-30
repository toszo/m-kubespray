EPIKUBESPRAY := epiphanyplatform/epikubespray:0.0.1

makefile_dir_path := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
#makes it easier to replace the value
M_NAME            ?= epikubespray-module-test
M_SHARED_EXTERNAL ?= $(makefile_dir_path)/$(M_NAME)/shared

export

all: apply

init: init-mkubespray

init-mkubespray: setup
	@docker run --rm \
		-v $(M_SHARED_EXTERNAL):/shared \
		-t $(EPIKUBESPRAY) \
		init \
		M_NAME=$(M_NAME)

apply: 
	@docker run --rm \
		-v $(M_SHARED_EXTERNAL):/shared \
		-t $(EPIKUBESPRAY) \
		apply \
		M_NAME=$(M_NAME)

setup: $(M_SHARED_EXTERNAL)

$(M_SHARED_EXTERNAL):
	@mkdir -p $(dir $@)