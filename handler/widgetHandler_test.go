package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"github.com/DanielFrag/widgets-spa-rv/utils"
	"github.com/DanielFrag/widgets-spa-rv/model"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func TestWidgets(t *testing.T) {
	widgetDBMock := WidgetDBMock{}
	dbInjector := func (next http.HandlerFunc) http.HandlerFunc {
		return func (w http.ResponseWriter, r *http.Request) {
			context.Set(r, "WidgetRepository", &widgetDBMock)
			next(w, r)
			return
		}
	}
	t.Run("CreateWidget", func(t *testing.T) {
		hfi := utils.HandlerFuncInjector{
			Dependencies: []func (http.HandlerFunc) http.HandlerFunc {
				dbInjector,
			},
			Handler: CreateWidget,
		}
		hfi.InjectDependencies()
		jsonReader := bytes.NewReader(utils.FormatJSON(model.Widget{
			Name: "sunda",
			Color: "blue",
			Price: "200.90",
			Inventory: 9,
			Melts: false,
		}))
		req, reqError := http.NewRequest("POST", "/", jsonReader)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 204 {
			t.Error("Wrong status code")
		}
		widgets, widgetsError := widgetDBMock.GetWidgets()
		if widgetsError != nil {
			t.Error("Widgets error: " + widgetsError.Error())
		}
		if len(widgets) != 1 {
			t.Error("Widgets not inserted")
		}
	})
	t.Run("GetWidgetByID", func(t *testing.T) {
		widgets, _ := widgetDBMock.GetWidgets()
		widgetID := widgets[0].ID.Hex()
		hfi := utils.HandlerFuncInjector{
			Dependencies: []func (http.HandlerFunc) http.HandlerFunc {
				dbInjector,
			},
			Handler: GetWidgetById,
		}
		hfi.InjectDependencies()
		req, reqError := http.NewRequest("GET", "/" + widgetID, nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{id}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode == 400 || result.StatusCode == 500 {
			t.Error("Wrong response")
		}
		var resWidget model.Widget
		jsonError := json.Unmarshal(reqRecorder.Body.Bytes(), &resWidget)
		if jsonError != nil {
			t.Error("Json error: " + jsonError.Error())
		}
		if resWidget.ID.Hex() != widgetID {
			t.Error("Wrong widget")
		}
	})
	t.Run("GetAllWidget", func(t *testing.T) {
		hfi := utils.HandlerFuncInjector{
			Dependencies: []func (http.HandlerFunc) http.HandlerFunc {
				dbInjector,
			},
			Handler: GetWidgets,
		}
		hfi.InjectDependencies()
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		var result []model.Widget
		jsonError := json.Unmarshal(reqRecorder.Body.Bytes(), &result)
		if jsonError != nil {
			t.Error("Json error: " + jsonError.Error())
		}
		widgetDB, _ := widgetDBMock.GetWidgets()
		if len(widgetDB) != len(result) {
			t.Error("Inconsistent data")
		}
	})
	t.Run("GetWidgetByID", func(t *testing.T) {
		widgets, _ := widgetDBMock.GetWidgets()
		widgetID := widgets[0].ID.Hex()
		newColor := widgets[0].Color + "2"
		hfi := utils.HandlerFuncInjector{
			Dependencies: []func (http.HandlerFunc) http.HandlerFunc {
				dbInjector,
			},
			Handler: ChangeWidget,
		}
		hfi.InjectDependencies()
		jsonReader := bytes.NewReader(utils.FormatJSON(model.Widget{
			Color: newColor,
		}))
		req, reqError := http.NewRequest("PUT", "/" + widgetID, jsonReader)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{id}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 204 {
			t.Error("Wrong response")
		}
		widgets, _ = widgetDBMock.GetWidgets()
		if widgets[0].Color != newColor {
			t.Error("Widget not updated")
		}
	})
}