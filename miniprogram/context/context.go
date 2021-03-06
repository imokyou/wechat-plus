package context

import (
	"github.com/imokyou/wechat-plus/v2/credential"
	"github.com/imokyou/wechat-plus/v2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
