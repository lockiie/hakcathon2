package routers

import (
	ctrls "github.com/lucas/hackathon/src/controllers"
	"github.com/lucas/hackathon/src/db"
	f "github.com/lucas/hackathon/src/functions"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

const (
	empresas            = "/empresas"
	categorias          = "/categorias"
	produtos            = "/produtos"
	produtoEnviarImagem = "/produtos/enviarImagem"
	paramID             = "/:id"
	swagger             = "/swagger"
	imagens             = "/imagens"
)

func init() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static(swagger, "./static/dist")
	app.Static("/images", "./images")
	app.Get("/dashboard", monitor.New())

	api := app.Group("/api")
	defer db.Pool.Close()

	api.Static(imagens, "./images")

	api.Post(empresas, ctrls.CreateEmpresa)
	api.Put(empresas+paramID, ctrls.UpdateEmpresa)
	api.Delete(empresas+paramID, ctrls.DeleteEmpresa)
	api.Get(empresas, ctrls.QueryEmpresa)
	api.Get(empresas+paramID, ctrls.QueryEmpresaByID)

	api.Post(categorias, ctrls.CreateCategoria)
	api.Put(categorias+paramID, ctrls.UpdateCategoria)
	api.Delete(categorias+paramID, ctrls.DeleteCategoria)
	api.Get(categorias, ctrls.QueryCategoria)
	api.Get(categorias+paramID, ctrls.QueryCategoriaByID)

	api.Post(produtos, ctrls.CreateProduto)
	api.Post(produtoEnviarImagem, ctrls.UploadImage)
	api.Put(produtos+paramID, ctrls.UpdateProduto)
	api.Delete(produtos+paramID, ctrls.DeleteProduto)
	api.Get(produtos, ctrls.QueryProduto)
	api.Get(produtos+paramID, ctrls.QueryProdutoByID)
	//fmt.Println(f.GoDotEnvVariable("portAPI"))

	app.Listen(":" + f.GoDotEnvVariable("portAPI"))
}
