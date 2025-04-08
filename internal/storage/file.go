package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/tony8888lrz/go-todo/internal/todo"
)

// FileStorage is a struct that implements the Storage interface for file-based storage.
type FileStorage struct {
	filePath string
	mutex    sync.Mutex
	data     []todo.Todo
	nextID   int
}

func NewFileStorage(filePath string) (*FileStorage, error) {
	// Initialize the FileStorage with the given filename.
	// Load existing data from the file if it exists.
	// If the file does not exist, create a new one.
	// Return a pointer to the FileStorage instance.
	fs := &FileStorage{
		filePath: filePath,
		data:     []todo.Todo{},
		nextID:   0,
	}
	if err := fs.loadFromFile(); err != nil {
		// 如果文件不存在，也不算严重错误，可忽略
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}
	return fs, nil
}

// Create 实现了 Storage.Create()；会先给Todo分配一个自增ID，存进内存，并写回文件
func (fs *FileStorage) Create(t todo.Todo) (todo.Todo, error) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	fs.nextID++
	t.ID = fs.nextID
	fs.data = append(fs.data, t)

	// 写回文件
	if err := fs.saveToFile(); err != nil {
		return t, err
	}
	return t, nil
}

// List 返回所有 Todo
func (fs *FileStorage) List() ([]todo.Todo, error) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	// 返回一个副本或者直接返回
	result := make([]todo.Todo, len(fs.data))
	copy(result, fs.data)
	return result, nil
}

// Delete 删除指定 ID 的待办事项
func (fs *FileStorage) Delete(id int) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	index := -1
	for i, t := range fs.data {
		if t.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("todo not found")

	}
	fs.data = append(fs.data[:index], fs.data[index+1:]...)

	return fs.saveToFile()
}

// loadFromFile 从 filePath 加载 JSON 数据到 fs.data
func (fs *FileStorage) loadFromFile() error {
	f, err := os.Open(fs.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	var loaded []todo.Todo
	if err := decoder.Decode(&loaded); err != nil {
		return err
	}

	fs.data = loaded
	maxID := 0
	for _, t := range loaded {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	fs.nextID = maxID

	return nil
}

// saveToFile 将 fs.data 中的数据写入到 filePath，正确的io
func (fs *FileStorage) saveToFile() error {
	f, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ") // 美化排版(可选)
	return encoder.Encode(fs.data)
}
