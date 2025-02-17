package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Hello world")
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{
		// {ID: 1, Title: "Buy milk", Completed: false},
		// {ID: 2, Title: "Buy eggs", Completed: false},
		// {ID: 3, Title: "Buy bread", Completed: false},
	}

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{
			Title: c.FormValue("title"),
			// Completed: c.FormValue("completed") == "true",
		}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Title == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Title is required",
			})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true

				return c.Status(fiber.StatusOK).JSON(todos[i])
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Todo not found",
		})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)

				return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo not found"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
