package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"

	"go.uber.org/zap"
)

const regexTmp = "\"([^\"]*)\""

var logger *zap.Logger

// nolint
type Field struct {
	Name    string // value_name
	Type    string // value_type
	Tag     string // value_tag
	Comment string // 注释
}

// Schema代表数据库中的表
type Schema struct {
	TableName     string            // 表名
	Fields        []*Field          // 约束条件
	ModelType     string            // 表明是什么model类型
	FieldMap      map[string]*Field // 记录字段名和Field的映射关系
	StructComment string            // struct的注释
}

func (schema *Schema) GetField(name string) *Field {
	return schema.FieldMap[name]
}

// nolint
func ParseStruct(curPath string) (*Schema, error) {
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, curPath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	// 记录表名和详细信息
	schema := &Schema{
		FieldMap: make(map[string]*Field),
	}
	for _, d := range f.Decls {
		switch specDecl := d.(type) {
		case *ast.GenDecl:
			if specDecl.Tok != token.TYPE {
				continue
			}
			// 拿到struct的spec
			spec := specDecl.Specs[0]
			switch tp := spec.(*ast.TypeSpec).Type.(type) {
			case *ast.StructType:
				_ = tp.Struct

				schema.StructComment = specDecl.Doc.Text()
				switch detailedInfo := specDecl.Specs[0].(*ast.TypeSpec).Type.(type) {
				case *ast.StructType:
					_ = detailedInfo.Struct

					fields := detailedInfo.Fields
					for _, field := range fields.List {
						if field.Names != nil {
							if !ast.IsExported(field.Names[0].Name) {
								// 使用正则表达式提取出符合要求的部分
								if field.Tag != nil {
									regexVal := regexp.MustCompile(regexTmp)
									schema.TableName = strings.Split(regexVal.FindStringSubmatch(field.Tag.Value)[1], ",")[0]
								}
							} else {
								typeName, _ := baseTypeName(field.Type.(ast.Expr))
								tmpField := &Field{
									Name: field.Names[0].Name, // 字段名
									Type: typeName,
								}
								if field.Tag != nil {
									tmpField.Tag = field.Tag.Value
								}
								// 取到注释部分
								if field.Comment != nil {
									comment := make([]byte, len(field.Comment.List[0].Text))
									copy(comment, field.Comment.List[0].Text)
									if field.Comment.List[0].Text != "" && len(comment) > 2 {
										comment = comment[3:len(comment)]
									}
									tmpField.Comment = string(comment)
								}
								schema.Fields = append(schema.Fields, tmpField)
								schema.FieldMap[tmpField.Name] = tmpField
							}
						} else {
							// 判断是Model还是minModel
							name, _ := baseTypeName(field.Type.(ast.Expr))
							schema.ModelType = name
						}
					}
				}
			}

		default:
			continue
		}
	}
	return schema, nil
}

// 拿到type对应的name
func baseTypeName(x ast.Expr) (name string, imported bool) {
	switch t := x.(type) {
	case *ast.Ident:
		return t.Name, false
	case *ast.SelectorExpr:
		if _, ok := t.X.(*ast.Ident); ok {
			return t.Sel.Name, true
		}
	case *ast.ParenExpr:
		return baseTypeName(t.X)
	case *ast.StarExpr:
		return baseTypeName(t.X)
	}
	return
}
