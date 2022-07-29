package parser

import (
	"fmt"
	"go/ast"
	"os"
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
	filename1 := "../testdatas/models.go.txt"
	filename2 := "../testdatas/copy_format_models.go.txt"

	res, err := FormatTags(filename1, formatter)
	if err != nil {
		t.Error(err)
	}

	expectData, err := os.ReadFile(filename2)
	if err != nil {
		t.Error(err)
	}

	diff := Diff(filename1, res, filename2+".orig", expectData)
	if len(diff) != 0 {
		t.Errorf("expect not equal diff data:\n%s", string(diff))
	}
}

type nochangeFormatter struct {
}

func (*nochangeFormatter) Match(st *ast.StructType) (fields []*ast.Field) {
	return st.Fields.List
}
func (*nochangeFormatter) Format(field *ast.Field) {}

func TestFormatNochangeTags(t *testing.T) {
	formatter := &nochangeFormatter{}

	filename := "../testdatas/models.go.txt"

	res, err := FormatTags(filename, formatter)
	if err != nil {
		t.Error(err)
	}

	expectData, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	diff := Diff(filename, res, filename, expectData)
	if len(diff) != 0 {
		t.Errorf("expect not equal diff data:\n%s", string(diff))
	}
}
