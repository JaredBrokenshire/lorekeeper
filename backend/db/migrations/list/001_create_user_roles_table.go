package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateUserRolesTable struct{}

func (m *CreateUserRolesTable) GetName() string {
	return "CreateUserRolesTable"
}

func (m *CreateUserRolesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("user_roles", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("user_id").Type("INT UNSIGNED").NotNull()
	table.Column("role_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *CreateUserRolesTable) Down(con *sqlx.DB) {
	builder.DropTable("user_roles", con).MustExec()
}
