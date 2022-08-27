package mapq

import "reflect"

// Node 节点
type Node interface {
	Eval(data map[string]interface{}) interface{}
}

// 这些函数提供给你，也许可以帮上忙。。。
func toF64(i interface{}) float64 {
	switch v := i.(type) {
	case int:
		return float64(v)
	case float64:
		return v
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	}
	return 0
}
func trytoF64(i interface{}) interface{} {
	switch v := i.(type) {
	case int, float64, int32, int64, uint32, uint64, float32:
		return toF64(v)

	}
	return i
}

func equal(left, right interface{}) bool {
	return reflect.DeepEqual(trytoF64(left), trytoF64(right))
}

// BinNode 双目表达式节点
type BinNode struct {
	Left, Right Node
	Op          int
}

// Eval 查询
func (n *BinNode) Eval(data map[string]interface{}) interface{} {
	panic("not implemented")
}

// 别的节点。。。。
