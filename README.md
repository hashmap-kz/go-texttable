# Pretty printing utility with table-like style

### Usage:
```
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
```


### Result:
```
POOL               ID    PGS  DESCR
pool_name_00000  4924  22903       
pool_name_00001  2092  10713       
pool_name_00002  8143  26624       
pool_name_00003  6975   1490       
pool_name_00004  6177    263       
pool_name_00005   773   1439       
pool_name_00006  3295   8256       
pool-name          n1     n2     n3
pool-name          n1     n2     n3
pool-name         n15     n2    n32
pool-name         n14    n23    n33
pool-name         n13    n22    n34
```