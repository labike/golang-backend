package result

type Codes struct {
	SUCCESS         uint
	FAILED          uint
	Message         map[uint]string
	NOAUTH          uint
	AUTHFORMATERROR uint
}

var ApiCode = &Codes{
	SUCCESS:         200,
	FAILED:          501,
	NOAUTH:          403,
	AUTHFORMATERROR: 405,
}

func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:         "Success",
		ApiCode.FAILED:          "Failed",
		ApiCode.NOAUTH:          "身份认证失败",
		ApiCode.AUTHFORMATERROR: "身份认证失败",
	}
}

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
