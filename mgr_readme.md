# Mgr Extension

Manager - Mgr extension aims to provide handy db access methods via `Mgr` struct.

Assuming we have a table `books` in database, and the generated package name is `models`, then handy methods will be exposed in `models.BooksMgr`.

For example:

```go
book1, err := models.BooksMgr.GetById(1)
book1, err := models.BooksMgr.GetById(1, "optional selected fields")

book2, err := models.BooksMgr.GetByISBN("11312321")
book2, err := models.BooksMgr.GetByISBN("11312321", "optional selected fields")

books, err := models.BooksMgr.FindAllByAuthorID(123)
books, err := models.BooksMgr.FindAllByAuthorID(123, "optional selected fields")

books, err := models.BooksMgr.FindByAuthorID(author_id, limit, offset, "optional selected fields")
books, err := models.BooksMgr.FindByAuthorIDPublishYear(123, 2019, limit, offset, "optional selected fields")
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

### FindAllByXXX

`FindAllByXXX` methods returns **struct slice**.

These methods are generated for indexes in table.

* **optional** selected fields is supported
* ordering support should be implemented

### FindByXXX

`FindByXXX` methods returns **struct slice**, its similar to `FindAllByXXX`, but with pagination support via limit / offset parameters.

* ordering support should be implemented
