package nested

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"

	// драйвер работы с postgres
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Component struct {
	Db *sql.DB
}

func GetEnv(t *testing.T) (*Component, context.Context) {
	comp := &Component{}
	strCon := fmt.Sprintf("dbname=%s host=%s port=%d user=%s password=%s sslmode=%s",
		"sample",
		"localhost",
		5432,
		"postgres",
		"postgres",
		"disable",
	)
	db, err := sql.Open("postgres", strCon)
	if err != nil {
		t.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
	comp.Db = db

	return comp, context.Background()
}

func TestTree(t *testing.T) {
	env, ctx := GetEnv(t)

	data := make([]*Set, 0, 100)
	// CREATE
	for i := 0; i < 100; i++ {
		obj := NewSet(env.Db, 0)
		obj, err := obj.Create(ctx, "11111: "+strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
		data = append(data, obj)
		obj, err = obj.Create(ctx, "22222: "+strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
		obj, err = obj.Create(ctx, "33333: "+strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
		obj, err = obj.Create(ctx, "44444: "+strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
		obj, err = obj.Create(ctx, "55555: "+strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
	}

	// MOVE
	for i := range data {
		if i%5 == 0 {
			if err := data[i].Move(ctx, data[i+1]); err != nil {
				t.Fatal(err)
			}
		}
	}

	// COPY
	for i := range data {
		if i%5 == 0 {
			if err := data[i].Copy(ctx, data[i+1]); err != nil {
				t.Fatal(err)
			}
		}
	}

	// DELETE
	for i := range data {
		if i%10 == 0 {
			if err := data[i].Delete(ctx); err != nil {
				t.Fatal(err)
			}
		}
	}

	// GETS
	cnt, err := data[5].GetBranchCountNode(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if cnt != 14 {
		t.Fatal("error: GetBranchCountNode")
	}
	cnt, err = data[5].GetChildNodeCount(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if cnt != 3 {
		t.Fatal("error: GetChildNodeCount")
	}
	d, err := data[5].GetSandPath(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(d) != 1 {
		t.Fatal("error: GetSandPath")
	}

	// CHECK
	tol := NewTool(env.Db)
	if err := tol.Check(ctx); err != nil {
		t.Fatal(err)
	}

	// REPAIR
	if _, err := queries.Raw(
		`UPDATE `+table+` SET keyl = 0, keyr = 0, level = 0`,
	).ExecContext(ctx, env.Db); err != nil {
		t.Fatal(err)
	}
	if err := tol.Repair(ctx); err != nil {
		t.Fatal(err)
	}
	if err := tol.Check(ctx); err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}
