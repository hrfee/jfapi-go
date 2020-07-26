package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func admin(c echo.Context) error {
	

func main() {
	tempPath, _ := filepath.Abs("/home/hrfee/.jf-accounts/user_configuration.json")
	file, _ := ioutil.ReadFile(tempPath)
	var jsonConfig map[string]interface{}
	json.Unmarshal(file, &jsonConfig)
	jf := new(Jellyfin)
	var server, username, password string
	fmt.Printf("server: ")
	fmt.Scanf("%s", &server)
	fmt.Printf("username: ")
	fmt.Scanf("%s", &username)
	fmt.Printf("password: ")
	fmt.Scanf("%s", &password)
	jf.init(server, "test", "test", "test", "test")
	jf.authenticate(username, password)
	fmt.Println(jf.setConfiguration("fe7fae23cbb142c2a3b6f8b9ecaabe8a", jsonConfig))
	displayprefs, _, _ := jf.getDisplayPreferences("fe7fae23cbb142c2a3b6f8b9ecaabe8a")
	fmt.Println(displayprefs)
	fmt.Println(jf.setDisplayPreferences("fe7fae23cbb142c2a3b6f8b9ecaabe8a", displayprefs))
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
