package infra

import (
	"github.com/DanielFrag/widgets-spa-rv/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//WidgetMGO wrap the session to access widget data
type WidgetMGO struct {
	session *mgo.Session
}

//CreateWidget insert a new widget in DB
func (w *WidgetMGO) CreateWidget(widget model.Widget) error {
	w.session = getSession()
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("widget")
	return widgetCollection.Insert(widget)
}

//GetWidgets return all widgets
func (w *WidgetMGO) GetWidgets() ([]model.Widget, error) {
	w.session = getSession()
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("widget")
	var widget []model.Widget
	err := widgetCollection.Find(bson.M{}).All(&widget)
	return widget, err
}

//GetWidgetByID return a single widget based on its ID
func (w *WidgetMGO) GetWidgetByID(widgetID string) (model.Widget, error) {
	w.session = getSession()
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("widget")
	var widget model.Widget
	err := widgetCollection.Find(bson.M{
		"_id": bson.ObjectIdHex(widgetID),
	}).One(&widget)
	return widget, err
}

//UpdateWidget set new values for an existed widget
func (w *WidgetMGO) UpdateWidget(widgetID string, widgetData map[string]interface{}) error {
	w.session = getSession()
	defer w.session.Close()
	widgetCollection := w.session.DB(getDbName()).C("widget")
	changes := make(bson.M)
	for key, value := range widgetData {
		changes[key] = value
	}
	return widgetCollection.Update(bson.M{
		"_id": bson.ObjectIdHex(widgetID),
	}, bson.M{
		"$set": changes,
	})
}

//GetWidgetDB return the object to access the widget data
func GetWidgetDB() *WidgetMGO {
	return &WidgetMGO{}
}
