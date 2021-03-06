package model

import (
	"fmt"
)

// 转化列名, 防止报错
// ref https://github.dev/Xhofe/alist model/util
func columnName(name string) string {
	return fmt.Sprintf("`%s`", name)
}
