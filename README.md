# go-todo

Author: Runze Liao, Instructed by Yi Zhao & Yizhe Tang

## 目录结构

```shell
go-todo/
├── cmd/
│   └── main.go           # 项目入口，解析命令行参数并调用业务逻辑
├── internal/
│   ├── todo/
│   │   ├── model.go      # Todo 数据模型
│   │   └── service.go    # TodoService: 业务逻辑定义（增删查等）
│   └── storage/
│       ├── file.go       # FileStorage: 基于JSON文件的存储实现
│       └── memory.go     # MemoryStorage: 使用内存存储的示例（不会持久化）
├── go.mod                # Go Modules 配置文件
├── go.sum                # Go Modules 依赖锁定文件
└── README.md             # 项目说明文档
```

- **`cmd/main.go`**：  
  - 解析命令行子命令（`add`、`list`、`delete`）并执行对应操作。  
- **`internal/todo`**：  
  - `model.go` 定义了 `Todo` 结构体的数据模型。  
  - `service.go` 提供了对 Todo 的业务操作（添加、查询、删除等），使用**接口**解耦存储层。  
- **`internal/storage/file.go`**：  
  - 使用 JSON 文件的方式实现了 `todo.Storage` 接口，数据会持久化到 `todos.json`（或其他指定文件）。  
- **`internal/storage/memory.go`**：  
  - 使用内存切片保存数据，仅用于简单演示；进程退出后数据会丢失。

## 快速开始

1. **克隆或下载本仓库**

   ```bash
   git clone https://github.com/YourUsername/go-todo.git
   cd go-todo
   ```

2. **初始化或更新依赖**（可选）

   ```bash
   go mod tidy
   ```

3. **编译或直接运行**

   - **直接运行：**

     ```bash
     go run cmd/main.go <subcommand> [options]
     ```
   - **编译为可执行文件：**

     ```bash
     go build -o todo-tool cmd/main.go
     ./todo-tool <subcommand> [options]
     ```

## 使用示例

下面以直接运行的方式演示。

1. **添加待办事项**  
   ```bash
   go run cmd/main.go add -title "Buy milk"
   ```
   输出示例：
   ```
   Added Todo: ID=1, Title=Buy milk
   ```

2. **列出所有待办事项**  
   ```bash
   go run cmd/main.go list
   ```
   输出示例：
   ```
   All Todos:
   ID=1, Title=Buy milk, Status=Not Done
   ```

3. **删除待办事项**  
   ```bash
   go run cmd/main.go delete 1
   ```
   输出示例：
   ```
   Deleted Todo with ID 1
   ```

4. **再次列出**  
   ```
   go run cmd/main.go list
   ```
   若已删除所有记录，则输出：
   ```
   No todos found.
   ```

## 配置与文件存储

- 在 `main.go` 中，示例默认使用 `FileStorage` 将数据存到 `todos.json`。  
- 若 `todos.json` 文件不存在，程序首次运行会自动创建；若已存在，则会加载其中原有的待办数据。  
- 你可以在 `NewFileStorage("todos.json")` 中替换其他路径或文件名。

## 未来扩展 Todo

1. **Web 服务化**  
   - 使用 [Gin](https://github.com/gin-gonic/gin) 或其他 Web 框架，将命令行改造成一个 RESTful API 服务，实现增删改查 API。  
   - 在启动时运行 `FileStorage` 或 `MemoryStorage`，在整个服务器生命周期内保留数据。

2. **数据库支持**  
   - 将存储层替换为真正的数据库（SQLite、MySQL、PostgreSQL 等），只需实现同样的 `Storage` 接口即可。  
   - 能够在多台服务器或多人协作场景下共享数据。

3. **标记完成 / 更新任务**  
   - 增加 `UpdateTodo` 或 `MarkDone` 等操作，为待办事项添加“已完成”的功能，并在 `list` 时显示进度。

4. **更多命令行子命令**  
   - 加上 `edit` 命令来更新标题或截止日期等字段；  
   - 加入 `--done` 标记，更灵活地显示完成与未完成任务。

5. **加上单元测试**  
   - 为 `todo.Service` 和存储层写一些测试，使用 `go test` 保证质量。  
   - 使用 `testify` 等库来做断言，或直接用 Go 原生测试方式。

6. **并发优化**  
   - 如果将来需要同时处理大量请求，可以考虑锁优化或更换成数据库以获得高并发。

## 许可

本项目示例可以自由使用或改造，无特殊许可证限制，仅用于学习Go语言。