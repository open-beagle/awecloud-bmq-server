package util

import (
	"errors"
	"image/color"

	"github.com/mojocn/base64Captcha"
	"github.com/open-beagle/awecloud-bmq-server/pkg/model"
)

func GetCaptcha(width, height int) (captcha model.Captcha, err error) {
	param := configJsonBody{
		CaptchaType: "string",
		DriverString: &base64Captcha.DriverString{
			Height:          height,
			Width:           width,
			NoiseCount:      0,
			ShowLineOptions: 2,
			Length:          4,
			// Source:          "abcdefghigklmnopqrstuvwxyz124567890",
			Source:  "124567890",
			BgColor: &color.RGBA{255, 255, 255, 1},
			//Fonts: []string{"3Dumb.ttf", "ApothecaryFont.ttf", "Comismsh.ttf", "DENNEthree-dee.ttf", "DeborahFancyDress.ttf",
			//	"Flim-Flam.ttf", "RitaSmith.ttf", "actionj.ttf", "chromohv.ttf"},
			Fonts: []string{"RitaSmith.ttf"},
		},
	}
	captcha.Id, captcha.Captcha, err = generateCaptcha(param)
	return
}

func VerifyCaptcha(id, value string) (err error) {
	param := configJsonBody{}
	param.Id = id
	param.VerifyValue = value
	right := verify(param)
	if !right {
		err = errors.New("验证码不正确")
	}
	return
}

var store = base64Captcha.DefaultMemStore

func generateCaptcha(param configJsonBody) (id, b64s string, err error) {
	var driver base64Captcha.Driver
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = c.Generate()
	return
}

type configJsonBody struct {
	Id            string                       `json:"id"`
	CaptchaType   string                       `json:"captcha_type"`
	VerifyValue   string                       `json:"verify_value"`
	DriverAudio   *base64Captcha.DriverAudio   `json:"driver_audio"`
	DriverString  *base64Captcha.DriverString  `json:"driver_string"`
	DriverChinese *base64Captcha.DriverChinese `json:"driver_chinese"`
	DriverMath    *base64Captcha.DriverMath    `json:"driver_math"`
	DriverDigit   *base64Captcha.DriverDigit   `json:"driver_digit"`
}

func verify(body configJsonBody) bool {
	return store.Verify(body.Id, body.VerifyValue, false)

}
