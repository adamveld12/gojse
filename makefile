.PHONY: ci clean deploy dev docker kill_deploy test

default: docker dev

dev: clean gojse
	@docker run -it --rm \
	-e GOJSE_USERNAME=$$GOJSE_USERNAME \
	-e GOJSE_PASSWORD=$$GOJSE_PASSWORD \
	-v $$PWD:/app/ \
	adamveld12/gojse

clean:
	@rm -rf ./gojse

ci: clean test docker
	@docker push adamveld12/gojse

docker: gojse
	@docker build -t adamveld12/gojse .

gojse:
	@CGO_ENABLED=0 GOOS=linux go build .

test:
	@go test -v -cover
