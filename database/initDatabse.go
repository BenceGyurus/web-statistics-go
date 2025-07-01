package database

func CreateDatabases() error {
	if err := Session.Query("CREATE KEYSPACE IF NOT EXISTS statistics WITH REPLICATION = { 'class': 'SimpleStrategy', 'replication_factor': 1 }").Exec(); err != nil {
		return err
	}

	if err := Session.Query(`
		CREATE TABLE IF NOT EXISTS statistics.traffic (
			id UUID PRIMARY KEY,
			sessionId UUID,
			timestamp TIMESTAMP,
			path TEXT,
			site TEXT,
		)
	`).Exec(); err != nil {
		return err
	}

	return nil
}
