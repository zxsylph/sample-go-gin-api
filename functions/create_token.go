package functions

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username string) (string, error) {
	fmt.Println("CreateToken")

	secretKey := []byte(os.Getenv("JWT_USER_SECRET"))
	expired, err := strconv.ParseInt(os.Getenv("JWT_USER_EXPIRE"), 10, 64)
	if err != nil {
		// if can't convert string to uint
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(time.Minute * time.Duration(expired)).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	// claims := &datatype.SignedDetails{
	// 	ID:     id,
	// 	UserId: username,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * time.Duration(expired))},
	// 	},
	// }

	// fmt.Printf("claims: %v\n", claims)

	// rcpt := jose.Recipient{
	// 	Algorithm:  jose.PBES2_HS256_A128KW,
	// 	Key:        os.Getenv("JWT_USER_SECRET"),
	// 	PBES2Count: 4096,
	// 	PBES2Salt:  []byte(os.Getenv("JWT_USER_SECRET")),
	// }

	// encrypter, err := jose.NewEncrypter(jose.A128CBC_HS256, rcpt, nil)
	// if err != nil {
	// 	return "", err
	// }

	// claimsStr, err := json.Marshal(claims)
	// if err != nil {
	// 	return "", err
	// }

	// fmt.Printf("claimsStr: %s\n", string(claimsStr[:]))

	// object, err := encrypter.Encrypt(claimsStr)
	// if err != nil {
	// 	return "", err
	// }

	// serialized, err := object.CompactSerialize()
	// if err != nil {
	// 	return "", err
	// }

	return tokenString, nil
}
