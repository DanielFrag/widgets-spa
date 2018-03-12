package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/DanielFrag/widgets-spa-rv/model"
	"github.com/DanielFrag/widgets-spa-rv/repository"
	"github.com/DanielFrag/widgets-spa-rv/utils"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func CreateWidget(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body request: "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	widgetRepository, widgetRepositoryError := extractWidgetRepository(r)
	if widgetRepositoryError != nil {
		http.Error(w, widgetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	var widget model.Widget
	jsonError := json.Unmarshal(body, &widget)
	if jsonError != nil {
		http.Error(w, "Json error: "+jsonError.Error(), http.StatusInternalServerError)
		return
	}
	widgetError := widgetRepository.CreateWidget(widget)
	if widgetError != nil {
		http.Error(w, "Error adding the widget: "+widgetError.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("added!"))
	return
}

func GetWidgets(w http.ResponseWriter, r *http.Request) {
	widgetRepository, widgetRepositoryError := extractWidgetRepository(r)
	if widgetRepositoryError != nil {
		http.Error(w, widgetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	widget, widgetError := widgetRepository.GetWidgets()
	if widgetError != nil {
		http.Error(w, "Error: "+widgetError.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(widget))
	return
}

func GetWidgetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, "Error id not provided", http.StatusBadRequest)
		return
	}
	widgetRepository, widgetRepositoryError := extractWidgetRepository(r)
	if widgetRepositoryError != nil {
		http.Error(w, widgetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	widgets, widgetsError := widgetRepository.GetWidgetByID(vars["id"])
	if widgetsError != nil {
		http.Error(w, "Error: "+widgetsError.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(widgets))
	return
}

func ChangeWidget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body request: "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	widgetRepository, widgetRepositoryError := extractWidgetRepository(r)
	if widgetRepositoryError != nil {
		http.Error(w, widgetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	var m map[string]interface{}
	jsonError := json.Unmarshal(body, &m)
	if jsonError != nil {
		http.Error(w, "Json error: "+jsonError.Error(), http.StatusInternalServerError)
		return
	}
	updateError := widgetRepository.UpdateWidget(vars["id"], m)
	if updateError != nil {
		http.Error(w, "widget not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(204)
}

func extractWidgetRepository(r *http.Request) (repository.WidgetRepository, error) {
	contextWidgetRepository := context.Get(r, "WidgetRepository")
	if contextWidgetRepository == nil {
		return nil, errors.New("Can't access the context widget repository")
	}
	widgetRepository, widgetRepositoryOk := contextWidgetRepository.(repository.WidgetRepository)
	if !widgetRepositoryOk {
		return nil, errors.New("Can't access the widget repository")
	}
	return widgetRepository, nil
}
