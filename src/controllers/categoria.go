package controllers

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/lucas/hackathon/src/db"
	"github.com/lucas/hackathon/src/models"
	"github.com/lucas/hackathon/src/repositories"
)

/*

{
    "name":"Tilibra",
    "description": "<p>ótima marca de cadernos</p>",
    "active": true,
    "id": "A1"
}

*/

func CreateCategoria(c *fiber.Ctx) error {
	e := models.Categoria{}

	err := json.Unmarshal(c.Body(), &e)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	if err = e.Validators(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	repoCategoria := repositories.NewRepoCategoria(&ctx, conn)

	if err := repoCategoria.Insert(&e); err != nil {
		const errMessage = "Erro ao inserir categoria, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(e)
}

//UpdateCategoria altera uma marca
func UpdateCategoria(c *fiber.Ctx) error {
	e := models.Categoria{}

	err := json.Unmarshal(c.Body(), &e)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	if err = e.Validators(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	repoCategoria := repositories.NewRepoCategoria(&ctx, conn)

	if err := repoCategoria.Update(&e, c.Params(paramID)); err != nil {
		const errMessage = "Erro ao alterar categoria, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

//DeleteCategoria Excluir uma marca
func DeleteCategoria(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	repoCategoria := repositories.NewRepoCategoria(&ctx, conn)
	if err := repoCategoria.Delete(c.Params(paramID)); err != nil {
		const errMessage = "Erro ao excluir categoria, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

//QueryCategoria faz consulta de Categoria e retorna um array de Categoria
func QueryCategoria(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()
	repoCategoria := repositories.NewRepoCategoria(&ctx, conn)

	Categoria, err := repoCategoria.Query(c)

	if err != nil {
		const errMessage = "Erro ao buscar categoria, "
		return tryErrorDB(c, errMessage, err.Error())

	}
	if len(*Categoria) > 0 {
		return c.Status(fiber.StatusOK).JSON(Categoria)
	}
	return c.SendStatus(fiber.StatusOK)
}

//QueryCategoriaByID faz consulta de Categoria e retorna um item de Categoria
func QueryCategoriaByID(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer conn.Close()
	repoCategoria := repositories.NewRepoCategoria(&ctx, conn)

	brand, err := repoCategoria.QueryByID(c.Params(paramID))

	if err != nil {
		const errMessage = "Erro ao buscar categoria, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(brand)
}
