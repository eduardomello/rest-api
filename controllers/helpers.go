package controllers

import "gopkg.in/mgo.v2/bson"

func GetId(vars map[string]string) bson.ObjectId {

	id := vars["todoId"]

	if !bson.IsObjectIdHex(id) {
		return ""
	}

	return bson.ObjectIdHex(id)
}
