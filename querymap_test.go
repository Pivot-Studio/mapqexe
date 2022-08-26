package mapq

import (
	"encoding/json"
	"testing"
)

func TestQueryMap(t *testing.T) {
	type args struct {
		data  map[string]interface{}
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// 基础功能测试
		{"basic true case", args{map[string]interface{}{"a": 1}, "a==1"}, true, false},
		{"basic false case", args{map[string]interface{}{"a": 1}, "a==2"}, false, false},
		{"and true case", args{map[string]interface{}{"a": 1, "b": 2}, "a==1&&b==2"}, true, false},
		{"and false case", args{map[string]interface{}{"a": 1, "b": 2}, "a==1&&b==3"}, false, false},
		{"or true case", args{map[string]interface{}{"a": 1, "b": 2}, "a==1||b==2"}, true, false},
		{"or true case2", args{map[string]interface{}{"a": 1, "b": 2}, "a==1||b==3"}, true, false},
		{"or false case", args{map[string]interface{}{"a": 1, "b": 2}, "a<=0||b>3"}, false, false},
		{"LEQ SM case", args{map[string]interface{}{"a": 1, "b": 2}, "a<2&&b>=2"}, true, false},
		{"nested case", args{map[string]interface{}{"a": 1, "b": 2, "c": map[string]interface{}{"d": 3}},
			"a==1&&b==2&&c.d==3"}, true, false},
		{"NEQ case", args{map[string]interface{}{"a": 1, "b": 2}, "a!=1||b!=2"}, false, false},
		{"parentheses case", args{map[string]interface{}{"a": 1, "b": 2}, "a==1&&!(b==2||b==3)"}, false, false},
		// 高级功能测试
		{"add case", args{map[string]interface{}{"a": 1, "b": 2}, "a+b==3"}, true, false},
		{"add case false", args{map[string]interface{}{"a": 1, "b": 2}, "a+b<3"}, false, false},
		{"mul case", args{map[string]interface{}{"a": 3, "b": 2}, "a*b==6"}, true, false},
		{"div case", args{map[string]interface{}{"a": 3, "b": 2}, "a/b==1.5"}, true, false},
		{"parentheses mul case", args{map[string]interface{}{"a": 3, "b": 2}, "(a+b)*b==10"}, true, false},
		{"parentheses mul false case", args{map[string]interface{}{"a": 3, "b": 2}, "(a+b)*b<5"}, false, false},
		{"string case", args{map[string]interface{}{"a": "1", "b": "2"}, "a=='1'&&b=='2'"}, true, false},
		{"null case", args{map[string]interface{}{"a": 1}, "b==null&&a!=null"}, true, false},
		{"unary case", args{map[string]interface{}{"a": 1}, "!(a==+1)&&-a==-1"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryMap(tt.args.data, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QueryMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRunQuery(b *testing.B) {
	obj := `
	{
		"component_id": "82e2e6e6db5e837242e9314084ec3cf7",
		"component_name": "文本",
		"component_type": "TEXT",
		"component_required": true,
		"component_recipient_id": "a9ee75702212e12f27cf1462b04b4df1",
		"component_width": 60,
		"component_height": 18,
		"component_page": 1,
		"component_pos_x": 235,
		"component_pos_y": 155,
		"component_extra": {"FontSize":12}
	}`
	datamap := make(map[string]interface{})
	json.Unmarshal([]byte(obj), &datamap)
	p := &Parser{}
	root, _ := p.Parse("!(component_pos_x>300||component_pos_y<160&&component_extra.FontSize==12)||component_type=='TEXT'")
	b.Run("query bench", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			_, err := RunQuery(root, datamap)
			if err != nil {
				b.Error(err)
			}
		}
	})
}
