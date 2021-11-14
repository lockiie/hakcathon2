package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lucas/hackathon/src/models"
)

//RepoBrands é uma estrutura para inserir um usuário ao banco de dados
type RepoEmpresa struct {
	ctx  *context.Context
	conn *sql.Conn
}

//NewRepoBrands Inicia um novo repositorio de marcas
func NewRepoEmpresa(ctx *context.Context, conn *sql.Conn) *RepoEmpresa {
	return &RepoEmpresa{ctx, conn}
}

//Insert é para inserir um novo usuário ao banco de dados
func (db RepoEmpresa) Insert(e *models.Empresa) error {
	res, err := db.conn.ExecContext(
		*db.ctx,
		`INSERT INTO empresa(empresa, whatsapp) VALUES (?, ?)`,
		e.Empresa, e.WhatsApp,
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = uint32(id)

	return err
}

//Update é para atulizar uma marca
func (db RepoEmpresa) Update(b *models.Empresa, code string) error {
	res, err := db.conn.ExecContext(
		*db.ctx,
		`UPDATE empresa SET empresa =?, whatsapp = ? WHERE id = ?`,
		b.Empresa, b.WhatsApp, code,
	)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()

	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("Nenhuma linha foi alterada")
	}
	return nil
}

//Delete é deletar uma marca
func (db RepoEmpresa) Delete(code string) error {
	_, err := db.conn.ExecContext(
		*db.ctx,
		`DELETE empresa WHERE id = ?`,
		code,
	)
	if err != nil {
		return nil
	}
	return nil
}

//Query retorna uma consulta do banco de dados
func (db RepoEmpresa) Query(c *fiber.Ctx) (*[]models.Empresa, error) {
	var args []interface{}
	var where string
	if c.Query("empresa") != "" {
		sql := "empresa like ?"
		addWhere(&where, sql)
		args = append(args, bindParamLikeFull(c.Query("empresa")))
	}
	if c.Query("whatsapp") != "" {
		sql := "whatsapp like ?"
		addWhere(&where, sql)
		args = append(args, bindParamLikeFull(c.Query("whatsapp")))
	}
	pag := models.Pagination{}
	pag.Limit = c.Query(LIMIT)
	pag.OffSet = c.Query(OFFSET)
	where += pag.Pag()
	fmt.Println(`SELECT id, empresa, whatsapp
	FROM empresa` + where)
	rows, err := db.conn.QueryContext(
		*db.ctx,
		`SELECT id, empresa, whatsapp
		 FROM empresa`+where, args...,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var empresas []models.Empresa

	for rows.Next() {
		var empresa models.Empresa
		rows.Scan(&empresa.ID, &empresa.Empresa, &empresa.WhatsApp)
		empresas = append(empresas, empresa)
	}
	return &empresas, nil
}

//QueryByID retorna uma consulta do banco de dados que trans 1 registro no máximo
func (db RepoEmpresa) QueryByID(code string) (models.Empresa, error) {

	row := db.conn.QueryRowContext(
		*db.ctx,
		`SELECT id, empresa, whatsapp
		FROM empresa
		WHERE ID = ?`, code,
	)
	var empresa models.Empresa
	err := row.Scan(&empresa.ID, &empresa.Empresa, &empresa.WhatsApp)
	if err != nil {
		return empresa, errors.New("Nenhum registro encontrado")
	}
	return empresa, err
}
