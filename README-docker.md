# 1. コンテナをビルド
docker-compose build

# 2. コンテナを起動
docker-compose up -d

# 普段（環境構築）
## コンテナ起動＆入る
### コンテナを起動
`docker-compose up -d`
### コンテナに入る
`docker exec -it atcoder_env /bin/bash`

## エイリアスの設定(Makefileでもいい)
### エイリアスを設定する場合
#### atcoderと打ってコンテナ入る
`echo 'alias atcoder="docker exec -it atcoder_env /bin/bash"' >> ~/.zshrc`
#### ojgoと打ってテストケースでテストする
`alias ojgo="oj t -c \"go run ./main.go\" -d test/"`
#### addgoと打ってtemplate.goをmain.goという名前で作成する
`alias addgo="cp /go/src/work/template.go ./main.go"`
### エイリアスの設定を適用
`source ~/.zshrc`もしくは`source ~/.bashrc`

## コンテナ停止、削除
## 作業終わったらコンテナ停止
`docker-compose down`
## Atcoder環境しばらく使わないなら、不要なコンテナやネットワーク設定、データ保存領域（ボリューム）、イメージを削除
`docker-compose down --rmi all --volumes`

# Makefile使う場合
イメージをビルドしコンテナを起動 + コンテナ内のBashで使う用のエイリアスの設定
`make up`
コンテナに入る。
`make atcoder`