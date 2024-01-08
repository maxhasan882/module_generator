package file

import (
	"fmt"
	"os"
	"path/filepath"
)

const basePath = "code_gen/"

// ExtractFileNamesFromDirectory returns a list of file names in the specified directory.
func ExtractFileNamesFromDirectory(directoryPath string) ([]string, error) {
	var fileNames []string

	// Open the directory
	dir, err := os.Open(directoryPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// Read the list of files in the directory
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	// Iterate through the file list and add file names to the result
	for _, fileInfo := range fileInfos {
		// Check if it's a regular file
		if fileInfo.Mode().IsRegular() {
			fileNames = append(fileNames, fileInfo.Name())
		}
	}

	return fileNames, nil
}

// CreateOrUpdateModule creates or updates a module.
func CreateOrUpdateModule(path, fileName, content string) error {
	// Create the directory if it doesn't exist
	if err := createDirectory(getPath(path)); err != nil {
		return fmt.Errorf("failed to create directory for module: %w", err)
	}

	// Open or create the file
	file, err := os.OpenFile(filepath.Join([]string{path, fileName}...), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open or create file for module: %w", err)
	}
	defer file.Close()

	// Write the content
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write content to module file: %w", err)
	}

	return nil
}

// CreateDirectory creates a directory if it doesn't exist.
func createDirectory(dir string) error {
	// Use Stat to check if the directory already exists
	stat, err := os.Stat(dir)
	if err == nil && stat.IsDir() {
		// The directory already exists
		return nil
	}

	// The directory doesn't exist, so create it
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	return nil
}

func CreateOrUpdateSchemaModule(filename, content, packageName string) error {
	modulePath := "infrastructure/repository/internal/" + packageName + "/" + filename

	if _, err := os.Stat(getPath(modulePath)); os.IsNotExist(err) {
		// If the file does not exist, create the directory and the file
		if err := os.MkdirAll(getPath("infrastructure/repository/internal/"+packageName), 0755); err != nil {
			return err
		}
	}

	file, err := os.OpenFile(getPath(modulePath), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file information to check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// If the file was just created or is empty, write the package declaration
	if fileInfo.Size() == 0 {
		_, err := file.WriteString(fmt.Sprintf("package %s", packageName))
		if err != nil {
			return err
		}
	}

	// Write the content
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func RemoveDirectory(directoryPath string) error {
	err := os.RemoveAll(getPath(directoryPath))
	if err != nil {
		return err
	}
	return nil
}

func getPath(path string) string {
	return path
}
