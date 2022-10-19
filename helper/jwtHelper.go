package helper

import (
	"context"
	"fmt"
	"time"

	"example.com/m/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TokenHelper interface {
	CreateJWT(string) (*TokenDetails, error)                   // id를 이용해 토큰 구조체를 반환
	CreateAuth(*TokenDetails) error                            // 토큰 구조체의 Uuid:Token으로 redis에 저장
	ExtractToken(*gin.Context) string                          // 해당 콘텍스트의 토큰을 추출
	VerifyToken(*gin.Context) (*jwt.Token, error)              // ExtractToken -> 토큰의 서명방법이 유효한지 확인
	TokenVaild(*gin.Context) error                             // ExtractToken -> VerifyToken -> 토큰이 위조되었는지 확인
	ExtractTokenMetadata(*gin.Context) (*AccessDetails, error) // ExtractToken -> VerifyToken -> token에서 redis 접근을 위한 uuid와 유저 id 추출
	FetchAuth(*AccessDetails) (string, error)                  // uuid로 redis에 접근해 존재하는지를 확인
}

type TokenDetails struct { // token에 들어가야야 하는 정보를 담은 구조체
	AccessToken string
	//RefreshToken string
	AccessUuid string // redis에 저장할 때 key값
	//RefreshUuid  string // redis에 저장할 때 key값
	AtExpires int64
	//RtExpires    int64
}

type AccessDetails struct { // redis에 접근하기 위한 구조체
	AccessUuid string
	UserId     string
}

var accessString = []byte("secret") // access token의 key
//var refreshString = []byte("resecret") // refresh token의 key

// 토큰 제작
func CreateJWT(id string) (*TokenDetails, error) {

	var err error

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix() // 엑세스 토큰 15분
	td.AccessUuid = uuid.NewString()

	//td.RtExpires = time.Now().Add(time.Hour * 1).Unix() // 리프레시 토큰 1시간
	//td.RefreshUuid = uuid.NewString()

	// 엑세스 토큰
	atClaims := jwt.MapClaims{}
	atClaims["id"] = id
	atClaims["exp"] = td.AtExpires

	atClaims["access_uuid"] = td.AccessUuid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims) // claims를 담아 토큰 만들기
	td.AccessToken, err = at.SignedString([]byte(accessString))
	if err != nil {
		return nil, err
	}

	// 리프레시 토큰
	/*rtClaims := jwt.MapClaims{}
	rtClaims["id"] = id
	rtClaims["exp"] = td.RtExpires
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.RefreshToken, err = rt.SignedString([]byte(refreshString))
	if err != nil {
		return nil, err
	}*/

	return td, nil
}

// 토큰 정보를 redis에 저장
func CreateAuth(td *TokenDetails) error {

	at := time.Unix(td.AtExpires, 0)
	//rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := redis.Client.Set(context.Background(), td.AccessUuid, td.AccessToken, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	/*errRefresh := redis.Client.Set(context.Background(), td.RefreshUuid, td.RefreshToken, rt.Sub(now)).Err()
	if errAccess != nil {
		return errRefresh
	}*/

	return nil
}

// token 추출
func ExtractToken(c *gin.Context) string {

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		return ""
	}

	return token
}

// token의 파싱 방법을 찾아 비교
func VerifyToken(c *gin.Context) (*jwt.Token, error) {

	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method : %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// token 유효성 검사
func TokenVaild(c *gin.Context) error {

	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// token에서 redis에 접근하기위 한 access uuid와 유저의 id를 추출
func ExtractTokenMetadata(c *gin.Context) (*AccessDetails, error) {

	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId := fmt.Sprintf("%s", claims["id"])

		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

// redis에 저장된 값이 없거나 토큰의 유효기간이 만료(자연적으로) 되었는지 확인
func FetchAuth(authD *AccessDetails) (string, error) {

	userId, err := redis.Client.Get(context.Background(), authD.AccessUuid).Result()
	if err != nil {
		return "", nil
	}

	return userId, nil
}
