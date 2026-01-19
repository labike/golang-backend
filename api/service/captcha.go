package service

import (
	"go-admin/util"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

var store = util.RedisStore{}

// 生成验证码
func CaptchaMake() (id, b64s string) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      1,
		ShowLineOptions: 2 | 4,
		Length:          6,
		Source:          "123456789zxcvbnmasdfghjklqwertyuiop",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64, _, _ := captcha.Generate()
	return lid, lb64
}

func CaptVerify(id string, capt string) bool {
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
