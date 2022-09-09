// Code generated by "script/dtogen". DO NOT EDIT.
package daocore

import (
    "context"
    "database/sql"
    "strings"

    "github.com/Masterminds/squirrel"
    "github.com/nnaakkaaii/go-http-server-template/pkg/dberror"
)

const UserTableName = "users"

var UserAllColumns = []string{
    "id",
    "first_name",
    "last_name",
    "email",
    "password",
}

var UserColumnsWOMagics = []string{
    "id",
    "first_name",
    "last_name",
    "email",
    "password",
}

var UserPrimaryKeyColumns = []string{
    "id",
}

type User struct {
    ID string
    FirstName string
    LastName string
    Email string
    Password string
}

func (t *User) Values() []interface{} {
    return []interface{}{
        t.ID,
        t.FirstName,
        t.LastName,
        t.Email,
        t.Password,
    }
}

func (t *User) SetMap() map[string]interface{} {
    return map[string]interface{}{
        "id": t.ID,
        "first_name": t.FirstName,
        "last_name": t.LastName,
        "email": t.Email,
        "password": t.Password,
    }
}

func (t *User) Ptrs() []interface{} {
    return []interface{}{
        &t.ID,
        &t.FirstName,
        &t.LastName,
        &t.Email,
        &t.Password,
    }
}

func IterateUser(sc interface{ Scan(...interface{}) error}) (User, error) {
    t := User{}
    if err := sc.Scan(t.Ptrs()...); err != nil {
        return User{}, dberror.MapError(err)
    }
    return t, nil
}

func SelectOneUserByEmail(ctx context.Context, txn *sql.Tx, email *string) (User, error) {
    eq := squirrel.Eq{}
    if email != nil {
        eq["email"] = *email
    }
    query, params, err := squirrel.
        Select(UserAllColumns...).
        From(UserTableName).
        Where(eq).
        ToSql()
    if err != nil {
        return User{}, dberror.MapError(err)
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return User{}, dberror.MapError(err)
    }
    return IterateUser(stmt.QueryRowContext(ctx, params...))
}

func SelectOneUserByID(ctx context.Context, txn *sql.Tx, id *string) (User, error) {
    eq := squirrel.Eq{}
    if id != nil {
        eq["id"] = *id
    }
    query, params, err := squirrel.
        Select(UserAllColumns...).
        From(UserTableName).
        Where(eq).
        ToSql()
    if err != nil {
        return User{}, dberror.MapError(err)
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return User{}, dberror.MapError(err)
    }
    return IterateUser(stmt.QueryRowContext(ctx, params...))
}



func InsertUser(ctx context.Context, txn *sql.Tx, records []*User) error {
    for i := range records {
        if records[i] == nil {
            records = append(records[:i], records[i+1:]...)
        }
    }
    if len(records) == 0 {
        return nil
    }
    sq := squirrel.Insert(UserTableName).Columns(UserColumnsWOMagics...)
    for _, r := range records {
        if r == nil {
            continue
        }
        sq = sq.Values(r.Values()...)
    }
    query, params, err := sq.ToSql()
    if err != nil {
        return err
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return dberror.MapError(err)
    }
    if _, err = stmt.Exec(params...); err != nil {
        return dberror.MapError(err)
    }
    return nil
}

func UpdateUser(ctx context.Context, txn *sql.Tx, record User) error {
    sql, params, err := squirrel.Update(UserTableName).SetMap(record.SetMap()).
        Where(squirrel.Eq{
        "id": record.ID,
    }).
        ToSql()
    if err != nil {
        return err
    }
    stmt, err := txn.PrepareContext(ctx, sql)
    if err != nil {
        return dberror.MapError(err)
    }
    if _, err = stmt.Exec(params...); err != nil {
        return dberror.MapError(err)
    }
    return nil
}

func UpsertUser(ctx context.Context, txn *sql.Tx, record User) error {
    updateSQL, params, err := squirrel.Update(UserTableName).SetMap(record.SetMap()).ToSql()
    if err != nil {
        return err
    }
    updateSQL = strings.TrimPrefix(updateSQL, "UPDATE "+UserTableName+" SET ")
    query, params, err := squirrel.Insert(UserTableName).Columns(UserColumnsWOMagics...).Values(record.Values()...).SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE "+updateSQL, params...)).ToSql()
    if err != nil {
        return err
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return dberror.MapError(err)
    }
    if _, err = stmt.Exec(params...); err != nil {
        return dberror.MapError(err)
    }
    return nil
}

func DeleteOneUserByEmail(ctx context.Context, txn *sql.Tx, email *string) error {
    eq := squirrel.Eq{}
    eq["email"] = email

    query, params, err := squirrel.
        Delete(UserTableName).
        Where(eq).
        ToSql()
    if err != nil {
        return dberror.MapError(err)
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return dberror.MapError(err)
    }
    if _, err = stmt.Exec(params...); err != nil {
        return dberror.MapError(err)
    }
    return nil
}

func DeleteOneUserByID(ctx context.Context, txn *sql.Tx, id *string) error {
    eq := squirrel.Eq{}
    eq["id"] = id

    query, params, err := squirrel.
        Delete(UserTableName).
        Where(eq).
        ToSql()
    if err != nil {
        return dberror.MapError(err)
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return dberror.MapError(err)
    }
    if _, err = stmt.Exec(params...); err != nil {
        return dberror.MapError(err)
    }
    return nil
}
