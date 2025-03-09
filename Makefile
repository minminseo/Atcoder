# コンテナ名
CONTAINER_NAME=atcoder_env

# デフォルトで `all` を実行（コンテナ起動＋環境セットアップ）
all: up setup-container

# コンテナを起動（Docker Compose V1/V2 両対応）
up:
	@if docker compose up --build -d 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build -d; \
	fi

# コンテナを停止
down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# コンテナに入る（エイリアス不要）
atcoder:
	docker exec -it $(CONTAINER_NAME) /bin/bash

# `ojgo` を実行（コンテナ内で `oj` コマンドを実行）
ojgo:
	docker exec -it $(CONTAINER_NAME) bash -c "oj t -c 'go run main.go' -d test"

# `addgo` を実行（テンプレートコピー）
addgo:
	docker exec -it $(CONTAINER_NAME) bash -c "cp /work/template.go ./main.go"

# コンテナ内 (`bash`) の `.bashrc` に `ojgo` と `addgo`、`BROWSER` の設定を追加
setup:
	@if docker ps --format '{{.Names}}' | grep -q "^$(CONTAINER_NAME)$$"; then \
		docker exec -it $(CONTAINER_NAME) bash -c 'echo "alias ojgo=\"oj t -c \\\"go run main.go\\\" -d tests/\"" >> ~/.bashrc'; \
		docker exec -it $(CONTAINER_NAME) bash -c 'echo "alias addgo=\"cp /atcoder/work/template.go ./main.go\"" >> ~/.bashrc'; \
		docker exec -it $(CONTAINER_NAME) bash -c 'echo "export BROWSER=google-chrome" >> ~/.bashrc'; \
		docker exec -it $(CONTAINER_NAME) bash -c 'source ~/.bashrc'; \
		echo "コンテナ内の .bashrc にエイリアスと BROWSER 環境変数を追加しました！"; \
	else \
		echo "エラー: コンテナ $(CONTAINER_NAME) が起動していません。まず 'make up' で起動してください！"; \
	fi

# テストの実行
test:
	@echo "Running tests..."
	docker exec -it $(CONTAINER_NAME) bash -c "go test ./... -v"

# `make clean` でコンテナやキャッシュを削除
clean:
	@echo "Cleaning up Docker containers and volumes..."
	docker-compose down -v
	docker system prune -f
	docker volume prune -f

.PHONY: all up down atcoder ojgo addgo setup-container test clean
