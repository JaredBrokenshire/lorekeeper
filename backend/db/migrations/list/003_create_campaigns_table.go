package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateCampaignsTable struct{}

func (m *CreateCampaignsTable) GetName() string {
	return "CreateCampaignsTable"
}

func (m *CreateCampaignsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("campaigns", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.Column("dungeon_master_id").Type("INT UNSIGNED").NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateCampaignsTable) Down(con *sqlx.DB) {
	builder.DropTable("campaigns", con).MustExec()
}
