# Based on https://medium.com/@olebedev/live-code-reloading-for-golang-web-projects-in-19-lines-8b2e8777b1ea
PID = /tmp/lss.pid

live-reload: restart
	@fswatch -e '.*/workspace/.*' -e '**/*.log' -o . | xargs -n1 -I{}  make restart || make kill

kill:
	@kill `cat $(PID)` || true

before:
	@echo "Build binary"
	@go build -o /tmp/lss lss.go

restart: kill before
	 @/tmp/lss server -b localhost -p 5000 & echo $$! > $(PID)

.PHONY: serve restart kill before # let's go to reserve rules names
