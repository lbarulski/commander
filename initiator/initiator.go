package main

import (
	"commander/libs/config"
	"commander/libs/userProviders/custom"
	"commander/libs/userProviders/github"
	"commander/libs/userProviders"
	"os/exec"
	"os"
	"io/ioutil"
	"strings"
	"os/user"
	"strconv"
)

func main() {
	cfg := config.Load()

	providers := []userProviders.Provider{
		custom.Initialize(cfg),
		github.Initialize(cfg),
	}

	var usersList []*userProviders.User
	for _, p := range providers {
		users := p.GetUsers()
		for _, u := range users {
			usersList = append(usersList, u)
		}
	}

	for _, u := range usersList {
		createUser(u)
	}
}

func createUser(u *userProviders.User) {
	//fmt.Println("Creating user: " + u.Username)
	//return;

	err := exec.Command("/usr/sbin/adduser", "-g", "GECOS",  "-s", "/usr/local/bin/shell", "-D", u.Username).Run()
	if err != nil {
		panic(err.Error())
	}

	os.Mkdir("/home/" + u.Username + "/.ssh", 0700)

	keys := strings.Join(u.PublicKeys, "\n")

	err = ioutil.WriteFile("/home/" + u.Username + "/.ssh/authorized_keys", []byte(keys), 0644)

	if err != nil {
		panic(err.Error())
	}

	usr, err := user.Lookup(u.Username)
	if err != nil {
		panic(err.Error())
	}

	uid, _ := strconv.Atoi(usr.Uid)
	gid, _ := strconv.Atoi(usr.Gid)

	for _, pth := range []string{
		"/home/" + u.Username + "/.ssh",
		"/home/" + u.Username + "/.ssh/authorized_keys",
		} {
		err = os.Chown(pth, uid, gid)

		if err != nil {
			panic(err.Error())
		}
	}
}