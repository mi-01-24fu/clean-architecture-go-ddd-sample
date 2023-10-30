# clean-architecture-go-ddd-sample

## DB設定
sql-migrate ファイル作成
```
$ cd config
$ sql-migrate <fileName>
```
sql-migrate create table
```
$ cd config
$ sql-migrate up
```
sql-migrate drom table
```
$ cd config
$ sql-migrate down
```
sql-boiler 実行
```
$ cd config
$ sqlboiler psql
```