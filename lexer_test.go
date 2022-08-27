package mapq

import "testing"

func TestLexer_ScanType(t *testing.T) {
	type fields struct {
		input string
		pos   int
		runes []rune
	}
	type args struct {
		code int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantToken string
		wantErr   bool
	}{
		{"add case", fields{input: "+456"}, args{code: TYPE_PLUS}, "+", false},
		{"int case", fields{input: "123"}, args{code: TYPE_INT}, "123", false},
		{"true case", fields{input: "true"}, args{code: TYPE_RES_TRUE}, "true", false},
		{"float case", fields{input: "123.456"}, args{code: TYPE_FLOAT}, "123.456", false},
		{"string case", fields{input: "'123_dsaa--sdada'"}, args{code: TYPE_STR}, "123_dsaa--sdada", false},
		{"var case", fields{input: "a10"}, args{code: TYPE_VAR}, "a10", false},
		{"var case 2", fields{input: "_a10_bA"}, args{code: TYPE_VAR}, "_a10_bA", false},
		{"larger case", fields{input: ">a10bA"}, args{code: TYPE_LG}, ">", false},
		{"combined op case", fields{input: ">="}, args{code: TYPE_LEQ}, ">=", false},
		{"err case", fields{input: ">"}, args{code: TYPE_LEQ}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("panic error: %v", err)
				}
			}()
			l := &Lexer{
				pos:   tt.fields.pos,
				runes: tt.fields.runes,
			}
			l.SetInput(tt.fields.input)
			gotToken, err := l.ScanType(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lexer.ScanType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("Lexer.ScanType() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}
