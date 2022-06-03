package constant

const (
	COMMON int32 = (1 << 10) * 0 // 0
)

var (
	SUCCESS = &BaseStatus{StatusCode: COMMON + 0, StatusMsg: "", ShowMsgKey: ""}

	ERR_UNKNOW = &BaseStatus{StatusCode: COMMON + 1, StatusMsg: "未知错误", ShowMsgKey: ""}

	ERR_SERVICE_INTERNAL = &BaseStatus{StatusCode: COMMON + 2, StatusMsg: "啊哦，服务器打瞌睡了", ShowMsgKey: "error_message_comment_post_fail"}
	ERR_INVALID_PARAM    = &BaseStatus{StatusCode: COMMON + 3, StatusMsg: "参数不合法", ShowMsgKey: "invalid_parameters"}

	ERR_USER_CLOSE_CONN = &BaseStatus{StatusCode: COMMON + 100, StatusMsg: "用户断开连接", ShowMsgKey: ""}
)

type BaseStatus struct {
	StatusCode int32
	StatusMsg  string
	ShowMsgKey string
}

func (this *BaseStatus) GetStatusCode() int32 {
	return this.StatusCode
}

func (this *BaseStatus) GetError() (int32, string) {
	return this.StatusCode, this.StatusMsg
}
