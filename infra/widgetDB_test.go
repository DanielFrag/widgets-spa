package infra

import (
	"testing"
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

func TestWidgetMGO(t *testing.T) {
	var widget model.Widget
	wName, wColor, wPrice, wInventory, wMelts := "sunda", "red", "10.91", int32(12), true
	t.Run("StartDB", func(t *testing.T) {
		startDBError := StartDB()
		if startDBError != nil {
			t.Error("Can't starts the DB")
		}
		ds.dbName = ds.dbName + "_test"
	})
	defer func() {
		mgoSession := getSession()
		dropDatabaseError := mgoSession.DB(getDbName()).DropDatabase()
		if dropDatabaseError != nil {
			panic(dropDatabaseError)
		}
		StopDB()
	}()
	t.Run("CheckEmptyDB", func(t *testing.T) {
		widgetMGO := GetWidgetDB()
		ws, wsError := widgetMGO.GetWidgets()
		if wsError != nil {
			t.Error("Trying get widgets error: " + wsError.Error())
		}
		if len(ws) != 0 {
			t.Error("Testing with no empty DB")
		}
	})
	t.Run("CreateFirstWidget", func(t *testing.T) {
		widgetMGO := GetWidgetDB()
		wError := widgetMGO.CreateWidget(model.Widget{
			Name: wName,
			Color: wColor,
			Price: wPrice,
			Inventory: wInventory,
			Melts: wMelts,
		})
		if wError != nil {
			t.Error("Can't create the first widget register")
		}
	})
	t.Run("RecoverAllWidget1", func(t *testing.T) {
		widgetMGO := GetWidgetDB()
		wResult, wError := widgetMGO.GetWidgets()
		if wError != nil {
			t.Error("Can't retrieve the widget registers")
		}
		if len(wResult) != 1 {
			t.Error("Can't find the inserted widget")
		}
		widget = wResult[0]
		if wResult[0].Color != wColor || wResult[0].Inventory != wInventory || wResult[0].Melts != wMelts || wResult[0].Name != wName || wResult[0].Price != wPrice {
			t.Error("Inconsistent widget data")
		}
	})
	t.Run("CreateSecondWidget", func(t *testing.T) {
		widgetMGO := GetWidgetDB()
		wError := widgetMGO.CreateWidget(model.Widget{
			Name: wName + "2",
			Color: wColor + "2",
			Price: wPrice + "2",
			Inventory: int32(wInventory + 1),
			Melts: !wMelts,
		})
		if wError != nil {
			t.Error("Can't create the second widget register")
		}
	})
	t.Run("RecoverAllWidget2", func(t *testing.T) {
		widgetMGO := GetWidgetDB()
		wResult, wError := widgetMGO.GetWidgets()
		if wError != nil {
			t.Error("Can't retrieve the widget registers")
		}
		if len(wResult) != 2 {
			t.Error("Can't find the inserted widgets")
		}
	})
	t.Run("RecoverSingleWidget", func(t *testing.T) {
		widgetMGO := GetWidgetDB()
		wResult, wResultError := widgetMGO.GetWidgetByID(widget.ID.Hex())
		if wResultError != nil {
			t.Error("Can't recover the first inserted register")
		}
		if wResult.Color != wColor || wResult.Inventory != wInventory || wResult.Melts != wMelts || wResult.Name != wName || wResult.Price != wPrice {
			t.Error("Inconsistent widget data")
		}
	})
	t.Run("SearchUnexistingWidget", func(t *testing.T) {
		widgetMGO := GetWidgetDB()
		wrongID := bson.NewObjectIdWithTime(time.Now())
		_, wResultError := widgetMGO.GetWidgetByID(wrongID.Hex())
		if wResultError == nil {
			t.Error("Can't recover an unexisted register")
		}
	})
	t.Run("UpdateWidget", func(t *testing.T) {
		widgetMGO := GetWidgetDB()
		wUpdateMap := map[string]interface{} {
			"name": "foo",
			"color": "bar",
		}
		updateError := widgetMGO.UpdateWidget(widget.ID.Hex(), wUpdateMap)
		if updateError != nil {
			t.Error("Can't update the selected document")
		}
		wResult, wResultError := widgetMGO.GetWidgetByID(widget.ID.Hex())
		if wResultError != nil {
			t.Error("Can't find the widget document: " + wResultError.Error())
		}
		if wResult.Name != "foo" || wResult.Color != "bar" {
			t.Error("Error retrieving users")
		}
	})
}