package context

import (
	"github.com/zhangxa/wechat/v2/credential"
	"github.com/zhangxa/wechat/v2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenContextHandle
}
