# シンボル解決をする
## symbolTable
- newで定義済みをなんとかする。
- addEntryでテーブルにシンボルとアドレスを追記。
- containsでシンボルが含むかどうかを返す。
- getAddressでアドレスの値を得る。
## シンボルの流れ
- 1週目にLCommandでラベルがないか確認する。
- 2周目にACommandの変数を作成する。
## parse処理の変更点
- ACommand時にシンボルか直接アドレスが書いてあるかを判断する。
## 詳細
### symbolTable
#### new
SymbolTable構造体を作る。
```
type SymbolTable struct {
    mp map[string]int
}
```
初期化時に定義済みのやつらをぶち込む。
#### addEntry
`func addEntry(symbol string, address int) error`
mapに追加する。duplicate keyの場合はエラーにする。
#### contains
`func contains(symbol string) bool`
あればtrue, なければfalse
#### getAddress
`func getAddress(symbol string) (int, error)`
あれば返す。ない場合は明示的にエラーを返すのがよさそう
### シンボルの流れ
#### 1週目
isXCommandシリーズをcommandパッケージとして解放する。
初期のidxを0とする。
そして、A, Cコマンド時にはidxをインクリメントする。
Ｌコマンドが来たら、シンボルにそのシンボルを、アドレスにidxの値を入れる。
終わったらcurrentIdxの値をリセット(=0)
#### 2週目
1週目とは違うidxを0にセットする。
Aコマンドが来たら、idxをインクリメントする。
そして、シンボルはそのまま入れて、アドレスには16+idxの値を入れる。
終わったらcurrentIdxの値をリセット(=0)
### parserの変更点
- parserパッケージにシンボリックリンクするメソッドを作成する。
    - 上の1週目、2周目の処理を行うものを作成する。
- toACommandで、valを数値に変換する前にシンボルがあるかどうかを確認し、あったらそのシンボルとアドレスの値を返しちゃう。キャッシュのような感じ。
- type ACommandをValueからAddressに変更する。そっちの方が正しい。
- type LCommandをSymbolのみにする。そっちの方が正しい。
- prepareをNewという名前にして公開する。それにともない、Parseはmethodに切り替える。
- A, LCommandでsymbolの値を返すSymbol methodを用意する。

