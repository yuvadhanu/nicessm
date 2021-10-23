package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//Context :""
type Context struct {
	CTX     context.Context
	DB      *mongo.Database
	Session mongo.Session
	Client  *mongo.Client
	SC      mongo.SessionContext
	Auth    Authentication
}
