package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/test"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Test    test.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
