package graphqlfields

import (
	"github.com/graphql-go/graphql/language/ast"
)

// GetSelectedFields getting list of passed fields in graphql
// https://github.com/graphql-go/graphql/issues/157
func GetSelectedFields(selectionPath []string, info graphql.ResolveInfo) []string {
	fields := info.FieldASTs
	var collect, collectFragment []string

	if len(info.Fragments) != 0 {
		for _, fragment := range info.Fragments {
			for _, selection := range fragment.GetSelectionSet().Selections {
				collectFragment = append(collectFragment, selection.(*ast.Field).Name.Value)
			}
		}
	}
	for _, propName := range selectionPath {
		found := false
		collect = collectFragment
		for _, field := range fields {
			if field.Name.Value == propName {
				selections := field.SelectionSet.Selections
				fields = make([]*ast.Field, 0)
				for _, selection := range selections {
					switch value := selection.(type) {
					case *ast.Field:
						fields = append(fields, value)
						name := value.Name.Value
						if name == "__typename" || name == "id" {
							continue
						}
						collect = append(collect, value.Name.Value)
					}
				}
				found = true
				break
			}
		}
		if !found {
			return collect
		}
	}
	// добавляем id, так как он всегда требуется для подзапров
	collect = append(collect, "id")
	return collect
}
