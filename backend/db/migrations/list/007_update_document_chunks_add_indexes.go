package list

import (
	"os"

	"github.com/jmoiron/sqlx"
)

type UpdateDocumentChunksAddIndexes struct{}

func (m *UpdateDocumentChunksAddIndexes) GetName() string {
	return "UpdateDocumentChunksAddIndexes"
}

func (m *UpdateDocumentChunksAddIndexes) Up(con *sqlx.DB) {
	con.MustExec("CREATE INDEX idx_campaign_files_campaign_id ON campaign_files (campaign_id)")
	con.MustExec("CREATE INDEX idx_document_chunks_campaign_file_id ON document_chunks (campaign_file_id)")
	con.MustExec("CREATE UNIQUE INDEX uq_document_chunks_file_chunk ON document_chunks (campaign_file_id, chunk_index)")
}

func (m *UpdateDocumentChunksAddIndexes) Down(con *sqlx.DB) {
	if os.Getenv("DB_DRIVER") == "mysql" {
		con.MustExec("DROP INDEX uq_document_chunks_file_chunk ON document_chunks")
		con.MustExec("DROP INDEX idx_document_chunks_campaign_file_id ON document_chunks")
		con.MustExec("DROP INDEX idx_campaign_files_campaign_id ON campaign_files")
		return
	}

	con.MustExec("DROP INDEX IF EXISTS uq_document_chunks_file_chunk")
	con.MustExec("DROP INDEX IF EXISTS idx_document_chunks_campaign_file_id")
	con.MustExec("DROP INDEX IF EXISTS idx_campaign_files_campaign_id")
}
