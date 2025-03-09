# Python + Node.js ベースイメージ
FROM python:3.10

# 作業ディレクトリの設定
WORKDIR /atcoder

# 必要なパッケージをインストール
RUN apt-get update && apt-get install -y curl \
    && curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
    && apt-get install -y nodejs \
    && npm install -g atcoder-cli \
    && curl -OL https://go.dev/dl/go1.21.5.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz \
    && rm go1.21.5.linux-amd64.tar.gz \
    && pip install --upgrade pip \
    && pip install online-judge-tools \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

#addgoのエイリアスをbashrcに追加
# RUN echo 'alias addgo="cp /work/template.go ./main.go"' >> ~/.bashrc # Makefileでやるからいらない
# bashrcを毎回読み込むようにbash_profileに設定
# RUN echo 'source ~/.bashrc' >> ~/.bash_profile # docker-compose.ymlで環境変数したからいらない


# 環境変数を設定
ENV PATH="/usr/local/go/bin:${PATH}"

# コンテナ起動時のデフォルトコマンド
CMD ["/bin/bash"]
