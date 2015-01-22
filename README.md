#democopy
##what's this?
Sourceエンジン製のデモファイルをzip圧縮して  
WebServerのディレクトリに日時別に配置するツールです  
##How to use.
1. [Golang](http://golang-jp.org/doc/install)をインストールし`${GOPATH}`を設定してください
2. `$ go get github.com/aoisensi/democopy` コマンドを実行してください
3. cronなどで  
`$ democopy -s="/path/to/demo/files/directory" -d="/path/to/web/server/directory"`  
を定期実行してください
4. 最新のファイルはゲーム中の可能性があるのでzip圧縮しません

##License
[![WTFPL](http://www.wtfpl.net/wp-content/uploads/2012/12/wtfpl-badge-4.png)](http://www.wtfpl.net/)
