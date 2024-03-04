## REST API を FW なしで作成

### movie をリソースとして、URI と HTTP メソッドを定義

| URI         | HTTP method |
| ----------- | ----------- |
| /movies     | GET         |
| /movies/:id | GET         |
| /movies     | POST        |
| /movies/:id | PUT         |
| /movies/:id | DELETE      |

- POST メソッドで新規作成する場合

```shell
curl -X POST -H "Content-Type: application/json" -d '{"id": 4, "name": "movie name", "release": 1995}' localhost:8080/movies
```

- PUT メソッドで更新する場合

```shell
curl -X PUT -H "Content-Type: application/json" -d '{"id": 4, "name": "new movie name!", "release": 1996}' localhost:8080/movies/4
```

- DELETE メソッドで削除する場合

```shell
curl -X DELETE localhost:8080/movies/4
```
