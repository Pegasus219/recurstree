package example

import (
	"encoding/json"
	"fmt"
	tree "github.com/Pegasus219/recurstree"
	"testing"
)

type Org struct {
	Id   int
	Pid  int
	Name string
}

func (o *Org) IsChild(id int) (b bool) {
	return o.Pid == id
}

func (o *Org) GetId() (i int) {
	return o.Id
}

func TestTree(t *testing.T) {
	orgList := []tree.RowInterface{
		&Org{Id: 12, Pid: 0, Name: "1st Layer"},
		&Org{Id: 131, Pid: 12, Name: "2nd Layer(A)"},
		&Org{Id: 132, Pid: 12, Name: "2nd Layer(B)"},
		&Org{Id: 7634, Pid: 131, Name: "3rd Layer"},
	}
	orgData, _ := json.Marshal(orgList)
	fmt.Printf("原始数据：%s\n", orgData)

	nodeTree := &tree.Tree{
		Data:     orgList[0], //根主节点
		Children: []*tree.Tree{},
	}
	tree.MakeTree(orgList, nodeTree)
	data, _ := json.Marshal(nodeTree)
	fmt.Printf("生成树：%s", data)
}
