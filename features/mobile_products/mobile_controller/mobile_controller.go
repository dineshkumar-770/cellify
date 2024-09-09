package controller

import (
	"cellify_backend/database"
	mobilemodel "cellify_backend/features/mobile_products/mobile_model"
	"cellify_backend/response"
	"cellify_backend/utils"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type MobileInfoController struct {
	MobileInforamtion mobilemodel.MobileInfo `json:"mobile_info"`
	h                 utils.Helper
}

func (m *MobileInfoController) SaveMobileInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	res := response.Response{
		Status:  "Failed",
		Message: "Error",
	}

	mobileParams := r.Body
	err := json.NewDecoder(mobileParams).Decode(&m.MobileInforamtion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = err.Error() + ", required all parameter!"
		json.NewEncoder(w).Encode(res)
		return
	}

	id := uuid.New()
	newID := m.h.RemoveDashesFromString(id.String())
	m.MobileInforamtion.MobileID = "mobile_" + newID

	dbName := os.Getenv("DBNAME")
	db := database.MyDataBase{}
	client, err := db.DataBaseINIT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error() + ", Error in database initialization"
		json.NewEncoder(w).Encode(res)
		return
	}

	collection := client.Database(dbName).Collection("Mobiles")

	result, err := collection.InsertOne(context.Background(), m.MobileInforamtion)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error() + ", Error in saving mobile info data!"
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Status = "Success"
	res.Message = "Mobile Info saved successfully!"
	res.Resp = result
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
