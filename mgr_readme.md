# Mgr Extension

Manager - Mgr extension aims to provide handy db access methods via `Mgr` struct.

Assuming we have a table `books` in database, and the generated package name is `models`, then handy methods will be exposed in `models.BooksMgr`.

For example:

```go
book1, err := models.BooksMgr.GetById(1)
book1, err := models.BooksMgr.GetById(1, "optional selected fields")

book2, err := models.BooksMgr.GetByISBN("11312321")
book2, err := models.BooksMgr.GetByISBN("11312321", "optional selected fields")

books, err := models.BooksMgr.FindByAuthorIDPublishYear(123, 2019).OrderBy("isbn desc").Select("isbn", "author_id").All()
books, err := models.BooksMgr.FindByAuthorIDPublishYear(123, 2019).OrderBy("isbn desc").Select("isbn", "author_id").Paginate(10, 20)
books, total, err := models.BooksMgr.FindByAuthorIDPublishYear(123, 2019).OrderBy("isbn desc").Select("isbn", "author_id").PaginateWithTotal(10, 20)
```

## Handy Methods

### GetByXXX

`GetByXXX` methods returns **one** struct.

These methods are generated for `primary key` and field with `unique index`

* **optional** selected fields is supported

### GetByXXXs

* To be implemented

`GetByXXXs` methods returns map.

These methods are generated for `primary key` and field with `unique index`, and turn to `in` query.

* **optional** selected fields is supported

### FindByXXX

`FindAllByXXX` methods returns **struct slice**.

These methods are generated for indexes in table.

* **optional** selected fields is supported : Select()
* support ordering : OrderBy()
* support general queries : Where()
* finished with All(), Paginate() or PaginateWithTotal()

## Timestamp conversion

sqlboiler can only scan datetime and timestamp value into time.Time or byte[]. However, to prevent time zone affect, we usually pass time value as unix timestamp which is a int64 value between frontend and backend. In this case, developers have to convert the timestamp type manually during development which causes extra efforts.

### solution

We scan datetime or timestamp data into a new struct (types.time) which contains time.Time and its unix value.

## Better transaction

Currently, for every operation in a single transaction, developers have to check error and write rollback or commit code mannually.

### solution

Generate a helper and accepts a callback to handling rollbacks or commit.

```go
// Transact transact handler
func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	err = txFunc(tx)
	return err
}
```

## Sharding

We support database sharding with a yaml configuration file. Assume we have a table `books` and we sharded it into two shards `books_shards_1` and `books_shards_2`. The shard key we use to devide the shards is `author_id`. We can write a `sqlboiler.yaml`:

``` yaml
default: &default
  adapter: mysql
  encoding: utf8mb4
  username: root
  password: root
  master:
    - localhost:3306

tables:
  books:
    shard: true
    shard_key: author_id
    shards:
      - book_shard_1:
          <<: *default
          database: books_shard_1
      - book_shard_2:
          <<: *default
          database: books_shard_2
```

And to query from `books` table, we need to load the configuration from `sqlboiler.yaml` first.

```go
// load sharding config
if err := boil.LoadConfig("sqlboiler.yaml"); err != nil {
    panic(err)
}
// open sql connection
db, err := sql.Open("mysql", "?parseTime=true")
if err != nil {
    panic(err)
}
// set executor
boil.SetDB(db)
// get book with id equals 1
book, err := models.BooksMgr.GetByID(1)
if err != nil {
    panic(err)
}
// get books published by Penguin
books, err := models.Books(qm.Where("publisher=?", "Penguin")).All(db)
if err != nil {
    panic(err)
}
```

The default sharding algorithm is `modulo`.
