package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateRacesTable struct{}

func (m *CreateRacesTable) GetName() string {
	return "CreateRacesTable"
}

func (m *CreateRacesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("races", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.MustExec()
}

func (m *CreateRacesTable) Down(con *sqlx.DB) {
	builder.DropTable("races", con).MustExec()
}
