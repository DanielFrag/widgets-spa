package repository

import (
	"github.com/DanielFrag/widgets-spa-rv/model"
	"github.com/DanielFrag/widgets-spa-rv/infra"
)

type WidgetRepository interface {
	CreateWidget(model.Widget) error
	GetWidgets() ([]model.Widget, error)
	GetWidgetByID(string) (model.Widget, error)
	UpdateWidget(string) error
}

func GetWidgetRepository() WidgetRepository {
	return infra.GetWidgetDB()
}
