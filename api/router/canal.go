package router

import (
	"github.com/withlin/canal-go/protocol/entry"
)

type Canal struct {
	FuncMap map[string]func([]*entry.Column)
}

// Init 表名@方法 如 pc_acitivity@Insert
func (c *Canal) Init() {

}
