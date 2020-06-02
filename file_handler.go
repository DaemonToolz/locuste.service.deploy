package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// ListVersions Récupère la liste des versions logicielles disponibles
func ListVersions() []string {
	files, err := ioutil.ReadDir("./repo/versions")
	if err != nil {
		log.Fatal(err)
	}

	folders := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			folders = append(folders, f.Name())
		}
	}

	return folders
}

// Unzip extraction d'un fichier archive
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)
	for _, f := range r.File {
		err := extractAndWriteFile(dest, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractAndWriteFile(dest string, f *zip.File) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer func() {
		if err := rc.Close(); err != nil {
			panic(err)
		}
	}()

	path := filepath.Join(dest, f.Name)

	if f.FileInfo().IsDir() {
		os.MkdirAll(path, f.Mode())
	} else {
		os.MkdirAll(filepath.Dir(path), f.Mode())
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer func() {
			if err := f.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

// CopyDirectory Copier un répertoire cible (récursif)
func CopyDirectory(scrDir, dest string, indicator *FileCopyInfo) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			log.Println("Impossible de charger les informations pour", sourcePath)
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath, indicator); err != nil {
				return err
			}
		default:
			indicator.CurrentFile = fileInfo.Name()
			indicator.FileIndex++
			go broadcastIndicator(*indicator)
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		if err := os.Chmod(destPath, entry.Mode()); err != nil {
			return err
		}

	}
	return nil
}

// Copy Copier un fichier cible
func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

// Exists Vérifie si un fichier / répertoire existe
func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateIfNotExists Créé un répertoire s'il n'existe pas
func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

// CountFiles Compte tous les fichiers (récursif)
func CountFiles(DirPath string, parent string) int {
	files, err := ioutil.ReadDir(DirPath)
	total := 0

	if err != nil {
		fmt.Println("Error opening file:", err)
		return total
	}

	for _, f := range files {
		if f.IsDir() == false {
			total++
		} else {
			total += CountFiles(DirPath+"/"+f.Name(), parent+"/"+f.Name())
		}
	}

	return total
}

// RemoveContents Supprime le contenu d'un répertoire
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteVersionFiles Supprime le contenu du répertoire des versions
func DeleteVersionFiles(root, version string) error {
	d, err := os.Open(root)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		if strings.Contains(name, version) {
			err = os.RemoveAll(filepath.Join(root, name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
