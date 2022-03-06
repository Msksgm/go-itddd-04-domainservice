# GO-ITDDD-04-DOMAINSERVICE

zenn の記事「[Go でドメインサービスを実装（「入門ドメイン駆動設計」Chapter4））](https://zenn.dev/msksgm/articles/20220311-go-itddd-04-domainservice)」のサンプルコードです。

# 実行環境

- Go
  - 1.17
- docker compose

# 実行方法

## コンテナを起動・マイグレーション

コンテナの起動

```bash
> make up
docker compose up -d
# 完了まで待つ
```

マイグレーション

```bash
> docker compose exec app bash db-migration.sh
1/u user (12.838ms)
```

## 実行

test-user 登録 1 回目

```bash
> make run
docker compose exec app go run main.go
2022/03/05 10:20:04 successfully connected to database
2022/03/05 10:20:04 test-user is successfully added in users table
```

test-user 登録 2 回目

```bash
> make run
docker compose exec app go run main.go
2022/03/05 10:22:24 successfully connected to database
2022/03/05 10:22:24 main.CreateUser(): test-user is already exists.
```

# テスト

```bash
> make test
docker compose exec app go test ./...
ok      github.com/Msksgm/itddd-go-04-domainservice     0.002s
ok      github.com/Msksgm/itddd-go-04-domainservice/domain/model/transport      0.003s
ok      github.com/Msksgm/itddd-go-04-domainservice/domain/model/user   0.003s

```
