package result

type Codes struct {
	SUCCESS            uint
	FAILED             uint
	Message            map[uint]string
	NOAUTH             uint
	AUTHFORMATERROR    uint
	MissingLoginParams uint
	CodeExpired        uint
	CaptStatus         uint
	PasswordError      uint
	StatusEnabled      uint
	POSTEXISTS         uint
}

var ApiCode = &Codes{
	SUCCESS:            200,
	FAILED:             501,
	NOAUTH:             403,
	AUTHFORMATERROR:    405,
	MissingLoginParams: 407,
	CodeExpired:        408,
	CaptStatus:         409,
	PasswordError:      410,
	StatusEnabled:      411,
	POSTEXISTS:         412,
}

func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:            "成功",
		ApiCode.FAILED:             "失败",
		ApiCode.NOAUTH:             "身份认证失败",
		ApiCode.AUTHFORMATERROR:    "身份认证失败",
		ApiCode.MissingLoginParams: "缺少登录必须参数",
		ApiCode.CodeExpired:        "验证码过期",
		ApiCode.CaptStatus:         "验证码错误",
		ApiCode.PasswordError:      "账号或密码错误",
		ApiCode.StatusEnabled:      "该账号已被禁用",
		ApiCode.POSTEXISTS:         "岗位已存在",
	}
}

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
