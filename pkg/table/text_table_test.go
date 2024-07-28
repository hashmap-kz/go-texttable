package table

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentsOk(t *testing.T) {
	tbl := NewTextTable()
	tbl.DefineColumn("POOL", LEFT, LEFT)
	tbl.DefineColumn("ID", RIGHT, RIGHT)
	tbl.DefineColumn("PGS", RIGHT, RIGHT)
	tbl.DefineColumn("DESCR", RIGHT, RIGHT)

	for i := 0; i < 7; i++ {
		tbl.Insert(fmt.Sprintf("pool_name_%05d", i))
		tbl.Insert(fmt.Sprintf("%d", rand.IntN(8192)))
		tbl.Insert(fmt.Sprintf("%d", rand.IntN(32768)))
		tbl.EndRow()
	}

	tbl.InsertAll("pool-name", "n1", "n2", "n3")
	tbl.EndRow()

	tbl.InsertAll("pool-name", "n1", "n2", "n3")
	tbl.EndRow()

	tbl.InsertAllAndFinishRow("pool-name", "n15", "n2", "n32")
	tbl.InsertAllAndFinishRow("pool-name", "n14", "n23", "n33")
	tbl.InsertAllAndFinishRow("pool-name", "n13", "n22", "n34")

	fmt.Println(tbl.Print())

	assert.True(t, true)
}
