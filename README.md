# 功能
MySQL数据库表和字段差异对比

# 安装
go get github.com/sylvat/mysqldiff

# 配置
## 配置文件configs/mysqldiff.yaml
```yaml
mysql:
  dev:
    dsn: dev:123456@tcp(x.x.x.x:3306)/test?timeout=90s&collation=utf8mb4_unicode_ci&parseTime=true
    db: test
  test:
    dsn: test:123456@tcp(x.x.x.x:3306)/test?timeout=90s&collation=utf8mb4_unicode_ci&parseTime=true
    db: test
```
## .env配置（可选）
可以在项目根目录下新建.env文件，env配置会覆盖mysqldiff.yaml中的配置
```
MYSQL_DEV_DSN=dev:123456@tcp(x.x.x.x:3306)/test?timeout=90s&collation=utf8mb4_unicode_ci&parseTime=true
MYSQL_DEV_DB=test
MYSQL_TEST_DSN=dev:123456@tcp(x.x.x.x:3306)/test?timeout=90s&collation=utf8mb4_unicode_ci&parseTime=true
MYSQL_TEST_DB=test
```

# 使用
```
$ go build
$ ./mysqldiff help  # 查看使用说明
$ ./mysqldiff run --src=dev --dest=test  # 比较源库dev和目标库test之间差异,dev和test需要在配置文件加入对应配置
$ ./mysqldiff run   # 默认src=dev,dest=test
```
