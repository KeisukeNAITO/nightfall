# Go製Hello World APIの起動メモ

テスト:

```bash
go test ./...
go test -cover ./...
```

実行:

```bash
go run .
```

APIアクセス方法:

```bash
curl "http://localhost:8080/api/v1/hello"
# {"message":"hello world"}
```

double:

```bash
curl "http://localhost:8080/api/v1/double/21"
# {"input":21,"double":42}
```
