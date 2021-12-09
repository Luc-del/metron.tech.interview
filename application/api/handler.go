package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"interview/application/pkg/models/discount"
	"interview/application/pkg/services/cashier"
	"net/http"
)

func status(version string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		status := http.StatusOK
		return c.Status(status).JSON(map[string]interface{}{
			"version":   version,
			"discounts": discount.GetDiscounts(),
		})
	}
}

func cost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("calling cost")
		var items struct {
			Ids []int `json:"basketBookIds"`
		}
		if err := c.BodyParser(&items); err != nil {
			return c.Status(http.StatusOK).JSON(map[string]interface{}{
				"error": err,
			})
		}

		fmt.Println("items:", items)
		cost := cashier.ComputeCost(discount.GetDiscounts(), items.Ids...)

		fmt.Printf("here's your receipt: %+v", cost)
		return c.Status(http.StatusOK).JSON(cost)
	}
}

func setDiscount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("calling discount settings")
		var params discount.Parameters
		if err := c.BodyParser(&params); err != nil {
			fmt.Println("error:", err)
			return c.Status(http.StatusOK).JSON(map[string]interface{}{
				"error": err,
			})
		}

		if err := params.Validate(); err != nil {
			fmt.Println("error:", err)
			return c.Status(http.StatusOK).JSON(map[string]interface{}{
				"error": err,
			})
		}
		fmt.Println("settings:", params)
		discount.Set(params)

		fmt.Println("new discounts:", discount.GetDiscounts())
		return c.Status(http.StatusOK).JSON(discount.GetDiscounts())
	}
}
