package usercontroller

import (
	"SociLinkApi/models"
	userrepository "SociLinkApi/repository/user"
	"SociLinkApi/types/errors"
	"gorm.io/gorm"
	"net/http"
)

func GetUsersByName(search string, db *gorm.DB) ([]models.User, *customerrors.RouteError) {
	if search == "" {
		return nil, customerrors.NewRouteError(http.StatusBadRequest, "search cannot be empty")
	}

	if user, err := userrepository.GetUsersByName(search, db); err != nil {
		return nil, customerrors.NewRouteError(http.StatusInternalServerError, err.Error())
	} else {
		return user, nil
	}
}
