package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tokenRes struct {
	Access_token string
	Token_type   string
	Expires_in   int
	Cert         string
	Issued       string
	Expires      string
}

func GetToken() gin.HandlerFunc {
	res := tokenRes{
		Access_token: "sZmbp2VNaQKh2CcNj03lVtvxoq_Oo__inYyHkmsC-YYgAsyTar7QawFlJmSrKbpMDSn7SOh7zQdWJIjFq9jMC5ALrfnANrnoFqSqYsIIZQ4WrifVzlvBOKqr0kOnBOHZZsJM-oVD7Sv5pa1QCI1c4b3iZUMvL85vDV3J4jNBaourcEwpImTZW4kHW52TXG7_AnWl3laF4DEhAMl5o7jXr1mTgRlN1BgZzKzY9QUJfslaZHeMnCUa6RLmGHbG7CAXzSfHHG2UCyJMduIX1PHCDhzIZ4OB7SxWlzJJH5L7DD2HIc6b",
		Token_type:   "bearer",
		Expires_in:   50971,
		Cert:         "",
		Issued:       "2021-05-12T11:54:38+03:00",
		Expires:      "2021-05-12T23:59:59+03:00",
	}

	return func(c *gin.Context) {

		fmt.Println(res)
		// resJson, err := json.Marshal(res)
		// if err != nil {
		// 	c.JSON(401, err)
		// }

		// c.JSON(http.StatusOK, resJson)
		c.JSON(http.StatusOK, gin.H{
			"access_token": res.Access_token,
			"token_type":   res.Token_type,
			"expires_in":   res.Expires_in,
			"cert":         res.Cert,
			"issued":       res.Issued,
			"expires":      res.Expires,
		})

	}
}

func GetTokenIdsub() gin.HandlerFunc {
	res := tokenRes{
		Access_token: "luOgfnWgsLvPGe_LjL54FP1DbsVzBREgJcBh6HvF-6y4uwaIQPmM4gnFI78gbxf-4xPQ3JEToUMEHzHHzV88ea85l8XLJL9brvgYyVgExb6UhIKGeeydr1eSfL0poUF0jbxX8fDFNMCbmREmSWMaYXea0xCbRHFo2FzlcOrJqjEh2xekJppbl8SjtNWS2OFlaMiouDMWzcoRgUB-jzcrz7-yCNAgYsjUvPwwXQ3Wt1zeWIGbIA2rUQavFtZiiDhATTzrHkaK7FnHkU3jZDN17SATSUTqC8L30V3TWfVWGPsgtQRbm7RxTpu4iOVwdTXLyrBqBg",
		Token_type:   "bearer",
		Expires_in:   50971,
		Cert:         "",
		Issued:       "2021-05-12T11:54:38+03:00",
		Expires:      "2021-05-12T23:59:59+03:00",
	}

	return func(c *gin.Context) {

		fmt.Println(res)
		// resJson, err := json.Marshal(res)
		// if err != nil {
		// 	c.JSON(401, err)
		// }

		// c.JSON(http.StatusOK, resJson)
		c.JSON(http.StatusOK, gin.H{
			"access_token": res.Access_token,
			"token_type":   res.Token_type,
			"expires_in":   res.Expires_in,
			"cert":         res.Cert,
			"issued":       res.Issued,
			"expires":      res.Expires,
		})

	}
}
