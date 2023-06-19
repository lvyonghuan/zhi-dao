package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"regexp"
	"strconv"
	"time"
)

const tokenSecret = "MHcCAQEEIDmdyHY5w5w24RA1embdpeFjAORml1L9LhX2E3HFFHHhoAoGCCqGSM49AwEHoUQDQgAETZfbJRz5nkLy/mgwWUDURpiz3ZpMhEdw7SLQq1axt84zMSjGHvJOX2rcEzFsWo9E/GmVvdFUoDPNl1WIOQTIqg=="        //token jwt秘钥
const refreshTokenSecret = "MHgCAQEEIQDN7vInhZuZ5TNnTSIDqg2Ibf/GT+0g2rFijapWuGAlO6AKBggqhkjOPQMBB6FEA0IABBvh0D7dq+/WDqpQ/7dG0AAA2vfwqtbqhaDiu/KC0kAZVpi1JUcECfJidyu0KBjvZV6DdpqaVzMyUoWJ7671rVs=" //刷新 jwt秘钥

func generateJWT(secret string, claims jwt.Claims) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("生成token错误：" + err.Error())
	}
	return tokenString, nil
}

func parseJWT(tokenString, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("ParseWithClaims错误：" + err.Error())
	}
	if claims, ok := token.Claims.(*jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("parseJWT错误：" + err.Error())
}

// 创造token和refresh_token
func createTokenAndRefreshToken(id string) (token, refreshToken string, err error) {
	tokenClaims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 12).String(),
	}
	token, err = generateJWT(tokenSecret, tokenClaims)
	if err != nil {
		return "", "", err
	}
	refreshTokenClaims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).String(),
	}
	refreshToken, err = generateJWT(refreshTokenSecret, refreshTokenClaims)
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, err
}

// CheckExp 检测token是否过期
func CheckExp(token, secret string) (err error, id int) {
	mapClaims, err := parseJWT(token, secret)
	if err != nil {
		log.Println("jwt解密错误")
		return errors.New("Error parsing JWT:" + err.Error()), 0
	}
	expStr := (*mapClaims)["exp"].(string)
	re := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d+) ([+-]\d{4}) ([A-Z]+) m=(.*)`)
	match := re.FindStringSubmatch(expStr)
	if match == nil {
		log.Println("解析错误")
		return errors.New("解析时间戳错误"), 0
	}
	exp, err := time.Parse("2006-01-02 15:04:05.9999999 -0700 MST", match[1]+" "+match[2]+" "+match[3])
	if err != nil {
		return errors.New("token time.Parse错误：" + err.Error()), 0
	}
	if time.Now().After(exp) {
		return errors.New("token过期"), 0
	}
	idStr := (*mapClaims)["id"].(string)
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return errors.New("token 字符转换错误：" + err.Error()), 0
	}
	return nil, id
}
