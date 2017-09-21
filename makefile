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


kill_deployed:
	@hyper rm -f jseminer1 jseminer2 jseminer3 jseminer4 jseminer5 jseminer6 jseminer7

deploy: ci
	@hyper pull adamveld12/gojse:latest
	@hyper run -d --size=l1 --name=jseminer1 --restart=always --env-file=./prod.env adamveld12/gojse
	@hyper run -d --size=l1 --name=jseminer2 --restart=always --env-file=./prod.env adamveld12/gojse
	@hyper run -d --size=l1 --name=jseminer3 --restart=always --env-file=./prod.env adamveld12/gojse
	@hyper run -d --size=l1 --name=jseminer4 --restart=always --env-file=./prod.env adamveld12/gojse
	@hyper run -d --size=l1 --name=jseminer5 --restart=always --env-file=./prod.env adamveld12/gojse
	@hyper run -d --size=l1 --name=jseminer6 --restart=always --env-file=./prod.env adamveld12/gojse
	@hyper run -d --size=l1 --name=jseminer7 --restart=always --env-file=./prod.env adamveld12/gojse
