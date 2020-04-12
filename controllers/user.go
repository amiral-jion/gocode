package controllers

import (
	"encoding/json"
	"github.com/amiraljion/gocode/models"
	"net/http"
	"regexp"
	"strconv"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w,r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		match := uc.userIdPattern.FindStringSubmatch(r.URL.Path)
		if len(match) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(match[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
		}
}

func (uc userController) getAll(w http.ResponseWriter, r *http.Request)  {
	encodeResponseAsJson(models.GetUsers(), w)
}

func (uc userController) get(id int, w http.ResponseWriter)  {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJson(u, w)
}

func (uc userController) post(w http.ResponseWriter, r *http.Request)  {
	u, err := uc.parsRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(u, w)
}

func (uc userController) put(id int, w http.ResponseWriter, r *http.Request)  {
	u, err := uc.parsRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(u, w)
}

func (uc userController) delete(id int, w http.ResponseWriter)  {
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc userController) parsRequest(r *http.Request)(models.User, error)  {
	dec := json.NewDecoder(r.Body)
	var user models.User

	err := dec.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
