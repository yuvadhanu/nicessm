package daos

import (
	"context"
	"fmt"
	"log"
	"nicessm-api-service/config"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"nicessm-api-service/redis"
	"nicessm-api-service/shared"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

//Daos :""
type Daos struct {
	mongoURL string
	mongoDB  string
	Shared   *shared.Shared
	Redis    *redis.RedisCli
	Config   *config.ViperConfigReader
}

//Counter :""
type Counter struct {
	Key   string `json:"key,omitempty" bson:"key,omitempty" form:"key"`
	Value int64  `json:"value,omitempty" bson:"value,omitempty" form:"value"`
}

//CollectionRegistory : ""
type CollectionRegistory struct {
	Code   string `json:"code" bson:"code,omitempty"`
	SUffix string `json:"suffix" bson:"suffix,omitempty"`
	Prefix string `json:"prefix" bson:"prefix,omitempty"`
	Digits int    `json:"digits" bson:"digits,omitempty"`
}

//GetDB :""
func GetDB(ctx context.Context, daos *Daos) (*mongo.Database, mongo.Session, *mongo.Client) {
	mongodbURL := daos.Config.GetString(daos.Shared.GetCmdArg(constants.ENV) + "." + constants.DBURL)

	// mongodbURL := "mongodb://localhost:27018,localhost:27019,localhost:27020/?replicaSet=rsSample"
	fmt.Println(mongodbURL)

	// mongodbURL := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL))
	if err != nil {
		panic(err.Error())
	}
	opts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := client.StartSession(opts)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(daos.mongoDB), sess, client
}

//GetDBV2 :""
func (d *Daos) GetDBV2(ctx context.Context) (*mongo.Database, mongo.Session) {
	mongodbURL := d.Config.GetString(d.Shared.GetCmdArg(constants.ENV) + "." + constants.DBURL)
	fmt.Println(mongodbURL)

	// mongodbURL := "mongodb://localhost:27018,localhost:27019,localhost:27020/?replicaSet=rsSample"
	// mongodbURL := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL))
	if err != nil {
		panic(err.Error())
	}
	// opts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(d.mongoDB), sess
}

//GetDBV3 :""
func (d *Daos) GetDBV3(ctx context.Context) *mongo.Client {
	mongodbURL := d.Config.GetString(d.Shared.GetCmdArg(constants.ENV) + "." + constants.DBURL)
	fmt.Println(mongodbURL)

	// mongodbURL := "mongodb://localhost:27018,localhost:27019,localhost:27020/?replicaSet=rsSample"
	// mongodbURL := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL))
	if err != nil {
		panic(err.Error())
	}

	return client
}

//GetDBV3 :""
func (d *Daos) GetDBV4(ctx context.Context) (*mongo.Client, *mongo.Database) {
	mongodbURL := d.Config.GetString(d.Shared.GetCmdArg(constants.ENV) + "." + constants.DBURL)
	fmt.Println(mongodbURL)
	// mongodbURL := "mongodb://localhost:27018,localhost:27019,localhost:27020/?replicaSet=rsSample"
	// mongodbURL := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL))
	if err != nil {
		panic(err.Error())
	}

	return client, client.Database(d.mongoDB)
}

//GetUniqueID :""
func (d *Daos) GetUniqueID(ctx *models.Context, key string) string {
	selector := bson.M{"key": key}
	update := bson.M{"$inc": bson.M{"value": 1}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var counter Counter
	var cr CollectionRegistory
	res := ctx.DB.Collection(constants.COLLCOUNTER).FindOneAndUpdate(ctx.CTX, selector, update, &opt)
	err := res.Decode(&counter)
	fmt.Println("Decode error CC", err)
	selector = bson.M{"code": key}
	crRes := ctx.DB.Collection(constants.COLLREGISTER).FindOne(ctx.CTX, selector, options.FindOne())
	err = crRes.Decode(&cr)
	fmt.Println("Decode error CR", err)
	dig := fmt.Sprintf("%dd", cr.Digits)
	str := "%v%0" + dig + "%v"
	fmt.Println(dig, str)
	return fmt.Sprintf(str, cr.Prefix, counter.Value, cr.SUffix)
}

//GetUniqueID
func (d *Daos) GetUniqueIDV2(ctx *models.Context, key string) (string, int64) {
	selector := bson.M{"key": key}
	update := bson.M{"$inc": bson.M{"value": 1}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var counter Counter
	var cr CollectionRegistory
	res := ctx.DB.Collection(constants.COLLCOUNTER).FindOneAndUpdate(ctx.CTX, selector, update, &opt)
	err := res.Decode(&counter)
	fmt.Println("Decode error CC", err)
	selector = bson.M{"code": key}
	crRes := ctx.DB.Collection(constants.COLLREGISTER).FindOne(ctx.CTX, selector, options.FindOne())
	err = crRes.Decode(&cr)
	fmt.Println("Decode error CR", err)
	dig := fmt.Sprintf("%dd", cr.Digits)
	str := "%v%0" + dig + "%v"
	fmt.Println(dig, str)
	return fmt.Sprintf(str, cr.Prefix, counter.Value, cr.SUffix), counter.Value
}

//GetDaos : ""
func GetDaos(s *shared.Shared, Redis *redis.RedisCli, conf *config.ViperConfigReader) *Daos {
	// fmt.Println(conf.GetString(s.GetCmdArg(constants.ENV) + ".mongodb_url"))
	fmt.Println(conf.GetString(s.GetCmdArg(constants.ENV) + ".mongodb_url"))
	return &Daos{conf.GetString(s.GetCmdArg(constants.ENV) + ".mongodb_url"),
		conf.GetString(s.GetCmdArg(constants.ENV) + ".database_name"),
		s,
		Redis,
		conf,
	}
}

// EnsureIndex will create index on collection provided
func (d *Daos) EnsureIndex(ctx *models.Context, collectionName string, indexQuery []string) error {

	opts := options.CreateIndexes().SetMaxTime(10 * time.Minute)

	index := []mongo.IndexModel{}

	for _, val := range indexQuery {
		temp := mongo.IndexModel{}
		temp.Keys = bsonx.Doc{{Key: val, Value: bsonx.Int32(1)}}
		index = append(index, temp)
	}
	_, err := ctx.DB.Collection(collectionName).Indexes().CreateMany(ctx.CTX, index, opts)
	// _, err := cd.Indexes().CreateMany(context.Background(), index, opts)
	if err != nil {
		fmt.Println("Error while executing index Query" + err.Error())
		return err
	}
	return nil
}
