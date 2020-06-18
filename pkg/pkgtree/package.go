package pkgtree

import (
	"strings"

	"github.com/golang/dep/gps/pkgtree"
)

func ListPackages(fileRoot, importRoot string) map[string][]string {
	tree, err := pkgtree.ListPackages(fileRoot, importRoot)
	if err != nil {
		panic(err)
	}

	packageImports := make(map[string][]string)
	for currentPackage, packageOrErr := range tree.Packages {
		imports := packageOrErr.P.Imports
		if len(imports) == 0 {
			continue
		}

		temp := make([]string, 0)
		for _, importPackage := range imports {
			if strings.HasPrefix(importPackage, importRoot) {
				importPackage = strings.TrimPrefix(importPackage, importRoot)
				temp = append(temp, importPackage)
			}
		}

		packageImports[strings.TrimPrefix(currentPackage, importRoot)] = temp
	}

	return packageImports
}
