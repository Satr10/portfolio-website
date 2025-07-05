package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type commandInput struct {
	Command string `json:"command"`
}

var allowedCommands = map[string]bool{
	"help":     true,
	"about":    true,
	"projects": true,
	"contact":  true,
	"clear":    true,
	"welcome":  true,
}

func CommandHtmx(c *fiber.Ctx) error {
	command := commandInput{}
	err := c.BodyParser(&command)
	if err != nil {
		fmt.Println("Error parsing command input:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid command input",
		})
	}
	cmdStr := strings.ToLower(strings.TrimSpace(command.Command))
	if _, ok := allowedCommands[cmdStr]; !ok {
		return c.Render("partials/commands/error", fiber.Map{
			"command": cmdStr,
		})
	}

	return c.Render("partials/commands/"+cmdStr, fiber.Map{
		"command": cmdStr,
	})

}
