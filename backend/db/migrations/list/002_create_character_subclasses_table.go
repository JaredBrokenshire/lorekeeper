package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateCharacterSubclassesTable struct{}

func (m *CreateCharacterSubclassesTable) GetName() string {
	return "CreateCharacterSubclassesTable"
}

func (m *CreateCharacterSubclassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_subclasses", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("INT UNSIGNED").NotNull()
	table.Column("subclass_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *CreateCharacterSubclassesTable) Down(con *sqlx.DB) {
	builder.DropTable("character_subclasses", con).MustExec()
}
