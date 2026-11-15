package list

import (
	"github.com/jmoiron/sqlx"
	"lorekeeper/pkg/go-migrations/builder"
)

type CreateDocumentChunksTable struct{}

func (m *CreateDocumentChunksTable) GetName() string {
	return "CreateDocumentChunksTable"
}

func (m *CreateDocumentChunksTable) Up(con *sqlx.DB) {
	table := builder.NewTable("document_chunks", con)
	table.Column("id").Type("INT UNSIGNED").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("campaign_file_id").Type("INT UNSIGNED").NotNull()
	table.Integer("page_number").NotNull()
	table.Integer("chunk_index").NotNull()
	table.Column("content").Type("TEXT").NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateDocumentChunksTable) Down(con *sqlx.DB) {
	builder.DropTable("document_chunks", con).MustExec()
}
