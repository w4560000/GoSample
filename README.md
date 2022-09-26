# GoSample

### GO 官網下載
```cmd
https://go.dev/doc/install
```

### VSCODE debug 安裝
```cmd
go get github.com/go-delve/delve/cmd/dlv
go install github.com/go-delve/delve/cmd/dlv
```

### GO 筆記

#### ENV
GOROOT = Golang 語言內建的程式庫、內建 package 放置路徑
GOPATH = 放置第三方 package 路徑
GO111MODULE  = on (程式引用 package 時使用 modules，而不會到 GOPATH 底下找)
               off (程式引用 package 時不使用 modules，一樣到 GOPATH 底下找)
               auto (自動判別，當前或上一層目錄存在 go.mod 以及 該專案目錄不在 GOPATH/src/ 下時為on，否則為off )
               設成on較好，因該功能開發出來就是為了解決 GOPATH/src 底下套件很複雜的問題

#### CMD
go run = 直譯式執行 (開發測試時 用 go rum 執行測試即可)
go build = 將程式碼編譯成 .exe檔
go install = 將套件編譯成 .a檔
go get  = 若有開啟 module，則會將第三方 package 下載至 GOPATH/pkg/mod
          若沒有則下載 至 GOPATH/src 下
### gin 安裝
```cmd
go get github.com/gin-gonic/gin
```

### vscode debug
修改 .vscode => lauch.json
program 指到要debug的檔案

### go package
一開始

GO111MODULE = 