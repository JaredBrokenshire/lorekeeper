package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type UpdateClassesAddLogo struct{}

func (m *UpdateClassesAddLogo) GetName() string {
	return "UpdateClassesAddLogo"
}

func (m *UpdateClassesAddLogo) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("logo_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *UpdateClassesAddLogo) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("logo_id")
	table.MustExec()
}
