# simpleAddをなんとかする。
simpleAddのアセンブリコードを吐けるような、parserをまずは作る。
コマンドが増えて、膨大になってしまうので
codeWriterは一旦interfaceだけ用意しておいて、次に考える。
動く仕組みを整えて起き、実際にどのようにbyte列に変換するかは後で考える。
## 設計
### process
まずは1つのファイルを読み込む。
ファイル名も拡張子を切り取った形で見る。
生の文字列をParser.Newに渡す。
### parse
parser.Newでまずは行ごとに中間語を区切る。
この際に、コメントのみの行、空行、コマンドの後ろについているコメント(`push constant 7 // 7をプッシュ`)を取り除く。そして、parser.Parse()を実行する。この時に各行をCodeWriter.Command interfaceを満たした構造体を生成して返す。
### codeWriter
codeWriter.Newでファイル名とcommandの配列をもらう。
codeWriter.Writeで書く処理を実行する。
## 具体実装設計
mainはprocess.Runを実行するだけとする
### process package
まずは1つのファイルに対して実行できるように考える。
それができたら、fileかdirかを判定して、dirならファイルごとにfor文で実行する。
#### run
1つのfileに対して実行する。
これは、`.vm`ファイルでなければエラーにする。
.vmなら、読んで中身をそのままparser.Newに投げる。
### parser package
文字列を種類ごとに分けることに責任を持つ。
#### Parser構造体
```
type Parser struct {
    currentIdx // 現在のコマンドのインデックス
    currentCommand // 現在のコマンドを空白区切りで出す
    commands // コマンドの配列
}
```
#### New関数
rawとして、ファイルの本文をもらうようにする。
これは、今後directoryを自体を全部パースできるようにするためである。
currentIdxは0にする。
commandsはこの時に空白のみのものとコメント行を取り除く。
commandsの前後の空白やタブとコメントの行も取り除く。
#### Parser.hasMoreCommand
currentIdxとcommandsの長さ比較
#### Parser.advance
currentIdxのインクリメントをする。
インクリメント後にcurrentCommandに新しいcommandを空白区切りの文字列の配列に変換してセットする。
#### Parser.commandType
コマンドの種類のparser.commandを返す。
具体的にはcurrentCommandの先頭の要素を見て返す。
今回はまずarithmetic(add), pushを抑える。
#### Parser.arg1
parser.command interfaceを受取る。
parser.arithmeticだった場合は、currentCommandの先頭の要素を返す。currentCommand[0]
parser.returnは明示的にerrorを返すようにする。
parser.return以外の場合はcurrentCommand[1]を返す。
#### Parser.arg2
parser.command interfaceを受け取る。
parser.push, parser.pop, parser.function, parser.callの場合はcurrentCommand[2]を返す。
それ以外の場合は明示的にエラーを返すようにする。
#### parser.push.parse
#### parser.alithmetic.parse
#### Parser.Parse
上のメソッドたちを組み合わせる。
最終的にはcodewriter.Commandの配列resutlsを返す
for p.hasMoreCommandで回し、p.commandTypeでparser.command interfaceを生成する。
command.parseでcodeWriter.commandを生成する。
それをresutlsに追加する。
そして、p.advanceで次に進める
####
### codeWriter package
#### CodeWriter構造体
```
type CodeWriter struct {
    file *os.File
    commands []Commmand
}
```
#### New
codeWriter構造体を生成する。
この時にファイル名を受け取っておき、もとのファイル名.asmで作成する。
#### CodeWriter.Write
forループでcommand.Convertを実行していく。出てきたbyte列をCodeWriter.fileに書き込んでいく。
#### CodeWriter.Close
file.Closeを実行するだけ。
#### Command interface
Convertメソッドを持たせる。
これは、書くべきbyte列とerrorを返す。
それぞれの種類ごとに頑張って書く
#### codeWriter.push.Convert
#### codeWriter.alithmetic.Convert