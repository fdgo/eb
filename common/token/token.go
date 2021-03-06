package token

import (
	"log"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
)

// CustomClaims 自定义的 metadata在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// Token jwt服务
type Token struct {
	rwlock     sync.RWMutex
	privateKey []byte
	conf       config.Config
}

func (srv *Token) get() []byte {
	srv.rwlock.RLock()
	defer srv.rwlock.RUnlock()

	return srv.privateKey
}

func (srv *Token) put(newKey []byte) {
	srv.rwlock.Lock()
	defer srv.rwlock.Unlock()

	srv.privateKey = newKey
}

// InitConfig 初始化
func (srv *Token) InitConfig(host, port, prefix string, path ...string) {
	etcdSource := etcd.NewSource(
		//设置配置中心地址
		etcd.WithAddress(host+":"+port),
		//设置前缀，不设置默认为 /micro/etcd
		etcd.WithPrefix(prefix),
		//是否移除前缀，这里设置为true 表示可以不带前缀直接获取对应配置
		etcd.StripPrefix(true),
	)
	srv.conf = config.NewConfig()
	err := srv.conf.Load(etcdSource)
	if err != nil {
		log.Fatal(err)
	}

	value := srv.conf.Get(path...).Bytes()
	if err != nil {
		log.Fatal(err)
	}

	srv.put(value)
	//log.Println("JWT privateKey:", string(srv.get()))
	srv.enableAutoUpdate(path...)
}

func (srv *Token) enableAutoUpdate(path ...string) {
	go func() {
		for {
			w, err := srv.conf.Watch(path...)
			if err != nil {
				log.Println(err)
			}
			v, err := w.Next()
			if err != nil {
				log.Println(err)
			}

			value := v.Bytes()
			srv.put(value)
			//log.Println("New JWT privateKey:", string(srv.get()))
		}
	}()
}

//Decode 解码
func (srv *Token) Decode(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return srv.get(), nil
	})

	if err != nil {
		return nil, err
	}
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, err
}

// Encode 将 User 用户信息加密为 JWT 字符串
// expireTime := time.Now().Add(time.Hour * 24 * 3).Unix() 三天后过期
func (srv *Token) Encode(issuer, userName string, expireTime int64) (string, error) {
	claims := CustomClaims{
		userName,
		jwt.StandardClaims{
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(srv.get())
}
