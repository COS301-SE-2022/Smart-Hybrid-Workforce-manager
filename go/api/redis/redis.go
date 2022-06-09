package redis

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"lib/logger"
	"math"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/gorilla/mux"
)

////////////////////////////////////////////////
//Databases
// 0 : testing
// 1 : user session token
// 2 : open
// 3 : open
// 4 : open

////////////////////////////////////////////////
//Structures and Variables

type RedisData struct{
	User_id string				`json:"User_id"`
	CreationTime time.Time 		`json:"CreationTime"`
	ExpirationTime time.Time	`json:"ExpirationTime"`
}

// Authenticated User Struct
type AuthUserData struct{
	Token string `json:"token"`
	ExpirationTime time.Time `json:"expr_time"`
}

//Redis clients
var redisClients[5] *redis.Client

//context of the program for instance of redis running
var ctx = context.Background()

func InitializeRedisClients() error{
	logger.Info.Println("Redis Initializing!")
	for i := 0; i < 5; i++{
		if redisClients[i] == nil {
			redisClients[i] = redis.NewClient(&redis.Options{
				Addr:     "redis:6379",
				Password: "archepassword1234",
				DB:       i,
			})
			err := redisClients[i].Set(ctx, "verify", true, 0).Err()
			if err != nil {
				fmt.Println(err)
				logger.Error.Fatal(err)
				return err
			}
			val, err := redisClients[i].Get(ctx, "verify").Result()
			if err != nil {
				logger.Error.Fatal(err)
				return err
			}
			_ = val
		}
	}
	logger.Info.Println("Redis Initialized")
	return nil
} 

func GetRedisClient(database int) redis.Client {
	if(database > len(redisClients)){
		//error
		return GetRedisClient(0);
	}
	if redisClients[database] == nil {
		redisClients[database] = redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "archepassword1234",
			DB:       database,
		})
		err := redisClients[database].Set(ctx, "verify", true, 0).Err()
		if err != nil {
			fmt.Println(err)
			logger.Error.Fatal(err)
		}
		val, err := redisClients[database].Get(ctx, "verify").Result()
		if err != nil {
			fmt.Println(err)
			logger.Error.Fatal(err)
		}
		_ = val
		return *redisClients[database]

	}
	val, err := redisClients[database].Get(ctx, "verify").Result()
	if err != nil {
		redisClients[database] = nil
		GetRedisClient(database);
	}
	_ = val
	return *redisClients[database];
}

func GetAuthClient() redis.Client{
	return GetRedisClient(1);
}

func ValidateUserToken(token string) bool{
	redisClient := GetAuthClient();
	val, err := redisClient.Get(ctx, token).Result();
	if err != nil {
		return false;
	}
	_ = val
	return true;
}

func AddAuthUser(user_id string) AuthUserData{
	udata := RedisData{user_id,time.Now(), time.Now().Add(time.Hour * time.Duration(1))}
	rdata,err := json.Marshal(udata)
	aUserData := AuthUserData{}
	if err != nil{
		logger.Error.Fatal(err)
		return aUserData
	}
	utoken := GenerateUserToken()
	redisClient := GetAuthClient()

	serr := redisClient.Set(ctx, utoken, rdata, 0).Err();
	if serr != nil{
		logger.Error.Fatal(serr)
		fmt.Println(serr)
		return aUserData
	}
	val, gerr := redisClient.Get(ctx, utoken).Result();
	if gerr != nil {
		logger.Error.Fatal(gerr)
		return aUserData
	}
	_ = val
	aUserData.Token = utoken
	aUserData.ExpirationTime = udata.ExpirationTime
	return aUserData
}

func GetAuthData(token string) *RedisData{
	redisClient := GetAuthClient()
	val, err := redisClient.Get(ctx, token).Result();
	if(err != nil){
		logger.Error.Println(err)
		return nil
	}
	var rd *RedisData
	fmt.Println(string(val))
	gerr := json.Unmarshal([]byte(val), &rd)
	_ = val
	fmt.Println(gerr)
	if(gerr != nil){
		logger.Error.Println(err)
		return nil
	}
	return rd
}

func GenerateUserToken() string{
    buff:= make([]byte, int(math.Ceil(float64(128)/2)))
    n,err := rand.Read(buff)
	if err != nil{
		logger.Warn.Println(err)
	}
	_ = n
    str := hex.EncodeToString(buff)
    return str[:128]
}

func GetUserInfo(request *http.Request) *RedisData{
	token := string(request.Header.Get("Authorization"))
	//check "bearer "
	if(token[0:7] != "bearer "){
		return nil
	}
	token = token[7:]
	return GetAuthData(token)
}
