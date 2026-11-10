package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateCharactersTable struct{}

func (m *CreateCharactersTable) GetName() string {
	return "CreateCharactersTable"
}

func (m *CreateCharactersTable) Up(con *sqlx.DB) {
	table := builder.NewTable("characters", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("user_id").Type("INT UNSIGNED").NotNull()
	table.String("name", 255).NotNull()
	table.Column("race_id").Type("INT UNSIGNED").NotNull()
	table.Integer("level").NotNull()
	table.MustExec()
}

func (m *CreateCharactersTable) Down(con *sqlx.DB) {
	builder.DropTable("characters", con).MustExec()
}
