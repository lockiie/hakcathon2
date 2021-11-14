package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lucas/hackathon/src/db"
	f "github.com/lucas/hackathon/src/functions"
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

func CreateProduto(c *fiber.Ctx) error {
	p := models.Produto{}

	err := json.Unmarshal(c.Body(), &p)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	if err = p.Validators(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()
	foto := uploadImage(c)
	if foto != "" {
		p.Foto = foto
	}

	repoProduto := repositories.NewRepoProduto(&ctx, conn)

	if err := repoProduto.Insert(&p); err != nil {
		const errMessage = "Erro ao inserir produto, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(p)
}

func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("foto")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)

	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("http://"+f.GoDotEnvVariable("hostAPI")+":"+f.GoDotEnvVariable("portAPI")+"/api/imagens/%s", image)

	// create meta data and send to client

	data := map[string]interface{}{

		"foto":    image,
		"fotoUrl": imageUrl,
		"header":  file.Header,
		"tamanho": file.Size,
	}

	return c.Status(fiber.StatusCreated).JSON(data)
}
func uploadImage(c *fiber.Ctx) string {
	fmt.Println("image")
	file, err := c.FormFile("foto")
	fmt.Println("image")
	if err == nil {

		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]

		image := fmt.Sprintf("%s.%s", filename, fileExt)
		fmt.Println(image)

		c.SaveFile(file, fmt.Sprintf("./images/%s", image))
		return uniqueId.String()

	}
	return ""
}

//UpdateProduto altera uma marca
func UpdateProduto(c *fiber.Ctx) error {
	p := models.Produto{}

	err := json.Unmarshal(c.Body(), &p)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	if err = p.Validators(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	foto := uploadImage(c)
	if foto != "" {
		p.Foto = foto
	}

	repoProduto := repositories.NewRepoProduto(&ctx, conn)

	if err := repoProduto.Update(&p, c.Params(paramID)); err != nil {
		const errMessage = "Erro ao alterar produto, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

//DeleteProduto Excluir uma marca
func DeleteProduto(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	repoProduto := repositories.NewRepoProduto(&ctx, conn)
	if err := repoProduto.Delete(c.Params(paramID)); err != nil {
		const errMessage = "Erro ao excluir produto, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

//QueryProduto faz consulta de Produto e retorna um array de Produto
func QueryProduto(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()
	repoProduto := repositories.NewRepoProduto(&ctx, conn)

	Produtos, err := repoProduto.Query(c)

	if err != nil {
		const errMessage = "Erro ao buscar produto, "
		return tryErrorDB(c, errMessage, err.Error())

	}
	if len(*Produtos) > 0 {
		return c.Status(fiber.StatusOK).JSON(Produtos)
	}
	return c.SendStatus(fiber.StatusOK)

}

//QueryProdutoByID faz consulta de Produto e retorna um item de Produto
func QueryProdutoByID(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer conn.Close()
	repoProduto := repositories.NewRepoProduto(&ctx, conn)

	brand, err := repoProduto.QueryByID(c.Params(paramID))

	if err != nil {
		const errMessage = "Erro ao buscar produto, "
		return tryErrorDB(c, errMessage, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(brand)
}
