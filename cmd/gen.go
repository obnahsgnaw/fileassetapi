package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

var frameworkPkg = "github.com/obnahsgnaw/frameworkapi"
var pkgPrefix = "github.com/efly/"

func main() {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	_, pkgName := filepath.Split(currentDir)

	pkg := pkgPrefix + pkgName
	project := strings.TrimSuffix(pkgName, "api")
	if project == "" {
		println("project invalid")
		os.Exit(2)
	}
	if err = renameGoMod(currentDir, pkg); err != nil {
		println(err.Error())
		os.Exit(3)
	}
	if err = renameDoc(currentDir, pkg); err != nil {
		println(err.Error())
		os.Exit(3)
	}
	if err = renameApi(currentDir, pkg, project, "backend", "backend"); err != nil {
		println(err.Error())
		os.Exit(4)
	}
	if err = renameApi(currentDir, pkg, project, "frontend", "frontend"); err != nil {
		println(err.Error())
		os.Exit(5)
	}
	if err = renameApi(currentDir, pkg, project, "tcp", "tcp"); err != nil {
		println(err.Error())
		os.Exit(6)
	}
	if err = renameApi(currentDir, pkg, project, "websocket", "wss"); err != nil {
		println(err.Error())
		os.Exit(7)
	}
	err = os.RemoveAll(filepath.Join(currentDir, "gen"))
	if err != nil {
		println(err.Error())
		os.Exit(8)
	}
}
func renameGoMod(currentDir, pkg string) error {
	file := filepath.Join(currentDir, "go.mod")
	return renameFilePackage(pkg, file)
}
func renameDoc(currentDir, pkg string) error {
	file := filepath.Join(currentDir, "doc", "doc.go")
	return renameFilePackage(pkg, file)
}
func renameFilePackage(pkg, file string) error {
	f, err := os.Stat(file)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	data = bytes.Replace(data, []byte(frameworkPkg), []byte(pkg), -1)
	err = os.WriteFile(file, data, f.Mode())
	if err != nil {
		return err
	}
	return nil
}
func renameApi(currentDir, projectPackage, projectName, dirName, kind string) error {
	// init buf.gen.yaml
	bufGen := filepath.Join(currentDir, dirName, "buf.gen.yaml")
	if err := renameFilePackage(projectPackage, bufGen); err != nil {
		return err
	}
	// rename api currentDir
	protoPkgFrom := "framework_" + kind + "_api"
	protoPkgTo := projectName + "_" + kind + "_api"
	apiDir := filepath.Join(currentDir, dirName, "apis", protoPkgFrom)
	toDir := filepath.Join(currentDir, dirName, "apis", protoPkgTo)
	if err := os.Rename(apiDir, toDir); err != nil {
		return err
	}
	// init proto
	return initProto(protoPkgFrom, protoPkgTo, toDir)
}
func initProto(protoPkgFrom, protoPkgTo, dir string) error {
	sub, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, de := range sub {
		file := filepath.Join(dir, de.Name())
		if de.IsDir() {
			if err = initProto(protoPkgFrom, protoPkgTo, file); err != nil {
				return err
			}
		} else {
			if err = renameProtoPackage(file, protoPkgFrom, protoPkgTo); err != nil {
				return err
			}
		}
	}
	return nil
}
func renameProtoPackage(file, from, to string) error {
	f, err := os.Stat(file)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	data = bytes.Replace(data, []byte(from), []byte(to), -1)
	err = os.WriteFile(file, data, f.Mode())
	if err != nil {
		return err
	}
	return nil
}
