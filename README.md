# leiserv
## 安装
```shell
go mod  tidy
```

## 启动
```shell
go run main.go
```
## 打包
```shell
go build -o leiserv main.go
```
## 运行
```shell
./leiserv -c config.yaml
```

## pm2

```shell
pm2 start package.json
```

### 启动项目
```shell
go run main.go
```
### 打包项目
```shell
go build -o leiserv main.go
```

### 使用 pm2 启动项目
```shell
pm2 start package.json --name leiserv
```
### package.json
```json
{
  "name": "leiserv",
  "version": "1.0.0",
  "description": "leiserv",
  "main": "main.go",
  "scripts": {
    "start": "go run main.go"
  },
  "author": "Richard",
  "license": "MIT"
}
```