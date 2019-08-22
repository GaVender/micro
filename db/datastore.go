package db

import "gopkg.in/mgo.v2"

const host = ""

var session *mgo.Session

func init() {
	var err error
	session, err = mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
}

func GetSession() *mgo.Session {
	return session.Copy()
}

func CloseSession() {
	session.Close()
}
