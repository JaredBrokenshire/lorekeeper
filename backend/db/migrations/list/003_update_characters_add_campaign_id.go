package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type UpdateCharactersAddCampaignId struct{}

func (m *UpdateCharactersAddCampaignId) GetName() string {
	return "UpdateCharactersAddCampaignId"
}

func (m *UpdateCharactersAddCampaignId) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("campaign_id").Type("INT UNSIGNED").Nullable()
	table.MustExec()
}

func (m *UpdateCharactersAddCampaignId) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("campaign_id")
	table.MustExec()
}
