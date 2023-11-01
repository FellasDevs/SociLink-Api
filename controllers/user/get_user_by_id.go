package usercontroller

import (
	"SociLinkApi/models"
	userrepository "SociLinkApi/repository/user"
	"SociLinkApi/types/errors"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetUserById(idString string, db *gorm.DB) (models.User, *customerrors.RouteError) {
	id, err := uuid.Parse(idString)

	if err != nil {
		return models.User{}, customerrors.NewRouteError(http.StatusBadRequest, "id is not a valid uuid")
	}

	if user, err := userrepository.GetUserById(id, db); err != nil {
		var statusCode int

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		return models.User{}, customerrors.NewRouteError(statusCode, err.Error())
	} else {
		return user, nil
	}
}
