package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type UpdateFilesAddAuthor struct{}

func (m *UpdateFilesAddAuthor) GetName() string {
	return "UpdateFilesAddAuthor"
}

func (m *UpdateFilesAddAuthor) Up(con *sqlx.DB) {
	table := builder.ChangeTable("files", con)
	table.Column("author_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *UpdateFilesAddAuthor) Down(con *sqlx.DB) {
	table := builder.ChangeTable("files", con)
	table.DropColumn("author_id")
	table.MustExec()
}
