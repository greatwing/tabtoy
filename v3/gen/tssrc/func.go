package tssrc

import (
	"github.com/davyxu/tabtoy/v3/model"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

// 将定义用的类型，转换为不同语言对应的复合类型

func init() {
	UsefulFunc["TSType"] = func(tf *model.TypeDefine) string {

		convertedType := model.LanguagePrimitive(tf.FieldType, "typescript")

		if tf.IsArray() {
			return convertedType + "[]"
		}

		return convertedType
	}
}

