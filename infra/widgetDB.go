package infra

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

type WidgetMGO struct {
	session *mgo.Session
}

func (w *WidgetMGO) CreateWidget(widget model.Widget) (string, error) {
	defer w.session.Close()
	return "", nil
}

func (w *WidgetMGO) GetWidgets() ([]model.Widget, error) {
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("Widget")
	var widget []model.Widget
	err := widgetCollection.Find(bson.M{}).All(&widget)
	return widget, err
}

func (w *WidgetMGO) GetWidgetByID(widgetID string) (model.Widget, error) {
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("Widget")
	var widget model.Widget
	err := widgetCollection.Find(bson.M{
		"_id": bson.ObjectIdHex(widgetID),
		}).All(&widget)
	return widget, err
}

func (w *WidgetMGO) UpdateWidget(widgetID string) error {
	defer w.session.Close()
	return nil
}

func GetWidgetDB() WidgetMGO {
	return WidgetMGO {
		session: getSession(),
	}
}