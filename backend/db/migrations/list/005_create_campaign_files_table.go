package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateCampaignFilesTable struct{}

func (m *CreateCampaignFilesTable) GetName() string {
	return "CreateCampaignFilesTable"
}

func (m *CreateCampaignFilesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("campaign_files", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("campaign_id").Type("INT UNSIGNED").NotNull()
	table.Column("file_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *CreateCampaignFilesTable) Down(con *sqlx.DB) {
	builder.DropTable("campaign_files", con).MustExec()
}
