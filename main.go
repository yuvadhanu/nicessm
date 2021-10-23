package main

import (
	"fmt"
	"log"
	"net/http"
	"nicessm-api-service/config"
	"nicessm-api-service/constants"
	"nicessm-api-service/daos"
	"nicessm-api-service/handlers"
	"nicessm-api-service/middlewares"
	"nicessm-api-service/redis"
	"nicessm-api-service/routes"
	"nicessm-api-service/services"
	"nicessm-api-service/shared"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	argsWithoutProg := os.Args[1:]
	config := config.Config()
	sh := shared.NewShared(shared.SplitCmdArguments(argsWithoutProg), config)
	redisConn := redis.Connect(config)
	db := daos.GetDaos(sh, redisConn, config)
	ser := services.GetService(db, sh, redisConn, config)
	han := handlers.GetHandler(ser, sh, redisConn, config)
	route := routes.GetRoute(han, sh, redisConn, config)
	rr := mux.NewRouter()
	commonRoute := rr.PathPrefix("/api").Subrouter()
	//UIRoute := rr.PathPrefix("/ui").Subrouter()
	rr.Use(middlewares.Log)
	rr.Use(middlewares.AllowCors)
	r := commonRoute.NewRoute().Subrouter()
	//r.Use(middlewares.JWT)
	nonau := commonRoute.NewRoute().Subrouter()
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("options called")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	//UI Routes
	//route.UIRoutes(UIRoute)
	//Common Routes
	route.CommonRoutes(r)
	route.FileRoutes(r)

	//Geo Location Routes
	route.StateRoutes(r)
	route.DistrictRoutes(r)
	route.VillageRoutes(r)
	route.BlockRoutes(r)
	route.GramPanchayatRoutes(r)

	//User Routes
	route.OrganisationRoutes(r)
	route.UserOrganisationRoutes(r)
	route.UserRoutes(r)
	route.UserAuthRoutes(nonau)
	route.UserTypeRoutes(r)

	//User Language
	route.LanguageRoutes(r)

	//User Cropseason
	route.CropseasonRoutes(r)

	//User Insect
	route.InsectRoutes(r)
	//user AgroEcologicalZone
	route.AgroEcologicalZoneRoutes(r)
	//User Market
	route.MarketRoutes(r)

	//User Aidlocation
	route.AidlocationRoutes(r)

	//User ProductConfig
	route.ProductConfigRoutes(r)

	//ACL Routes
	route.ModuleRoutes(r)
	route.MenuRoutes(r)
	route.TabRoutes(r)
	route.FeatureRoutes(r)
	route.ACLMasterUserRoutes(r)
	route.ACLAccess(r)
	//SoilTypeRoutes
	route.SoilTypeRoutes(r)
	//AssetRoutes
	route.AssetRoutes(r)

	// Project Routes
	route.ProjectRoutes(r)
	// ProjectKnowledgeDomain Routes
	route.ProjectKnowledgeDomainRoutes(r)
	// KnowledgeDomain Routes
	route.KnowledgeDomainRoutes(r)
	//SubDomainRoutes
	route.SubDomainRoutes(r)
	//TopicRoutes
	route.TopicRoutes(r)
	//TopicRoutes
	route.SubTopicRoutes(r)
	//FarmerRoutes
	route.FarmerRoutes(r)
	//ClusterRoutes
	route.ClusterRoutes(r)
	//CommonLandRoutes
	route.CommonLandRoutes(r)
	//LiveStockVaccinationRoutes
	route.LiveStockVaccinationRoutes(r)
	//BlockCropRoutes
	route.BlockCropRoutes(r)
	//StateLiveStockRoutes
	route.StateLiveStockRoutes(r)

	// ProjectState Routes
	route.ProjectStateRoutes(r)
	// ProjectUser Routes
	route.ProjectUserRoutes(r)
	// AidCategoryRoutes
	route.AidCategoryRoutes(r)
	// ContentRoutes
	route.ContentRoutes(r)
	// ContentcommentRoutes
	route.ContentCommentRoutes(r)
	// ContentTranslationRoutes
	route.ContentTranslationRoutes(r)
	// ProjectFarmer Routes
	route.ProjectFarmerRoutes(r)
	// ProjectPartner Routes
	route.ProjectPartnerRoutes(r)
	// DiseaseRoutes
	route.DiseaseRoutes(r)
	// BannedItemRoutes
	route.BannedItemRoutes(r)
	// Vaccine Routes
	route.VaccineRoutes(r)
	// CommodityCategory Routes
	route.CommodityCategoryRoutes(r)
	// CommodityFunction Routes
	route.CommodityFunctionRoutes(r)
	// Commodity Routes
	route.CommodityRoutes(r)
	// CommodityStage Routes
	route.CommodityStageRoutes(r)
	// CommodityVariety Routes
	route.CommodityVarietyRoutes(r)
	// CommoditySubVariety Routes
	route.CommoditySubVarietyRoutes(r)
	// DistrictWeatherDataRoutes
	route.DistrictWeatherDataRoutes(r)
	// SelfRegisterRoutes
	route.SelfRegisterRoutes(r)
	http.DefaultClient.Timeout = time.Minute * 10
	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), rr))
}
