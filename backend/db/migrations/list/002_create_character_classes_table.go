package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateCharacterClassesTable struct{}

func (m *CreateCharacterClassesTable) GetName() string {
	return "CreateCharacterClassesTable"
}

func (m *CreateCharacterClassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_classes", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("INT UNSIGNED").NotNull()
	table.Column("class_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *CreateCharacterClassesTable) Down(con *sqlx.DB) {
	builder.DropTable("character_classes", con).MustExec()
}
