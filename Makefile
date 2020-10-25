run:
	go build -o streamed-logs-viewer \
	&& go run trash/main.go | SLV_DEBUG=1 ./streamed-logs-viewer
