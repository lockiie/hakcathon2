package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/lucas/hackathon/src/models"
)

//RepoBrands é uma estrutura para inserir um usuário ao banco de dados
type RepoProduto struct {
	ctx  *context.Context
	conn *sql.Conn
}

//NewRepoBrands Inicia um novo repositorio de marcas
func NewRepoProduto(ctx *context.Context, conn *sql.Conn) *RepoProduto {
	return &RepoProduto{ctx, conn}
}

//Insert é para inserir um novo usuário ao banco de dados
func (db RepoProduto) Insert(p *models.Produto) error {

	res, err := db.conn.ExecContext(
		*db.ctx,
		`INSERT INTO produto(produto, foto, descricao, valor, categoria_id, empresa_id)
		VALUES(?,?,?,?,?,?)`,
		p.Produto, p.Foto, p.Descricao, p.Valor, p.Categoria.ID, p.Empresa.ID,
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = uint32(id)

	return err
}

//Update é para atulizar uma marca
func (db RepoProduto) Update(p *models.Produto, code string) error {
	_, err := db.conn.ExecContext(
		*db.ctx,
		`UPDATE produto set produto = ?, foto = ?, descricao = ?, valor = ?, categoria_id = ?, empresa_id = ?
		WHERE id = ?`,
		p.Produto, p.Foto, p.Descricao, p.Valor, p.Categoria.ID, p.Empresa.ID,
		code,
	)

	if err != nil {
		return err
	}

	return nil
}

//Delete é deletar uma marca
func (db RepoProduto) Delete(code string) error {
	res, err := db.conn.ExecContext(
		*db.ctx,
		`DELETE produto WHERE id = ?`,
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
func (db RepoProduto) Query(c *fiber.Ctx) (*[]models.Produto, error) {
	var args []interface{}
	var where string
	where = " "
	if c.Query("produto") != "" {
		sql := "produto like ?"
		addWhere(&where, sql)
		args = append(args, bindParamLikeFull(c.Query("produto")))
	}
	if c.Query("descricao") != "" {
		sql := "descricao like ?"
		addWhere(&where, sql)
		args = append(args, bindParamLikeFull(c.Query("descricao")))
	}
	if c.Query("valor_maior") != "" {
		sql := "valor > ?"
		addWhere(&where, sql)
		args = append(args, c.Query("valor_maior"))
	}
	if c.Query("valor_menor") != "" {
		sql := "valor < ?"
		addWhere(&where, sql)
		args = append(args, c.Query("valor_menor"))
	}
	if c.Query("categoria_id") != "" {
		sql := "categoria_id = ?"
		addWhere(&where, sql)
		args = append(args, c.Query("categoria_id"))
	}
	if c.Query("empresa_id") != "" {
		sql := "empresa_id = ?"
		addWhere(&where, sql)
		args = append(args, c.Query("empresa_id"))
	}
	pag := models.Pagination{}
	pag.Limit = c.Query(LIMIT)
	pag.OffSet = c.Query(OFFSET)
	where += pag.Pag()
	rows, err := db.conn.QueryContext(
		*db.ctx,
		`SELECT P.id, P.produto, P.foto, P.descricao, P.valor,
		        C.id, C.categoria,
		        E.id,E .empresa, E.whatsapp
        FROM produto P, categoria C, empresa E
        WHERE P.categoria_id = c.id
          AND P.empresa_id = e.id`+where, args...,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var produtos []models.Produto

	for rows.Next() {
		var pro models.Produto
		rows.Scan(&pro.ID, &pro.Produto, &pro.Foto, &pro.Descricao, &pro.Valor,
			&pro.Categoria.ID, &pro.Categoria.Categoria,
			&pro.Empresa.ID, &pro.Empresa.Empresa, &pro.Empresa.WhatsApp)
		produtos = append(produtos, pro)
	}
	return &produtos, nil
}

//QueryByID retorna uma consulta do banco de dados que trans 1 registro no máximo
func (db RepoProduto) QueryByID(code string) (models.Produto, error) {

	row := db.conn.QueryRowContext(
		*db.ctx,
		`SELECT P.id, P.produto, P.foto, P.descricao, P.valor, 
		        C.id, C.categoria,
		        E.id,E .empresa, E.whatsapp
        FROM produto P, categoria C, empresa E
        WHERE P.categoria_id = c.id
          AND P.empresa_id = e.id
		  AND P.ID = ?`, code,
	)
	var pro models.Produto
	err := row.Scan(&pro.ID, &pro.Produto, &pro.Foto, &pro.Descricao, &pro.Valor,
		&pro.Categoria.ID, &pro.Categoria.Categoria,
		&pro.Empresa.ID, &pro.Empresa.Empresa, &pro.Empresa.WhatsApp)
	if err != nil {
		return pro, errors.New("Nenhum registro encontrado")
	}
	return pro, err
}
