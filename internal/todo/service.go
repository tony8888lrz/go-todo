package todo

// Storage 接口（数据存储层的最小抽象）
type Storage interface {
	Create(t Todo) (Todo, error)
	List() ([]Todo, error)
	Delete(id int) error
}

// Service 接口（业务层需要实现的操作）
type Service interface {
	AddTodo(title string) (Todo, error)
	ListTodo() ([]Todo, error)
	DeleteTodo(id int) error
}

// TodoService 的具体实现
type TodoService struct {
	store Storage // 不直接 import storage 包，而是用接口来“依赖倒置”
}

// NewTodoService 构造函数
func NewTodoService(s Storage) *TodoService {
	return &TodoService{store: s}
}

func (ts *TodoService) AddTodo(title string) (Todo, error) {
	t := Todo{
		Title: title,
		// ... 其他初始化
	}
	return ts.store.Create(t)
}

func (ts *TodoService) ListTodo() ([]Todo, error) {
	return ts.store.List()
}

func (ts *TodoService) DeleteTodo(id int) error {
	return ts.store.Delete(id)
}
