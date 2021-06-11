# go-eth-getblock
## 1. Installation
- Install go from https://golang.org/
- Install dependency libraries
## 2. Run application
Run the terminal and go to the project root folder, and then run the command following.
```
go run server.go
```

## 3. Project structure
this project uses echo web framework.
### 3.1. server.go 
this is the entry file for golang

### 3.2. routes
define routes for 
```
/block/:block
/block/:block/txs/:txs
```

### 3.3. controllers
define controllers matches routings

### 3.4. services
define `EthBlockNumber` to get latest block number and `EthGetBlockByNumber` to get the block information by number.

### 3.5. lrucache
define lrucache module. key is int64

### 3.6. cache
define a global cache variable to use in this project.
global cache variable type is `lrucache.LRUCACHE`

