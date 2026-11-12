package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type UpdateCharactersAddProfilePicture struct{}

func (m *UpdateCharactersAddProfilePicture) GetName() string {
	return "UpdateCharactersAddProfilePicture"
}

func (m *UpdateCharactersAddProfilePicture) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("profile_picture_id").Type("INT UNSIGNED").NotNull()
	table.MustExec()
}

func (m *UpdateCharactersAddProfilePicture) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("profile_picture_id")
	table.MustExec()
}
