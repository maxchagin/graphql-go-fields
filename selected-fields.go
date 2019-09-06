package graphqlfields

import (
	"github.com/graphql-go/graphql/language/ast"
)

// GetSelectedFields getting list of passed fields in graphql
// https://github.com/graphql-go/graphql/issues/157
func GetSelectedFields(selectionPath []string, fields []*ast.Field) []string {
	var collect []string
	for _, propName := range selectionPath {
		found := false
		for _, field := range fields {
			if field.Name.Value == propName {
				selections := field.SelectionSet.Selections
				fields = make([]*ast.Field, 0)
				for _, selection := range selections {
					fields = append(fields, selection.(*ast.Field))
				}
				found = true
				break
			}
		}
		if !found {
			return collect
		}
	}

	for _, field := range fields {
		name := field.Name.Value
		// исключаем поля (id будет добавлен ниже, принудительно)
		if name == "id" || name == "__typename" {
			continue
		}
		collect = append(collect, field.Name.Value)
	}
	// добавляем id, так как он всегда требуется для подзапров
	collect = append(collect, "id")

	return collect
}
