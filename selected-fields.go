package graphqlfields

import (
	"github.com/graphql-go/graphql/language/ast"
)

// GetSelectedFields getting list of passed fields in graphql
// https://github.com/graphql-go/graphql/issues/157
func GetSelectedFields(selectionPath []string, astfields []*ast.Field) []string {
	var collect []string
	var fields = make([]ast.Selection, 0)
	for _, propName := range selectionPath {
		found := false
		for _, field := range astfields {
			if field.Name.Value == propName {
				selections := field.SelectionSet.Selections
				for _, selection := range selections {
					fields = append(fields, selection)
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
		var name string
		switch value := field.(type) {
		case *ast.Field:
			name = value.Name.Value
		case *ast.FragmentSpread:
			name = value.Name.Value
		}
		// exclude fields (id will be added below, forced)
		if name == "id" || name == "__typename" {
			continue
		}
		collect = append(collect, name)
	}
	// add id, as it is always required for subqueries
	collect = append(collect, "id")
	return collect
}
