package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/podanypepa/cursor-pagination/db"
)

type resp struct {
	Items  []db.User `json:"items"`
	Cursor cursor    `json:"cursor"`
}

type cursor struct {
	Prev string `json:"prev"`
	Next string `json:"next"`
}

// User GET /user
func User(c *fiber.Ctx) error {
	offset, err := strconv.Atoi(c.Query("from"))
	if err != nil {
		return err
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return err
	}

	rows, err := db.DB.Query(
		`SELECT * 
		FROM users 
		ORDER BY id ASC
		WHERE id > ? 
		LIMIT ?`,
		offset, limit,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []db.User
	var u db.User
	for rows.Next() {
		if err := rows.Scan(&u); err == nil {
			users = append(users, u)
		}
	}

	res := resp{
		Items: users,
		Cursor: cursor{
			Prev: "",
			Next: "",
		},
	}

	return c.JSON(res)
}
