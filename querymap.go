package mapq

// QueryMap 查询map
func QueryMap(data map[string]interface{}, query string) (bool, error) {
	p := &Parser{}
	n, err := p.Parse(query)
	if err != nil {
		return false, err
	}
	return n.Eval(data).(bool), nil
}

// RunQuery 查询
func RunQuery(root Node, data map[string]interface{}) (bool, error) {
	return root.Eval(data).(bool), nil
}
