package recurstree

type (
	//数据接口类型
	RowInterface interface {
		//判断是否是指定id的子节点数据
		IsChild(int) bool
		//获取数据id
		GetId() int
	}

	//层级结构树
	Tree struct {
		Data     RowInterface
		Children []*Tree
	}
)

//生成层级结构树
//最初的参数node必须是根主节点
func MakeTree(vList []RowInterface, node *Tree) {
	//返回数据集合中某个节点下的直属子节点集合
	childs := findChild(node, vList)
	for _, child := range childs {
		//将直属子节点加入树的当前节点下
		node.Children = append(node.Children, child)
		//判断该子节点下是否还有下属，如果有，继续递归
		if hasChild(child, vList) {
			MakeTree(vList, child)
		}
	}
}

//返回数据集合中某个节点下的直属子节点集合
func findChild(node *Tree, vList []RowInterface) (ret []*Tree) {
	for _, v := range vList {
		if v.IsChild(node.Data.GetId()) {
			ret = append(ret, &Tree{
				Data:     v,
				Children: []*Tree{},
			})
		}
	}
	return ret
}

//判断该子节点下是否还有下属
func hasChild(node *Tree, vList []RowInterface) (has bool) {
	for _, v := range vList {
		if v.IsChild(node.Data.GetId()) {
			has = true
			break
		}
	}
	return has
}
