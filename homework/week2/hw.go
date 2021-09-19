package main

import (
	"github.com/pkg/errors"
)

func main() {
	// 如果想要降级处理的话就直接处理，不把错误向上抛
	// 如果不处理需要直接跑到上层，在需要处理的地方通过业务逻辑处理
	err := errors.Wrap(errors.NotFound, "not found")
}
