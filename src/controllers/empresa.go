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

func CreateEmpresa(c *fiber.Ctx) error {
	e := models.Empresa{}

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

	repoEmpresa := repositories.NewRepoEmpresa(&ctx, conn)

	if err := repoEmpresa.Insert(&e); err != nil {
		const errMessage = "Erro ao inserir empresa, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(e)
}

//UpdateEmpresa altera uma marca
func UpdateEmpresa(c *fiber.Ctx) error {
	e := models.Empresa{}

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

	repoEmpresa := repositories.NewRepoEmpresa(&ctx, conn)

	if err := repoEmpresa.Update(&e, c.Params(paramID)); err != nil {
		const errMessage = "Erro ao alterar empresa, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

//DeleteEmpresa Excluir uma marca
func DeleteEmpresa(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	repoEmpresa := repositories.NewRepoEmpresa(&ctx, conn)
	if err := repoEmpresa.Delete(c.Params(paramID)); err != nil {
		const errMessage = "Erro ao excluir empresa, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

//QueryEmpresa faz consulta de Empresa e retorna um array de Empresa
func QueryEmpresa(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()
	repoEmpresa := repositories.NewRepoEmpresa(&ctx, conn)

	Empresas, err := repoEmpresa.Query(c)

	if err != nil {
		const errMessage = "Erro ao buscar empresa, "
		return tryErrorDB(c, errMessage, err.Error())

	}
	if len(*Empresas) > 0 {
		return c.Status(fiber.StatusOK).JSON(Empresas)
	}
	return c.SendStatus(fiber.StatusOK)

}

//QueryEmpresaByID faz consulta de Empresa e retorna um item de Empresa
func QueryEmpresaByID(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer conn.Close()
	repoEmpresa := repositories.NewRepoEmpresa(&ctx, conn)

	brand, err := repoEmpresa.QueryByID(c.Params(paramID))

	if err != nil {
		const errMessage = "Erro ao buscar empresa, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(brand)
}
