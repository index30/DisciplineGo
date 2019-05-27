## DisciplineGo
To be better Gopher  
高みを目指せたらいいなあという気持ちで始めたGoのあれこれ

## Contents
Pile basic skill up  
少しずつ積み重ねます
- basic_tips
    - `hello_world.go` ... Hello, World!!  
      Hello, Worldは飛ばす派です  
      print文系を盛り込んだだけ
    - `fmt_etc.go` ... Trying Stringer etc.  
    fmt系の色々を試す場所
    - `builder_pattern.go` ... Builder Pattern  
    ダックタイピングに関わるあれこれ  
    Builderを介して構築することで、ビルド時の抜け漏れや新たな要件を定義しやすくしている  
    加えて, BuilderやInterfaceによって実際のコードを隠蔽でき, テストも簡単になる  
    - `generator_pattern.go` ... Generator Pattern  
    ジェネレータについて
    - `make_cli/` ... About CLI Tool([link](#CLI))  
    メニューを入力するとランダムに返すCLI ([1]参照)
    - `process_log/` ... CLI Tool Part2
    プロセス実行ログCLI([1]参照)
    - `echo_app_mock/` ... Check about "echo"
    フレームワークEchoの確認. Hello Worldしていきましょう()

## Memorandum
### Etiquette
- `go fmt`: Most action before commit  
    Goのオフィシャルフォーマッタの適用
- `go lint`: Suggesting points that you should fix
    推奨される修正の提案

### Go Patterns
[GoPatterns](http://tmrts.com/go-patterns/)は慣用的な実践パターンのtips集

### Lint Tool
- [errcheck](https://github.com/kisielk/errcheck)
- [maligned](https://github.com/mdempsky/maligned)

### CLI

```
$ mkdir cli_dir
$ cd cli_dir
$ go mod init cli_dir 
$ (make file in cli_dir)
$ go build
```

## Author
[Yusei](https://github.com/index30)

## Reference
[[1](https://gihyo.jp/magazine/SD/archive/2019/201905)] Software Design 2019年5月号  