package database

import (
	"fmt"
	"strconv"

	"github.com/gocql/gocql"
)

type Cassandra struct {
	Session *gocql.Session
}

func (cassandra *Cassandra) Count(table string) (int64, error) {
	var count int64
	err := cassandra.Session.
		Query("SELECT COUNT(*) FROM " + table).
		Consistency(gocql.One).
		Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (cassandra *Cassandra) Insert(
	ttl int64,
	table string,
	fields string,
	values ...interface{},
) error {
	marks := "?"
	for index := 1; index < len(values); index++ {
		marks = marks + ", ?"
	}
	statement := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) IF NOT EXISTS",
		table,
		fields,
		marks,
	)
	if ttl > 0 {
		statement = statement + " USING TTL " + strconv.FormatInt(ttl, 10)
	}

	existingData := make(map[string]interface{})
	inserted, err := cassandra.Session.
		Query(statement, values...).
		Consistency(gocql.Quorum).
		MapScanCAS(existingData)
	if err != nil {
		return err
	}
	if !inserted {
		return fmt.Errorf("Already exists")
	}

	return nil
}

func (cassandra *Cassandra) Close() {
	if !cassandra.Session.Closed() {
		cassandra.Session.Close()
	}
}

func CassandraConnect(server string, keyspace string) (*Cassandra, error) {
	cluster := gocql.NewCluster(server)
	cluster.Keyspace = keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return &Cassandra{Session: session}, nil
}
