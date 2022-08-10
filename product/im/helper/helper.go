package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"im/define"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("im-key")

//生成token
func GenerateToken(identity, email string) (string, error) {
	UserClaim := &UserClaims{
		Identity:       identity,
		Email:          email,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

//解析token
func ParserToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return nil, fmt.Errorf("parser Token Error:%v", err)
	}
	return userClaim, nil
}

func SendMail(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <873677408@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码:<b>" + code + "</b>")
	// err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "873677408@qq.com", "password123", "smtp.qq.com"))
	//返回EOF时，关闭SSL重试
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "873677408@qq.com", define.MailPassword, "smtp.qq.com"), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.163.com:465",
	})

	if err != nil {
		return err
	}
	return nil
}

func GetUUID() string {
	return uuid.NewV4().String()
}

//生成验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
