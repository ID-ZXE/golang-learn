package schema

import (
	"go/ast"
	"orm/dialect"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

type ITableName interface {
	TableName() string
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	var tableName string
	// 是否自行定义了表名
	table, ok := dest.(ITableName)
	if !ok {
		tableName = modelType.Name()
	} else {
		tableName = table.TableName()
	}
	schema := &Schema{
		Model:    dest,
		Name:     tableName,
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		structField := modelType.Field(i)
		if !structField.Anonymous && ast.IsExported(structField.Name) {
			field := &Field{
				Name: structField.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(structField.Type))),
			}
			if v, ok := structField.Tag.Lookup("orm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, structField.Name)
			schema.fieldMap[structField.Name] = field
		}
	}
	return schema
}
