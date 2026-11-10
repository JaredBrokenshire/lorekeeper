package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateUsersTable struct{}

func (m *CreateUsersTable) GetName() string {
	return "CreateUsersTable"
}

func (m *CreateUsersTable) Up(con *sqlx.DB) {
	table := builder.NewTable("users", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("username", 255).NotNull().Unique()
	table.String("email", 255).NotNull().Unique()
	table.String("password", 255).NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateUsersTable) Down(con *sqlx.DB) {
	builder.DropTable("users", con).MustExec()
}
