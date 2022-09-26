docker-go: vendor
	docker run -it --rm -v $(shell pwd):/app -w /app kleverio/golang:cicd make build

vendor:
	go mod vendor

build: clean
	go build -o app

clean:
	rm -f app

docker-build: docker-go
	docker build -t kleverio/klog:sample .

docker-push: docker-build
	docker push kleverio/klog:sample

deploy:
	kubectl apply -f klog.yaml
	kubectl rollout restart deploy klog -n klog