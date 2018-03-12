package repository

import (
	"github.com/DanielFrag/widgets-spa-rv/infra"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

type WidgetRepository interface {
	CreateWidget(model.Widget) error
	GetWidgets() ([]model.Widget, error)
	GetWidgetByID(string) (model.Widget, error)
	UpdateWidget(string, map[string]interface{}) error
}

func GetWidgetRepository() WidgetRepository {
	return infra.GetWidgetDB()
}
