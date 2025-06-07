size:=1


core:
	go build .
	GOTRACEBACK=crash ./test-coredump --size $(size) --delay 1


core-1G:
	go build .
	GOTRACEBACK=crash ./test-coredump --size 1024 --delay 1


