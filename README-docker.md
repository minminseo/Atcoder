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

## エイリアスの設定
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

# Atcoderでの使い方（コンテナ内でのCLI使い方）
```bash:
# AtCoderのログイン
acc login

# AtCoderのコンテストを取得
acc new abc300
cd abc300

# 問題をダウンロード
oj d https://atcoder.jp/contests/abc300/tasks/abc300_a

# 解答を提出
oj s main.py
```

# その他(コンテナの起動停止状態の確認)
エイリアス（dpsコマンド）を追加
`echo 'alias dps="docker ps --format '\''table {{.Names}}\t{{.Status}}\t{{.Ports}}'\''"' >> ~/.zshrc`
関数(atcoder-status)を追加
```zsh
cat <<EOF >> ~/.zshrc

# AtCoderのコンテナが動いているか確認する関数
function atcoder-status() {
    if docker ps | grep -q "atcoder_env"; then
        echo "✅ Atcoder環境は起動中！"
    else
        echo "❌ Atcoder環境は停止中。"
    fi
}
EOF
```
設定を反映
`source ~/.zshrc`

現在動作しているDockerコンテナ一覧を表示
`dps`

atcoder-status → atcoder_env が動いているか確認
`atcoder-status`

alias atcoder="docker exec -it atcoder_env /bin/bash"を試す。！！！！