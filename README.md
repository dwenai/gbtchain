[![pipeline status](https://api.travis-ci.org/bityuan/bityuan.svg?branch=master)](https://travis-ci.org/bityuan/bityuan/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bityuan/bityuan)](https://goreportcard.com/report/github.com/bityuan/bityuan)

# 基于 chain33 区块链开发 框架 开发的 GBT公有链系统

#### 编译

```
git clone https://github.com/dwenai/gbtchain.git $GOPATH/src/github.com/dwenai/gbtchain
cd $GOPATH/src/github.com/dwenai/gbtchain
go build -i -o gbt
go build -i -o gbt-cli github.com/dwenai/gbtchain/cli
```

#### 运行

拷贝编译好的gbt, gbt-cli, gbtChain.toml这三个文件置于同一个文件夹下，执行：

```
./ts -f gbtChain.toml
```


