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
