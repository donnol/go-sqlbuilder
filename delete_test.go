// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"fmt"
	"testing"
)

func ExampleDeleteFrom() {
	sql := DeleteFrom("demo.user").
		Where(
			"status = 1",
		).
		Limit(10).
		String()

	fmt.Println(sql)

	// Output:
	// DELETE FROM demo.user WHERE status = 1 LIMIT 10
}

func ExampleDeleteBuilder() {
	db := NewDeleteBuilder()
	db.DeleteFrom("demo.user")
	db.Where(
		db.GreaterThan("id", 1234),
		db.Like("name", "%Du"),
		db.Or(
			db.IsNull("id_card"),
			db.In("status", 1, 2, 5),
		),
		"modified_at > created_at + "+db.Var(86400), // It's allowed to write arbitrary SQL.
	)

	sql, args := db.Build()
	fmt.Println(sql)
	fmt.Println(args)

	// Output:
	// DELETE FROM demo.user WHERE id > ? AND name LIKE ? AND (id_card IS NULL OR status IN (?, ?, ?)) AND modified_at > created_at + ?
	// [1234 %Du 1 2 5 86400]
}

func ExampleDeleteBuilder_SQL() {
	db := NewDeleteBuilder()
	db.SQL(`/* before */`)
	db.DeleteFrom("demo.user")
	db.SQL("PARTITION (p0)")
	db.Where(
		db.GreaterThan("id", 1234),
	)
	db.SQL("/* after where */")
	db.OrderBy("id")
	db.SQL("/* after order by */")
	db.Limit(10)
	db.SQL("/* after limit */")

	sql, args := db.Build()
	fmt.Println(sql)
	fmt.Println(args)

	// Output:
	// /* before */ DELETE FROM demo.user PARTITION (p0) WHERE id > ? /* after where */ ORDER BY id /* after order by */ LIMIT 10 /* after limit */
	// [1234]
}

func TestDelete(t *testing.T) {
	{
		db := NewDeleteBuilder().
			DeleteFrom("demo.user").
			WhereCondition(ValueFunc(Equal), "id", 1234).
			WhereCondition(ValueFunc(Equal), "name", "xx")
		query, args := db.Build()
		t.Logf("%s, %+v", query, args)
	}
	// {
	// 	db := NewDeleteBuilder()
	// 	db.DeleteFrom("demo.user").
	// 		WhereCondition(ValuesFunc(db.In), "id", 1234, 2235).
	// 		WhereCondition(ValuesFunc(db.In), "age", 12, 22)
	// 	query, args := db.Build()
	// 	t.Logf("%s, %+v", query, args)
	// }
	// {
	// 	db := NewDeleteBuilder()
	// 	db.DeleteFrom("demo.user").
	// 		WhereCondition(Func(db.IsNotNull), "id").
	// 		WhereCondition(Func(db.IsNull), "addr")
	// 	query, args := db.Build()
	// 	t.Logf("%s, %+v", query, args)
	// }
}
