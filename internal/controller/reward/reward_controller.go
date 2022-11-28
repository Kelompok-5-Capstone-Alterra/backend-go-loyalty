package controller

import (
	"backend-go-loyalty/internal/entity"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var instance *gorm.DB

func CreateReward(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reward entity.Reward
	json.NewDecoder(r.Body).Decode(&reward)
	instance.Create(&reward)
	json.NewEncoder(w).Encode(reward)
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	var rewards []entity.Reward
	instance.Find(&rewards)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rewards)
}

func FindRewardById(w http.ResponseWriter, r *http.Request) {
	rewardId := mux.Vars(r)["id"]
	if !checkIfRewardExists(rewardId) {
		json.NewEncoder(w).Encode("Reward Not Found!")
		return
	}
	var reward entity.Reward
	instance.First(&reward, rewardId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reward)
}

func UpdateReward(w http.ResponseWriter, r *http.Request) {
	rewardId := mux.Vars(r)["id"]
	if !checkIfRewardExists(rewardId) {
		json.NewEncoder(w).Encode("Reward Not Found!")
		return
	}
	var reward entity.Reward
	instance.First(&reward, rewardId)
	json.NewDecoder(r.Body).Decode(&reward)
	instance.Save(&reward)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reward)
}

func DeleteReward(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rewardId := mux.Vars(r)["id"]
	if !checkIfRewardExists(rewardId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Reward Not Found!")
		return
	}
	var reward entity.Reward
	instance.Delete(&reward, rewardId)
	json.NewEncoder(w).Encode("Reward Deleted Successfully!")
}

func checkIfRewardExists(rewardId string) bool {
	var reward entity.Reward
	instance.First(&reward, rewardId)
	if reward.RewardID == 0 {
		return false
	}
	return true
}
