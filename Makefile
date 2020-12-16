.PHONY: test

build:
	mkdir -p deployments/.build/bin
	env GOOS=linux go build -ldflags="-s -w" -o deployments/.build/bin/handler cmd/handler/main.go
	cd deployments/.build && zip -r sentron-sourcemaps.zip ./bin

clean:
	rm -rf deployments/.build

deploy: export AWS_DEFAULT_PROFILE := sentron
deploy: test clean build tf-init tf-plan tf-apply

test:
	go test ./...

test-e2e:
	./test/e2e/query-api.sh

tf-init:
	cd deployments && terraform init

tf-plan:
	cd deployments && terraform plan -out .terraform-plan

tf-apply:
	cd deployments && terraform apply -auto-approve .terraform-plan

tf-destroy:
	cd deployments && terraform destroy

tf-fmt:
	cd deployments && terraform fmt **/*.tf

bench:
	go test -bench . github.com/jpstevens/sentron-sourcemaps/internal/pkg/sourcemap -cpuprofile cpu.pb.gz -memprofile mem.pb.gz

pprof-cpu: bench
	go tool pprof -http=localhost:9090 sourcemap.test cpu.pb.gz

pprof-mem: bench
	go tool pprof -http=localhost:9090 sourcemap.test mem.pb.gz
