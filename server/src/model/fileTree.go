package model

import (
	"io/ioutil"
)

type FileNode struct {
	Name  string      `json:"name"`
	Size  int64       `json:"size"`
	IsDir bool        `json:"isDir"`
	Files []*FileNode `json:"files"`
}

func newFileNode(name string, size int64, isDir bool) *FileNode {
	return &FileNode{
		Name:  name,
		Size:  size,
		IsDir: isDir,
		Files: []*FileNode{},
	}
}

func BuildFileTree() *FileNode {
	root := newFileNode("root", 0, true)

	travelFiles(root, storageRootPath)

	return root
}

func travelFiles(cur *FileNode, path string) {
	files, _ := ioutil.ReadDir(path)

	for _, file := range files {
		fileNode := newFileNode(file.Name(), file.Size(), file.IsDir())
		if file.IsDir() {
			travelFiles(fileNode, path+"/"+file.Name())
		}
		cur.Files = append(cur.Files, fileNode)
	}
}
