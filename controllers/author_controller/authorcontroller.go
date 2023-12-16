package authorcontroller

import (
	"encoding/json"
	"errors"
	"go-native-api/config"
	"go-native-api/helpers"
	"go-native-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var author[]models.Author

	if error := config.DB.Find(&author).Error; error != nil {
		helpers.Response(w, http.StatusBadRequest, "Failed to get author", nil)
		return
	}

	helpers.Response(w, http.StatusOK, "success", author)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	if error := json.NewDecoder(r.Body).Decode(&author); error != nil {
		helpers.Response(w, http.StatusBadRequest, "Failed to decode request body", nil)
		return
	}
	defer r.Body.Close()


	if error := config.DB.Create(&author); error == nil {
		helpers.Response(w, http.StatusBadRequest, "Failed to create author", nil)
		return
	}

	helpers.Response(w, http.StatusOK, "Success Create Author", author)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author models.Author

	if error := config.DB.First(&author, id).Error; error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			helpers.Response(w, http.StatusBadRequest, "Author not found", nil)
			return
		}
		helpers.Response(w, http.StatusBadRequest, "Failed to get author", nil)
		return
	}

	helpers.Response(w, http.StatusOK, "success", author)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author models.Author

	if error := config.DB.First(&author, id).Error; error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			helpers.Response(w, http.StatusBadRequest, "Author not found", nil)
			return
		}
		helpers.Response(w, http.StatusBadRequest, "Failed to get author", nil)
		return
	}

	if error := json.NewDecoder(r.Body).Decode(&author); error != nil{
		helpers.Response(w, http.StatusBadRequest, error.Error(), nil)
		return
	}

	defer r.Body.Close()

	if error := config.DB.Where("id_author = ?", id).Updates(&author).Error; error != nil{
		helpers.Response(w, http.StatusBadRequest, error.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusOK, "Success Update Author", author)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author models.Author

	if error := config.DB.First(&author, id).Error; error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			helpers.Response(w, http.StatusBadRequest, "Author not found", nil)
			return
		}
		helpers.Response(w, http.StatusBadRequest, "Failed to get author", nil)
		return
	}

	if error := config.DB.Delete(&author).Error; error != nil {
		helpers.Response(w, http.StatusBadRequest, "Failed to delete author", nil)
		return
	}

	helpers.Response(w, http.StatusOK, "Success Delete Author", nil)
}