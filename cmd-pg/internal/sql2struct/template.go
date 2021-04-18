package sql2struct

import (
	"fmt"
	"html/template"
	"os"
	"cmd-pg/internal/cmdutils"
)

const structTp1 = `type {{ .TabelName | ToCamelCase}} struct {
{{ range .Columns}} {{ $length :=len .Comment}} {{ if gt $length 0 }}//
{{.Comment}} {{else}}// {{.Name}} {{ end }}
    {{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}
    {{.Type}} {{.Tag}}{{ else }}{{.Name}} {{ end }}
{{end}}}

func (mode {{.TableName | ToCamelCase}}) TableName() string {
   return "{{.TableName}}""
}`

type StructTemplate struct {
	structTp1 string
}

type StructColumn struct {
	Name       string
	Type       string
	Tag        string
	Comment    string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
   return &StructTemplate{structTp1: structTp1}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn{
	tp1Columns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tp1Columns = append(tp1Columns, &StructColumn{
			Name: column.ColumnName,
			Type: DBTypeToStructType[column.DataType],
			Tag:  tag,
			Comment: column.ColumnComment,
		})
	}
	return tp1Columns
}
var word string
func (t *StructTemplate) Generate(tableName string,tp1Columns []*StructColumn) error  {
	tp1 := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.ModeUnderscoreToUpperCamelCase,
	}).Parse(t.structTp1))

	tp1DB := StructTemplateDB{
		TableName: tableName,
		Columns: tp1Columns,
	}
	err := tp1.Execute(os.Stdout, tp1DB)
	if err != nil {
		return err
	}
	return err
}