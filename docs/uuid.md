# UUID生成方針

## SQLiteにはUUID生成関数がない

PostgreSQLの `gen_random_uuid()` に相当する関数はSQLiteに存在しない。

SQLite側で `lower(hex(randomblob(16)))` を使う方法もあるが、これはUUID形式（ハイフン区切り・バージョンビット）に準拠しておらずただのランダム16進数になるため採用しない。

## 方針: Go側で生成する

`github.com/google/uuid` を使い、`Create()` の中でIDを生成してINSERT時にプレースホルダへ渡す。

```go
import "github.com/google/uuid"

id := uuid.New().String() // "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
```
