# shisha-log-backend

### OpenAPI
<https://youngeek-0410.github.io/shisha-log-backend/>

### DB Diagram
<https://dbdiagram.io/d/65a3fb0dac844320aee1a1b0>

### データ登録手順
- shisha-log-backendディレクトリに移動
- `docker-compose exec -it api-server sh`
- `cd migration/`
- `go run main.go up` で↓が出ればオッケー
```
start migration
Before:
-------------------
version :  0
dirty   :  false
error   :  no migration
-------------------

Updated:
-------------------
version :  20240111045838
dirty   :  false
error   :  <nil>
-------------------
```
- `exit`
- `docker-compose exec -iT db mysql -uroot shisha-log < migration/seed.sql`

※マイグレーションのバージョンを戻すとき
- `docker-compose exec -it api-server sh`
- `cd migration/`
- `go run main.go down`