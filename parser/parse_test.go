package parser

import (
	"fmt"
	"go/ast"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type copyFormatter struct {
}

func (*copyFormatter) Match(st *ast.StructType) (fields []*ast.Field) {
	return st.Fields.List
}

func (*copyFormatter) Format(field *ast.Field) {
	if field.Tag != nil {
		s, _ := strconv.Unquote(field.Tag.Value)
		tag := reflect.StructTag(s)
		gtag := tag.Get("gorm")

		if gtag != "" {
			var tv strings.Builder
			tv.WriteRune('`')
			tv.WriteString(s)
			tv.WriteRune(' ')
			tv.WriteString(fmt.Sprintf(`custom:"%s"`, gtag))
			tv.WriteRune('`')
			field.Tag.Value = tv.String()
		}
	}
}

func TestFormatCopyTags(t *testing.T) {
	formatter := &copyFormatter{}
	diffData, err := FormatTags("../testdatas/models.go.txt", formatter)
	if err != nil {
		t.Error(err)
	}
	if len(diffData) == 0 {
		t.Errorf("diff data is empty")
	}
	fmt.Print(string(diffData))
}

type nochangeFormatter struct {
}

func (*nochangeFormatter) Match(st *ast.StructType) (fields []*ast.Field) {
	return st.Fields.List
}
func (*nochangeFormatter) Format(field *ast.Field) {}

func TestFormatNochangeTags(t *testing.T) {
	formatter := &nochangeFormatter{}
	diffData, err := FormatTags("../testdatas/models.go.txt", formatter)
	if err != nil {
		t.Error(err)
	}
	if len(diffData) != 0 {
		t.Errorf("diff data not empty \n %s", string(diffData))
	}
}
