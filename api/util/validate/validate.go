package validate

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslate "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		uni = ut.New(enT, zhT, enT)
		trans, _ = uni.GetTranslator("zh") // 翻译器
		_ = zhTranslate.RegisterDefaultTranslations(v, trans)
	}
}

func Validate(f func(obj interface{}) error, i interface{}) error {
	validatorErr := f(i)
	if validatorErr == nil {
		return nil
	}
	if errArr, ok := validatorErr.(validator.ValidationErrors); ok {
		for _, value := range errArr {
			errStr := value.Translate(trans)
			fieldName := value.StructField()
			tagName := value.Tag()
			ttt, _ := reflect.TypeOf(i).Elem().FieldByName(fieldName)
			label := ttt.Tag.Get("label")
			tip := ttt.Tag.Get("tip")
			if label != "" {
				errStr = strings.Replace(errStr, fieldName, label, -1)
			}
			if tip != "" {
				strSlice := strings.Split(tip, ",")
				for _, value := range strSlice {
					tips := strings.Split(value, "=")
					if len(tips) == 2 {
						if tips[0] == tagName && tips[1] != "" {
							errStr = tip
						}
					}
				}

			}
			return errors.New(errStr)
		}
	}
	return validatorErr
}
