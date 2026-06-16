/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/todo_tui/internal/domain"
	"example.com/todo_tui/internal/service"
	"example.com/todo_tui/internal/sqlite"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long:  "This command add new task.ex)add wash_dishes ",
	Run: func(cmd *cobra.Command, args []string) {

		db, err := sqlite.Connect_db()
		if err != nil {
			log.Fatal(err)
		}

		var newTask domain.Task

		newTask.Title = args[0]

		newTask.Description = description
		if dueDate != "" {
			due, err := time.Parse("2006-01-02", dueDate)
			if err != nil {
				log.Fatal(err)
			}
			newTask.DueDate = &due
		}

		var persedPriority domain.Priority
		persedPriority = domain.Priority(priority)
		newTask.Priority = persedPriority

		repo := sqlite.NewTaskRepository(db)
		s := service.NewTaskService(repo)
		ctx := context.Background()
		fmt.Println(newTask)
		createdTask, err := s.CreateTask(ctx, newTask)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(createdTask)
	},
}

var description string
var dueDate string
var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&description, "desc", "", "task description")

	addCmd.Flags().StringVar(&dueDate, "due", "", "task expire")

	addCmd.Flags().IntVar(&priority, "pri", 0, "task priority")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
