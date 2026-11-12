package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateSubclassesTable struct{}

func (m *CreateSubclassesTable) GetName() string {
	return "CreateSubclassesTable"
}

func (m *CreateSubclassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("subclasses", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("class_id").Type("INT UNSIGNED").NotNull()
	table.String("name", 255).NotNull()
	table.MustExec()
}

func (m *CreateSubclassesTable) Down(con *sqlx.DB) {
	builder.DropTable("subclasses", con).MustExec()
}
