package controller

import (
	"cellify_backend/database"
	"cellify_backend/features/addusers/models"
	"cellify_backend/response"
	"cellify_backend/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	User   models.User `json:"user"`
	helper utils.Helper
}

func (u *UserController) SaveUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := response.Response{
		Status:  "Failed",
		Message: "Error",
	}
	// var userInfo models.User
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&u.User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = err.Error() + ", required all parameter!"
		json.NewEncoder(w).Encode(res)
		return
	}

	dbName := os.Getenv("DBNAME")
	db := database.MyDataBase{}
	client, err := db.DataBaseINIT()
	if err != nil {
		log.Fatal(err)
	}
	id := uuid.New()
	newID := u.helper.RemoveDashesFromString(id.String())
	//assigning user its ID----
	u.User.UserId = "user_" + newID

	//encrypting password---
	hPass := u.helper.EncryptPassword(params["password"])
	u.User.Password = hPass

	coll := client.Database(dbName).Collection("Users")
	result, err := coll.InsertOne(context.Background(), u.User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = err.Error() + ", Error saving user info"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Status = "Success"
	res.Message = "User saved successfully!"
	res.Resp = result
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	res := response.Response{
		Status:  "Failed",
		Message: "Error",
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = err.Error() + ", required all parameter!"
		json.NewEncoder(w).Encode(res)
		return
	}

	userEmail := user.Email
	userPassword := user.Password
	dbName := os.Getenv("DBNAME")
	db := database.MyDataBase{}
	client, err := db.DataBaseINIT()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(dbName).Collection("Users")

	filter := bson.M{"email": userEmail}

	result := collection.FindOne(context.Background(), filter).Decode(&u.User)
	if result == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
		res.Message = "It seems you are not registered with us!"
		json.NewEncoder(w).Encode(res)
		return
	} else if result != nil {
		w.WriteHeader(http.StatusNotFound)
		res.Message = "Error in reteriving of your info! "
		json.NewEncoder(w).Encode(res)
		return
	} else {
		isPwdMatch := u.helper.ComparePassowrds(u.User.Password,userPassword)
		fmt.Println(u.User.Password,userPassword)

		if isPwdMatch {
			w.WriteHeader(http.StatusOK)
			res.Message = "Login Successful!"
			res.Resp = u.User
			json.NewEncoder(w).Encode(res)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			res.Message = "Incorrect Password! "
			json.NewEncoder(w).Encode(res)
			return
		}
	}
}
