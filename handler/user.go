package handler

import (
	dblayer "cloud-storage/db"
	"cloud-storage/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	pwd_salt        = "#890"
	tokenExpiration = 24 * time.Hour // token validity period
)

// signup user to database
func SingupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid Parameter"))
		return
	}

	enc_passwd := util.Sha1([]byte(passwd + pwd_salt))
	suc := dblayer.UserSignUp(username, enc_passwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}

}

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		data, err := ioutil.ReadFile("./static/view/home.html")
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(data)
// 		return
// 	}
// }

// handler for user sign-in
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodGet {
	// 	data, err := ioutil.ReadFile("./static/view/signin.html")
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Write(data)
	// 	return
	// }
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
		return
	}
	fmt.Println("User signin request")
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	enc_passwd := util.Sha1([]byte(password + pwd_salt))

	// 1. verify username && password
	pwd_checked := dblayer.UserSignIn(username, enc_passwd)
	if !pwd_checked {
		fmt.Println("Password not correct!")

		w.Write([]byte("FAILED"))
		return
	}
	fmt.Println("On sign in success")

	// 2. generate auth token
	token := GenToken(username)
	result := dblayer.UpdateToken(username, token)
	fmt.Println("New Token" + token)
	if !result {

		w.Write([]byte("FAILED"))
		return
	}

	// 3. redirect to main page
	// response := map[string]string{
	// 	"Location": "/user/home",
	// }
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)
	// fmt.Println("On sign in success, should redirect to home page")
	// fmt.Println("http://" + r.Host + "/static/view/home.html")
	// // w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. parse args
	r.ParseForm()
	username := r.Form.Get("username")
	// token := r.Form.Get("token")
	// isValid := IsTokenValid(token, username)

	// // check if token valid
	// if !isValid {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }
	// query user info
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// encapsulate response msg
	resp := util.RespMsg{
		Code: 0,
		Msg:  "ok",
		Data: user,
	}
	w.Write(resp.JSONBytes())

}

func GenToken(username string) string {
	// 40 digits token: md5(username + timestamp + token_salt) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	token_prefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return token_prefix + ts[:8]
}

func IsTokenValid(token string, username string) bool {

	if len(token) != 40 {
		fmt.Println("Token Length Not Matched")
		return false
	}
	// check if token expires
	tokenTimeHex := token[32:]
	// extract token timestamp
	var tokenTimeInt int64
	_, err := fmt.Sscanf(tokenTimeHex, "%x", &tokenTimeInt)
	// tokenTimeInt, err := fmt.Sscanf(tokenTimeHex, "%x", new(int64))
	if err != nil {
		fmt.Println("Invalid Token")
		return false
	}
	tokenTime := time.Unix(tokenTimeInt, 0)
	if time.Since(tokenTime) > tokenExpiration {
		fmt.Println("Token Expired")
		return false
	}
	// get user_token from user database
	userToken, _ := dblayer.GetTokenFromDB(username)
	return userToken == token
}
