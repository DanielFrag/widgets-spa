package infra

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

type WidgetMGO struct {
	session *mgo.Session
}

func (w *WidgetMGO) CreateWidget(widget model.Widget) error {
	w.session = getSession()
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("Widget")
	return widgetCollection.Insert(widget)
}

func (w *WidgetMGO) GetWidgets() ([]model.Widget, error) {
	w.session = getSession()
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("Widget")
	var widget []model.Widget
	err := widgetCollection.Find(bson.M{}).All(&widget)
	return widget, err
}

func (w *WidgetMGO) GetWidgetByID(widgetID string) (model.Widget, error) {
	w.session = getSession()
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("Widget")
	var widget model.Widget
	err := widgetCollection.Find(bson.M{
		"_id": bson.ObjectIdHex(widgetID),
		}).All(&widget)
	return widget, err
}

func (w *WidgetMGO) UpdateWidget(widgetID string, widgetData map[string]interface{}) error {
	w.session = getSession()
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("Widget")
	changes := make(bson.M)
	for key, value := range widgetData {
		changes[key] = value
	}
	return widgetCollection.Update(bson.M {
		"_id": bson.ObjectIdHex(widgetID),
	}, bson.M {
		"$set": changes,
	})
}

func GetWidgetDB() *WidgetMGO {
	return &WidgetMGO {}
}