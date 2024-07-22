/*
fmt.Println()で改行をしない場合、
プログラムの終了時にバッファがフラッシュされ、エラーメッセージや
システムのプロンプト（％など）が表示されることがある。そのため、
出力を明確に終了させるためには、fmt.Println()を使って明示的に
改行をする必要がある。
*/

package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Printf("%d", N)
	}
	fmt.Println("")
}
