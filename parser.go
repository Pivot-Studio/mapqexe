package mapq

// Parser 语法分析器
type Parser struct {
	lexer *Lexer
}

// 你的递归下降分析代码（如果你使用递归下降的话
// func (p *Parser) boolexp() (node Node, err error) {
// 	panic("not implemented")
// }

// func (p *Parser) boolean() (node Node, err error) {
// 	panic("not implemented")
// }
// 别的分析函数
// 。。。。。

// Parse 生成ast
func (p *Parser) Parse(str string) (n Node, err error) {
	panic("not implemented")
}
