package redis

////////////
//TODO
//
//

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
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
// 1 : User Session Token
// 2 : Google Calender Integration
// 3 : open
// 4 : open

////////////////////////////////////////////////
//Structures and Variables
type RedisData struct{
	User_id 			string		`json:"user_id"`
	User_Identifier		string		`json:"user_identifier"`
	User_Name			string		`json:"user_name"`
	User_Surname		string		`json:"user_surname"`
	Token 				string 		`json:"token"`
	CreationTime 		time.Time 	`json:"CreationTime"`
	ExpirationTime 		time.Time	`json:"ExpirationTime"`
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

func getRedisClient(database int) redis.Client {
	if(database > len(redisClients)){
		//error
		return getRedisClient(0);
	}
	if redisClients[database] == nil {
		redisClients[database] = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
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
		getRedisClient(database);
	}
	_ = val
	return *redisClients[database];
}

func getAuthClient() redis.Client{
	return getRedisClient(1);
}

func getCalenderClient() redis.Client{
	return getRedisClient(2);
}

func generateToken()string{
	buff := make([]byte, int(math.Ceil(float64(128)/2)))
    n,err := rand.Read(buff)
	if err != nil{
		logger.Warn.Println(err)
	}
	_ = n
    return hex.EncodeToString(buff)[:128]
}

func generateAuthToken(rclient *redis.Client, user_id string)(string,error){
	val, err := rclient.Get(ctx, user_id).Result();
	if err != nil{
		err = rclient.Set(ctx, user_id, "", 0).Err()
		if err != nil{
			logger.Error.Println("Redis database error: cannot set user token")
			return "",err
		}
	}
	token := ""
	for true{
		token = generateToken();
		if token == string(val){
			continue;
		}
		val, err := rclient.Get(ctx, token).Result();
		if err != nil{
			break
		}
		_ = val
	}
    err = rclient.Set(ctx, user_id, token, 0).Err()
	if err != nil{
		logger.Error.Println("Redis database error: cannot set user token")
		return "",err
	}
	return token,nil;

}

func GetRequestRedisData(request *http.Request) (*RedisData,error) {
	token := string(request.Header.Get("Authorization"))
	//check "bearer "
	if(len(token) < 100){
		//log bad auth token
		fmt.Print("bad token")		
		return nil,errors.New("bad auth token");
	}
	if(token[0:7] != "bearer "){
		return nil,errors.New("bad auth token");
	}
	token = token[7:]
	return getTokenRedisData(token);
}

func getTokenRedisData(token string) (*RedisData,error){
	redisClient := getAuthClient();
	val, err := redisClient.Get(ctx, token).Result();
	if err != nil{
		return nil,errors.New("token not found")
	}
	var redisData *RedisData
	err = json.Unmarshal([]byte(val), &redisData)
	if err != nil{
		return nil,errors.New("Redis data parse error")
	}
	if(redisData != nil){
		return redisData,nil;
	}
	return nil,errors.New("no data found");
		
}

func UserLogin(user_id string, user_identifier string, user_name string, user_surname string) (*RedisData,error){
	redisClient := getAuthClient();

	token, err := generateAuthToken(&redisClient, user_id);
	if err != nil{
		return nil, err;
	}

	redisData := &RedisData{
		User_id: user_id,
		User_Identifier: user_identifier,
		User_Name: user_name,
		User_Surname: user_surname,
		Token: token,
		CreationTime: time.Now(), 
		ExpirationTime: time.Now().Add(time.Hour),
	}

	rdata,err := json.Marshal(redisData)
	if err != nil{
		logger.Error.Fatal(err)
		return nil,err
	}

	err = redisClient.Set(ctx, token, rdata, time.Until(time.Now().Add(time.Hour))).Err();
	if err != nil{
		logger.Error.Fatal(err)
		return nil,errors.New("redis database error: cannot set token data")
	}
	val, err := redisClient.Get(ctx, token).Result();
	if err != nil {
		logger.Error.Fatal(err)
		return nil,errors.New("redis database error: cannot get set token");
	}
	_ = val;

	return redisData,nil;
}
