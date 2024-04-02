package internal

import (
	"log"
	"os"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
	"github.com/golang-migrate/migrate/v4"
	cassandra "github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var session *gocql.Session

// InitDB initializes the database connection
func InitDB() error {
	var err error

	//init envs
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the DB connection string from environment variable
	dbConnString := os.Getenv("DB_CONN_STRING")
	dbKeySpace := os.Getenv("DB_KEY_SPACE")
	if dbConnString == "" || dbKeySpace == "" {
		log.Fatal("DB_CONN_STRING or DB_KEY_SPACE environment variables not set")
	}

	// Connect to the cluster
	cluster := gocql.NewCluster(dbConnString)
	cluster.Keyspace = dbKeySpace
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout=20000000000
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	runMigrations(dbKeySpace)

	return nil
}

// GetSession returns the current database session
func GetSession() *gocql.Session {
	return session
}

// CloseDB closes the database session
func CloseDB() {
	if session != nil {
		session.Close()
	}
}

func runMigrations(dbKeySpace string) {
	driver, err := cassandra.WithInstance(session, &cassandra.Config{
		KeyspaceName: dbKeySpace,
	})
	if err != nil {
		log.Fatalf("could not create Cassandra migration driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/",
		dbKeySpace, driver)
	if err != nil {
		log.Fatalln("migration failed: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalln("an error occurred while applying migrations: %w", err)
	}
}
