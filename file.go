package gofs

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

type File struct {
	PathHandler
}

func NewFile(path PathHandler) File {
	return File{path}
}

func (f *File) Size() (int64, error) {
	// get the file size
	stat, err := os.Stat(f.String())
	return stat.Size(), err
}

func (f *File) GetMetaData() {

}

func (f *File) Delete() error {
	err := os.Remove(f.String())
	return err
}

func (f *File) Copy(destPath PathHandler) error {
	// copy the file to the new path
	destDir := destPath.Dir()
	err := destDir.CreateIfNotExist()
	if err != nil {
		return err
	}
	srcFile, err := os.Open(f.String())
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(filepath.Join(destPath.String(), f.Name()))
	if err != nil {
		return fmt.Errorf("creating destination file: %w", err)
	}
	defer dstFile.Close()

	// Allocate a buffer
	buf := make([]byte, 4096)
	for {
		n, err := syscall.Read(syscall.Handle(int(srcFile.Fd())), buf)
		if err != nil {
			return fmt.Errorf("reading from source file: %w", err)
		}
		if n == 0 {
			break // EOF
		}
		_, err = syscall.Write(syscall.Handle(int(dstFile.Fd())), buf[:n])
		if err != nil {
			return fmt.Errorf("writing to destination file: %w", err)
		}
	}

	return nil
}

func (f *File) Create(overwrite bool) error {
	if f.IsFile() || f.Exists() {
		if overwrite {
			f.Delete()
		} else {
			return nil
		}
	}
	_, err := os.Create(f.String())
	return err
}

func (f *File) CreateIfNotExists() error {
	return f.Create(false)
}

func (f *File) Read() ([]byte, error) {

}

func (f *File) ReadAll() ([]byte, error) {

}

func (f *File) ReadString() (string, error) {

}

func (f *File) ReadLines() ([]string, error) {

}

func (f *File) Write(data []byte) error {

}

func (f *File) WriteString(data string) error {

}

func (f *File) WriteLine(data []string) error {

}

func (f *File) Append(data []byte) error {

}

func (f *File) AppendString(data string, newLine bool) error {

}

func (f *File) AppendLine(data string) error {

}

func (f *File) ReadStream() (any, error) {

}

func (f *File) WriteStream() (any, error) {

}
