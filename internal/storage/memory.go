package storage

import (
	"errors"
	"sync"

	"github.com/tony8888lrz/go-todo/internal/todo"
)

// MemoryStorage 实现了 todo 包里的 Storage 接口
type MemoryStorage struct {
	data   []todo.Todo
	mutex  sync.Mutex
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:   make([]todo.Todo, 0),
		nextID: 1,
	}
}

// 实现 Storage.Create()
func (ms *MemoryStorage) Create(t todo.Todo) (todo.Todo, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	t.ID = ms.nextID
	ms.nextID++
	ms.data = append(ms.data, t)
	return t, nil
}

// 实现 Storage.List()
func (ms *MemoryStorage) List() ([]todo.Todo, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	copyData := make([]todo.Todo, len(ms.data))
	copy(copyData, ms.data)
	return copyData, nil
}

// 实现 Storage.Delete()
func (ms *MemoryStorage) Delete(id int) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	index := -1
	for i, t := range ms.data {
		if t.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("todo not found")
	}
	ms.data = append(ms.data[:index], ms.data[index+1:]...)
	return nil
}
