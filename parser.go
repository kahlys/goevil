package goevil

import (
	"io/ioutil"

	"golang.org/x/mod/modfile"
)

func getDependencies(file string) ([]string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return []string{}, err
	}

	f, err := modfile.Parse("", data, nil)
	if err != nil {
		return []string{}, err
	}

	res := []string{}
	for _, dep := range f.Require {
		if len(dep.Mod.Path) > 0 {
			res = append(res, dep.Mod.Path)
		}
	}
	return res, nil
}
