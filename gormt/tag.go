package gormt

import (
	"go/ast"
	"reflect"
	"strconv"
)

type GormTagFormatter struct {
}

func (*GormTagFormatter) Match(st *ast.StructType) (fields []*ast.Field) {
	return st.Fields.List
}

func (*GormTagFormatter) Format(field *ast.Field) {
	if field.Tag != nil {
		s, _ := strconv.Unquote(field.Tag.Value)
		tag := reflect.StructTag(s)
		gtag := tag.Get("gorm")

		if gtag != "" {
			field.Tag.Value = format(gtag)
		}
	}
}

func format(gtag string) string {

	return gtag
}
