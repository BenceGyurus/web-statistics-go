package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func DatabaseInitSession() error {
	cassandraHost := os.Getenv("CASSANDRA_CONTACT_POINT")
	cassandraUsername := os.Getenv("CASSANDRA_USERNAME")
	cassandraPassword := os.Getenv("CASSANDRA_PASSWORD")

	fmt.Println("Cassandra Host:", cassandraHost)

	cluster := gocql.NewCluster(cassandraHost)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: cassandraUsername, Password: cassandraPassword, AllowedAuthenticators: []string{"com.instaclustr.cassandra.auth.InstaclustrPasswordAuthenticator"}}

	sess, err := cluster.CreateSession()
	Session = sess
	if err != nil {
		log.Println("Error creating session:", err)
		return err
	} else {
		log.Println("Connected to Cassandra cluster successfully")
	}
	if err := CreateDatabases(); err == nil {
		return err
	}
	return nil
}
