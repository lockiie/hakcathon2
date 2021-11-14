package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/lucas/hackathon/src/models"
)

//RepoBrands é uma estrutura para inserir um usuário ao banco de dados
type RepoCategoria struct {
	ctx  *context.Context
	conn *sql.Conn
}

//NewRepoBrands Inicia um novo repositorio de marcas
func NewRepoCategoria(ctx *context.Context, conn *sql.Conn) *RepoCategoria {
	return &RepoCategoria{ctx, conn}
}

//Insert é para inserir um novo usuário ao banco de dados
func (db RepoCategoria) Insert(c *models.Categoria) error {
	res, err := db.conn.ExecContext(
		*db.ctx,
		`INSERT INTO categoria(categoria)
		VALUES(?)`,
		c.Categoria,
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = uint32(id)

	return err
}

//Update é para atulizar uma marca
func (db RepoCategoria) Update(p *models.Categoria, code string) error {
	_, err := db.conn.ExecContext(
		*db.ctx,
		`UPDATE categoria set categoria = ?
		WHERE id = ?`,
		p.Categoria,
		code,
	)

	if err != nil {
		return err
	}

	return nil
}

//Delete é deletar uma marca
func (db RepoCategoria) Delete(code string) error {
	res, err := db.conn.ExecContext(
		*db.ctx,
		`DELETE categoria WHERE id = ?`,
		code,
	)
	if err != nil {
		return nil
	}
	affected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("Nenhum registro foi deletado")
	}
	return nil
}

//Query retorna uma consulta do banco de dados
func (db RepoCategoria) Query(c *fiber.Ctx) (*[]models.Categoria, error) {
	var args []interface{}
	var where string
	if c.Query("categoria") != "" {
		sql := "categoria like ?"
		addWhere(&where, sql)
		args = append(args, bindParamLikeFull(c.Query("categoria")))
	}
	pag := models.Pagination{}
	pag.Limit = c.Query(LIMIT)
	pag.OffSet = c.Query(OFFSET)
	where += pag.Pag()
	rows, err := db.conn.QueryContext(
		*db.ctx,
		`SELECT id, categoria
		 FROM categoria`+where, args...,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categorias []models.Categoria

	for rows.Next() {
		var cat models.Categoria
		rows.Scan(&cat.ID, &cat.Categoria)
		categorias = append(categorias, cat)
	}
	return &categorias, nil
}

//QueryByID retorna uma consulta do banco de dados que trans 1 registro no máximo
func (db RepoCategoria) QueryByID(code string) (models.Categoria, error) {

	row := db.conn.QueryRowContext(
		*db.ctx,
		`SELECT id, categoria
		FROM categoria
		WHERE id = ?`, code,
	)
	var cat models.Categoria
	err := row.Scan(&cat.ID, &cat.Categoria)
	if err != nil {
		return cat, errors.New("Nenhum registro encontrado")
	}
	return cat, err
}
