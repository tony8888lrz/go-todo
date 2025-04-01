package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/tony8888lrz/go-todo/internal/storage"
	"github.com/tony8888lrz/go-todo/internal/todo"
)

func main() {
	// 定义一些命令行子命令：add/list/delete 增删改查
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// 给add命令加一个title参数
	addTitle := addCmd.String("title", "", "Title of the todo")

	// 创建存储和服务
	memStore := storage.NewMemoryStorage()
	todoService := todo.NewTodoService(memStore)

	// 解析命令
	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list' or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		// 执行service逻辑
		if *addTitle == "" {
			fmt.Println("Title is required. Use -title <your title>")
			os.Exit(1)
		}
		newTodo, err := todoService.AddTodo(*addTitle)
		if err != nil {
			fmt.Println("Error adding todo:", err)
			os.Exit(1)
		}
		fmt.Printf("Added Todo: ID=%d, Title=%s\n", newTodo.ID, newTodo.Title)
	case "list":
		listCmd.Parse(os.Args[2:])
		todos, err := todoService.ListTodo()
		if err != nil {
			fmt.Println("Error listing todos:", err)
			os.Exit(1)
		}
		if len(todos) == 0 {
			fmt.Println("No todos found.")
			return
		}
		fmt.Println("All Todos:")
		for _, t := range todos {
			status := "Not Done"
			if t.Done {
				status = "Done"
			}
			fmt.Printf("ID=%d, Title=%s, Status=%s\n", t.ID, t.Title, status)
		}
	case "delete":
		delCmd.Parse(os.Args[2:])
		if len(delCmd.Args()) < 1 {
			fmt.Println("Please provide the ID to delete. e.g. delete 1")
			os.Exit(1)
		}
		idStr := delCmd.Args()[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("Invalid ID: %s\n", idStr)
			os.Exit(1)
		}
		err = todoService.DeleteTodo(id)
		if err != nil {
			fmt.Println("Error deleting todo:", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted Todo with ID %d\n", id)
	default:
		fmt.Println("expected 'add', 'list' or 'delete' subcommands")
		os.Exit(1)
	}
}
