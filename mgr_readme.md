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

* selected fields is supported : Select()
* support ordering : OrderBy()
* support general queries : Where()
* finished with All(), Paginate() or PaginateWithTotal()

## Better transaction

Currently, for every operation in a single transaction, developers have to check error and write rollback or commit code mannually.

### Helper function

We provide a helper and accepts a callback to handling rollbacks or commit.

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

### Usage

Take a simple online transaction as example:

* Without helper

```go
// start transaction
tx, err := db.Begin()
if err != nil {
    panic(err)
}
// insert order
if err := order.Insert(tx, boil.Infer()); err != nil {
   tx.Rollback()
   panic(err)
}
// insert invoice
if err := invoice.Insert(tx, boil.Infer()); err != nil {
   tx.Rollback()
   panic(err)
}
// insert payment
if err := payment.Insert(tx, boil.Infer()); err != nil {
   tx.Rollback()
   panic(err)
}
err = tx.Commit()
if err != nil {
    panic(err)
}
```

* With helper

```go
// start transaction
err := boil.Transact(db, func(tx *sql.Tx) error {
    // insert order
    if err := order.Insert(tx, boil.Infer()); err != nil {
        return err
    }
    // insert invoice
    if err := invoice.Insert(tx, boil.Infer()); err != nil {
        return err
    }
    // insert payment
    err := payment.Insert(tx, boil.Infer())
    return err
})
if err != nil {
    panic(err)
}
```

It reduces the redundant `tx.Rollback()` or `tx.Commit()` code.

## To be implemented

* custom type in 'IN' query
* select with pagination with condition
* mysql functions
