package repository

import (
	"github.com/DanielFrag/widgets-spa/model"
)

type WidgetRepository interface {
	CreateWidget(model.Widget) (string, error)
	GetWidgets() ([]model.Widget, error)
	GetWidgetByID(string) (model.Widget, error)
	UpdateWidget(string) error
}
