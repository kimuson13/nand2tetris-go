# アセンブリ作ろう
## 処理の手順(シンボル無し)
- [ ] .asmファイルを読み込む
- [ ] .hackファイルを作成する
- [ ] parserで構文解析を行い、astの形式に落とし込む。
- [ ] astの内容を、もとに機械語に変換(codeパッケージ)
- [ ] .hackファイルを作成し、機械語を.hackファイルに書き込む。
## 詳しい設計
### .asmファイルを読み込む
- .asmファイル以外はエラーにする。
- パスがない場合は、os.Openのエラーをそのまま返す。
### parserで構文解析を行う。
Parser構造体を持つ。
これに仕様にあるように
- init
- hasMoreCommands
- advance
- commandType
- symbol
- dest
- comp
- jump
をメソッドとして持たせる。
それぞれのメソッドの設計はパッケージ構成に書く。
ここでは処理の手順を考える。

まずprepareでfileの内容をParser構造体でparseできるような形にする。
hasMoreCommandsがtrueの間はforで回して、1行ずつ処理していく。
最終的な戻り値は、Code.ACommand構造体かCode.CCommand構造体のスライスにする。
parserとしての戻り値はCommand interfaceを用意する。
次に、`commandType`でなんのコマンドかを判定する。
ここでの戻り値によって処理が変わる。
A_COMMANDのときは、symbolを実行して、構造体にする。
C_COMMANDのときは、dest, comp, jumpメソッドを実行して構造体に詰める。
その行の処理が終わったら、advanceメソッドを呼ぶ。

### 解析結果をもとに変換する。
CodeパッケージにCommand interfaceを用意して、ACommand, CCommandでconvertメソッドを用意しておく。それをforで回して順にparserの結果から変換していく。

### .hackファイルを作成し、機械語を.hackファイルに書き込む。
.asmファイルの接頭辞？を取る。Max.asmならMax.hackファイルを作成する。
どのような形式で受け取るかは解像度を上げてから決める。
ファイルの生成場所は、読み込んだファイルと同じ位置とする。
`/usr/local/hoge/Hoge.asm`が来た場合は`/usr/local/hoge/Hoge.hack`とする。なので、引数ではコマンドのargsの1つ目をそのまま受け取る。
`file.CreateHack(path string, hoge hoge?)`
## パッケージ構成
書いている内に分けたくなったら分ける。
### parser
#### Parser.Init()
フィールドは
- Commands
- CurrentIdx
を持つ。  
Commandsは読み込んだアセンブリファイルの中身を改行区切りで保持する。
CurrentIdxは現在の処理している行数。
#### hasMoreCommands
Commandsの長さがCurrentIdxより大きければtrue, ほかはfalse
#### advance
currentIdxを足すだけ。
#### commandType
その行の先頭の文字を見る。
`@`ならA_COMMAND,
`(`ならL_COMMANDを返す。
それ以外の場合、
- `=`と`;`が両方1つ含まれており、`=`が先に来る
- `=`が1つ含まれている。
- `;`が1つ含まれている
の3パターンであるかを確認し、そうでなければエラーを返す。

A_COMMAND, L_COMMAND, C_COMMANDはそれぞれ構造体として返す。
A_COMMAND, L_COMMANDはRawを持たせる。
C_COMMANDには、Comp, Dest, Jumpを持たせる。ここでは、それぞれを文字列として分けるだけにする。
そしてこれらをCommand interfaceとして返す。メソッドはString()を持たせておく。
A_COMMAND, L_COMMANDにはSymboler interfaceにも属させる。これはsymbolというメソッドを持つ。symbolはA_COMMAND, L_COMMANDのそれぞれのCode.Command interfaceに変換する役割を持つ。これでポリモーフィズムが保たれるはず。。。
#### symbol
引数にはsymboler interfaceを持たせて、ポリモーフィズムで実行する。
詰め替えて返すだけに近い。
#### dest, comp, jump
文字列を元にマッピングする。
### file
.hackファイルを作成し、書き込む処理を行う。
パスからfile名を抽出するのもここで行う。
stringのスライスを行ごとに最後に改行コードを付けて書き込む。
### code
Command interfaceとA, L, Cの実装を持つ。
それぞれ対応をハードコードする。。。
元気があればきちんと定数で定義して、それをマッピングできるような形を取りたい。。。
あとは、Commandのsliceを順に実行する関数を用意する。
これはstringのスライスとして返す。
### process
`process.Run(arg string) error`でメインの処理を行う。
mainパッケージはこれのみを呼び出すようにする。
process.Runのargは`go run main.go tests/max/MaxL.asm`のように1つ目の引数を受け取る。
引数のエラーはここで処理をする。
- 引数が多い。
- 引数の形式がファイルパスではない。
ことを処理する。

mainは本当にprocess.Runを実行して、エラーを返して終了するだけとする。
