default: build_sauth build_streamrcv

build_sauth:
	@cd sauth && make

build_streamrcv:
	@cd streamrcv && make