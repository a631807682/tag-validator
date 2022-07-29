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

type TagType int8

const (
	// table column definition
	Table TagType = iota + 1
	// column permission
	Permission
	// gorm internal implementation
	Internal
	// schema relation
	Relation
)

// https://gorm.io/docs/models.html
var tagNamesMap = map[string]TagType{
	"column":                 Table,
	"type":                   Table,
	"size":                   Table,
	"primaryKey":             Table,
	"primary_key":            Table, // same with primaryKey
	"unique":                 Table,
	"default":                Table,
	"precision":              Table,
	"scale":                  Table,
	"not null":               Table,
	"autoIncrement":          Table,
	"autoIncrementIncrement": Table,
	"index":                  Table,
	"uniqueIndex":            Table,
	"comment":                Table,

	"serializer":     Internal,
	"embedded":       Internal,
	"embeddedPrefix": Internal,
	"autoCreateTime": Internal,
	"autoUpdateTime": Internal,
	"check":          Internal,

	"<-": Permission,
	"->": Permission,
	"-":  Permission,

	"foreignKey":       Relation,
	"references":       Relation,
	"polymorphic":      Relation,
	"many2many":        Relation,
	"belongsto":        Relation,
	"polymorphicValue": Relation,
	"joinForeignKey":   Relation,
	"joinReferences":   Relation,
	"constraint":       Relation,
}
