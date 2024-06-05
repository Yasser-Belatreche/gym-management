package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
import "gym-management/src/lib/primitives/application_specific"

type GormPsqlPersistence struct {
	host     string
	user     string
	password string
	database string
	port     string
	clients  map[string]*gorm.DB
}

var instance *GormPsqlPersistence

func NewGormPsqlPersistence(host string, user string, password string, database string, port string) *GormPsqlPersistence {
	if instance != nil {
		return instance
	}

	instance = &GormPsqlPersistence{
		host:     host,
		user:     user,
		password: password,
		database: database,
		port:     port,
		clients:  make(map[string]*gorm.DB),
	}

	return instance
}

func (g *GormPsqlPersistence) Connect() *application_specific.ApplicationException {
	dns := "host=" + g.host + " user=" + g.user + " password=" + g.password + " dbname=" + g.database + " port=" + g.port + " sslmode=disable" + " TimeZone=UTC+1"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return application_specific.NewUnknownException("DATABASE_CONNECTION_ERROR", "Error connecting to the database", map[string]string{
			"error": err.Error(),
		})
	}

	g.clients["default"] = db

	return nil
}

func (g *GormPsqlPersistence) Disconnect() *application_specific.ApplicationException {
	if g.clients["default"] == nil {
		return application_specific.NewUnknownException("DATABASE_CONNECTION_ERROR", "No database connection to close", nil)
	}

	db, _ := g.clients["default"].DB()

	err := db.Close()
	if err != nil {
		return application_specific.NewUnknownException("DATABASE_CONNECTION_ERROR", "Error closing the database connection", map[string]string{
			"error": err.Error(),
		})
	}

	g.clients = make(map[string]*gorm.DB)

	return nil
}

func (g *GormPsqlPersistence) WithTransaction(session *application_specific.Session, method func() *application_specific.ApplicationException) *application_specific.ApplicationException {
	defaultDB, ok := g.clients["default"]
	if !ok {
		return application_specific.NewUnknownException("DATABASE_CONNECTION_ERROR", "No database connection to start a transaction", nil)
	}

	if _, ok := g.clients["default"]; !ok {
		return application_specific.NewUnknownException("DATABASE_CONNECTION_ERROR", "No database connection to start a transaction", nil)
	}

	defer func() {
		if r := recover(); r != nil {
			g.clients[session.CorrelationId].Rollback()
			delete(g.clients, session.CorrelationId)
			panic(r)
		}
	}()

	g.clients[session.CorrelationId] = defaultDB.Begin()

	err := method()
	if err != nil {
		g.clients[session.CorrelationId].Rollback()
		delete(g.clients, session.CorrelationId)
		return nil
	}

	g.clients[session.CorrelationId].Commit()

	delete(g.clients, session.CorrelationId)

	return nil
}

type Status struct {
	Provider string `json:"provider"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

func (g *GormPsqlPersistence) HealthCheck() struct {
	Provider string `json:"provider"`
	Status   string `json:"status"`
	Message  string `json:"message"`
} {
	if _, ok := g.clients["default"]; !ok {
		return Status{
			Status:   "DOWN",
			Provider: "GORM - PSQL",
			Message:  "No database connection",
		}
	}

	db, _ := g.clients["default"].DB()

	err := db.Ping()
	if err != nil {
		return Status{
			Status:   "DOWN",
			Provider: "GORM - PSQL",
			Message:  err.Error(),
		}
	}

	return Status{
		Status:   "UP",
		Provider: "GORM - PSQL",
		Message:  "Database connection is healthy",
	}
}

func (g *GormPsqlPersistence) GetClient(session *application_specific.Session) *gorm.DB {
	if client, ok := g.clients[session.CorrelationId]; ok {
		return client
	}

	db, ok := g.clients["default"]
	if !ok {
		panic("Database connection not initialized")
	}

	return db
}
