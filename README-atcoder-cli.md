# Atcoder CLI、online-judge-toolsの使い方
ログイン
`acc login`
abc100のURLを取得
`acc contect abc100`

abc100の問題別のURLを取得
`acc tasks abc100`

abc100用のディレクトリを作成。このコマンド実行後、abc100ディレクトリ配下にどの問題のディレクトリを配置するか選択肢が表示される。スペースで選択し、エンターを押すとabc100ディレクトリ配下に選択した問題のディレクトリ（a,b,cなど）が配置され、それぞれのディレクトリ内にはtestsというテストケースが入ったディレクトリが生成される。
`acc new abc100`
テスト(abc100/a/でabc100/tests/のテストケースを使ってmain.goを実行するなら)
`oj t -c "go run main.go" -d tests` or `ojgo`(エイリアス)
特定の問題のテストケースをダウンロード
`oj d https://atcoder.jp/contests/abc100/tasks/abc100_a`
