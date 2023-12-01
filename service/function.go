package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
)

func HashSHA256(input string) string { // 创建hash256
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func VerifyHash(input string, verify string) bool { // 验证hash256
	recomputedHash := HashSHA256(input)
	return recomputedHash == verify
}

// 中间件

func (_ *Utility) UserPasswdVerify(userName string, passwd string) bool { // 用户密码认证函数
	DBData, err := DataBase.ReadOneDB("user", bson.D{bson.E{Key: "name", Value: userName}}, gin.H{})
	if err != nil {
		return false
	}
	if DBData["name"] != userName {
		return false
	}

	if DBData["passwd"] != passwd {
		return false
	}
	return true
}

func (_ *Utility) ReturnHeader() gin.HandlerFunc { // 通过cookie认证
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Allow", "GET,POST,PUT,DELETE")
	}
}

func (_ *Utility) CheckLoginMiddleware() gin.HandlerFunc { // 通过cookie认证
	return func(ctx *gin.Context) {
		// cookie, err := ctx.Cookie("")
	}
}

func (_ *Utility) verifyHeaderLoginCode() gin.HandlerFunc { // 通过请求头认证
	return func(ctx *gin.Context) {

	}
}

func (_ *Utility) NoFoundPage(ctx *gin.Context) { // 404
	ctx.JSON(http.StatusNotFound, GeneralJSONHeader{
		Code: http.StatusNotFound,
		Msg:  "页面走丢了~",
		Path: ctx.Request.URL.Path,
		Data: nil,
	})
}

func (_ *Utility) Log(data ...any) { // Debug输出
	if gin.Mode() == gin.DebugMode {
		fmt.Println(data...)
	}
}

func GetEvn(key string) string {
	return os.Getenv(key)
}

func ServerBegin() { // 服务器启动前配置函数
	// 初始化环境变量
	err := godotenv.Load("Service.env")
	if err != nil {
		utilityFunction.Log(err)
	}

	DataBase = DBService{ // 初始化数据库函数
		Client: InitDB(),
	}
}
