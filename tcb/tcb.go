package tcb

import "github.com/imokyou/wechat-plus/context"

//Tcb Tencent Cloud Base
type Tcb struct {
	*context.Context
}

//NewTcb new Tencent Cloud Base
func NewTcb(context *context.Context) *Tcb {
	return &Tcb{
		context,
	}
}