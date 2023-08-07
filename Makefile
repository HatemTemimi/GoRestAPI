SHELL := /bin/bash

ifneq (,$(wildcard ./app.env))
    include app.env
    export
endif

serve:
	source $(PWD)/scripts/setup.sh
	