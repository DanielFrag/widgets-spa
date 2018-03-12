package handler

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

type UserDBMock struct {
	users []model.User
}

func (u *UserDBMock) GetUserByID(userID string) (model.User, error) {
	for i := range u.users {
		if u.users[i].ID.Hex() == userID {
			return u.users[i], nil
		}
	}
	return model.User{}, errors.New("Can't find the requested user")
}

func (u *UserDBMock) GetUsers() ([]model.User, error) {
	return u.users, nil
}

func (u *UserDBMock) GetUserByLogin(login, pass string) (model.User, error) {
	for i := range u.users {
		if u.users[i].Login == login {
			return u.users[i], nil
		}
	}
	return model.User{}, errors.New("Can't find the requested user")
}

func (u *UserDBMock) UpdateUserSession(userID, session string) error {
	for i := range u.users {
		if u.users[i].ID.Hex() == userID {
			u.users[i].Session = session
			return nil
		}
	}
	return errors.New("Can't find the requested user")
}

func (u *UserDBMock) InitializeUserDB() {
	u.users = []model.User{
		model.User{
			ID: bson.NewObjectIdWithTime(time.Now()),
			Login: "sunda",
			Password: "adnus",
			Session: "123",
			Gravatar: "www.anyware.com.br/sunda",
		},
		model.User{
			ID: bson.NewObjectIdWithTime(time.Now().Add(time.Second * 3)),
			Login: "foo",
			Password: "bar",
			Session: "456",
			Gravatar: "www.anyware.com.br/foo",
		},
	}
}

type WidgetDBMock struct {
	widgets []model.Widget
}

func (w *WidgetDBMock) CreateWidget(widget model.Widget) error {
	widget.ID = bson.NewObjectId()
	w.widgets = append(w.widgets, widget)
	return nil
}

func (w *WidgetDBMock) GetWidgets() ([]model.Widget, error) {
	return w.widgets, nil
}

func (w *WidgetDBMock) GetWidgetByID(widgetID string) (model.Widget, error) {
	for i := range w.widgets {
		if w.widgets[i].ID.Hex() == widgetID {
			return w.widgets[i], nil
		}
	}
	return model.Widget{}, errors.New("Can't find the widget")
}

func (w *WidgetDBMock) UpdateWidget(widgetID string, widgetData map[string]interface{}) error {
	for i := range w.widgets {
		if w.widgets[i].ID.Hex() == widgetID {
			for k, v := range widgetData {
				switch k {
				case "name":
					w.widgets[i].Name = v.(string)
				case "color":
					w.widgets[i].Color = v.(string)
				case "price":
					w.widgets[i].Price = v.(string)
				case "invetory":
					w.widgets[i].Inventory = v.(int32)
				case "melts":
					w.widgets[i].Melts = v.(bool)
				}
			}
			return nil
		}
	}
	return nil
}
