package result

type Codes struct {
	SUCCESS             uint
	FAILED              uint
	Message             map[uint]string
	NOAUTH              uint
	AUTHFORMATERROR     uint
	MissingLoginParams  uint
	CodeExpired         uint
	CaptStatus          uint
	PasswordError       uint
	StatusEnabled       uint
	POSTEXISTS          uint
	DEPTEXISTS          uint
	DEPTDELETEFAILED    uint
	MENUEXISTS          uint
	DELETEMENUFAILED    uint
	ADDROLEFAILED       uint
	MissParameter       uint
	CREATEUSERFAILED    uint
	UPLOADFILEFAILED    uint
	UPDATEUSERFAILED    uint
	RESETPASSWORDFAILED uint
	TWICEPASSWORDERROR  uint
	OLDPASSWORDERROR    uint
	TOKENINLAILD        uint
}

var ApiCode = &Codes{
	SUCCESS:             200,
	FAILED:              501,
	NOAUTH:              403,
	AUTHFORMATERROR:     405,
	MissingLoginParams:  407,
	CodeExpired:         408,
	CaptStatus:          409,
	PasswordError:       410,
	StatusEnabled:       411,
	POSTEXISTS:          412,
	DEPTEXISTS:          413,
	DEPTDELETEFAILED:    414,
	MENUEXISTS:          415,
	DELETEMENUFAILED:    416,
	ADDROLEFAILED:       417,
	MissParameter:       418,
	CREATEUSERFAILED:    419,
	UPLOADFILEFAILED:    420,
	UPDATEUSERFAILED:    421,
	RESETPASSWORDFAILED: 422,
	TWICEPASSWORDERROR:  423,
	OLDPASSWORDERROR:    424,
	TOKENINLAILD:        425,
}

func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:             "成功",
		ApiCode.FAILED:              "失败",
		ApiCode.NOAUTH:              "身份认证失败",
		ApiCode.AUTHFORMATERROR:     "身份认证失败",
		ApiCode.MissingLoginParams:  "缺少登录必须参数",
		ApiCode.CodeExpired:         "验证码过期",
		ApiCode.CaptStatus:          "验证码错误",
		ApiCode.PasswordError:       "账号或密码错误",
		ApiCode.StatusEnabled:       "该账号已被禁用",
		ApiCode.POSTEXISTS:          "岗位已存在",
		ApiCode.DEPTEXISTS:          "部门已存在",
		ApiCode.DEPTDELETEFAILED:    "部门删除失败",
		ApiCode.MENUEXISTS:          "菜单已存在",
		ApiCode.DELETEMENUFAILED:    "删除菜单失败",
		ApiCode.ADDROLEFAILED:       "添加角色失败",
		ApiCode.MissParameter:       "参数错误",
		ApiCode.CREATEUSERFAILED:    "创建用户失败",
		ApiCode.UPLOADFILEFAILED:    "上传文件失败",
		ApiCode.UPDATEUSERFAILED:    "更新用户信息失败",
		ApiCode.RESETPASSWORDFAILED: "重置密码失败",
		ApiCode.OLDPASSWORDERROR:    "旧密码错误",
		ApiCode.TWICEPASSWORDERROR:  "两次密码不一致",
		ApiCode.TOKENINLAILD:        "token验证失败",
	}
}

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
