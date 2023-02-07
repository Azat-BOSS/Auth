package Controllers

import (
	"Auth/API"
	"Auth/DB"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type TUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAllUsers(c *fiber.Ctx) error {
	rows, err := DB.DATABASE.Query("SELECT  * FROM users.users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"data":    nil,
			"success": false,
		})
	}
	defer rows.Close()

	users := []TUser{}
	for rows.Next() {
		user := TUser{}
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		users = append(users, user)
	}
	return c.Status(200).JSON(&fiber.Map{
		"error":   false,
		"data":    users,
		"success": true,
	})
}

func RegUser(c *fiber.Ctx) error {
	userModel := new(TUser)
	if err := c.BodyParser(userModel); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(userModel.Password), 3)
	convertedPass := string(passwordHash)

	_, err := DB.DATABASE.Query("INSERT INTO users.users (name, email, password) VALUES (?, ?, ?)",
		userModel.Name, userModel.Email, convertedPass)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"data":    userModel,
		"success": true,
	})
}

func LoginUser(c *fiber.Ctx) error {
	userModel := new(TUser)
	if err := c.BodyParser(userModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
		})
	}
	var DbHashPassword string
	var DbEmail string

	DB.DATABASE.QueryRow("SELECT email, password FROM users.users WHERE email = ?", userModel.Email).Scan(&DbEmail, &DbHashPassword)

	err := bcrypt.CompareHashAndPassword([]byte(DbHashPassword), []byte(userModel.Password))
	if err != nil && userModel.Email != DbEmail {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": err,
			"message": "email or password incorrect",
		})
	}

	var tokenJwt = API.CreateJWT(userModel.Email)

	cookie := fiber.Cookie{
		Name:    "access_token",
		Value:   tokenJwt,
		Expires: time.Now().Add(time.Hour * 24),
		Secure:  true,
	}
	c.Cookie(&cookie)

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	_, err := DB.DATABASE.Query("DELETE FROM users.users WHERE id = ?", paramsId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "can't delete post with id = " + paramsId,
			"success": false,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "deleted post with id = " + paramsId,
		"success": true,
	})
}
