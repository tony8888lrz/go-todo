package todo

import "time"

// 首字母大写表明此结构在包外可见（exported）。
type Todo struct {
	ID          int       // 待办事项ID
	Title       string    // 标题
	Done        bool      // 是否已完成
	CreatedTime time.Time // 创建时间
}
