run:
	go build -o streamed-logs-viewer \
	&& go run trash/main.go | ./streamed-logs-viewer
