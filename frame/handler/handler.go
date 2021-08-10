package handler

import (
	"encoding/json"
	"net/http"

	"github.com/covrom/hexagonarch/app/userspace"
	"github.com/covrom/hexagonarch/core/space"
	"github.com/covrom/hexagonarch/core/user"
)

type RestAPI struct {
	usp *userspace.UserSpace
}

func NewRestAPI(ust user.UserStorer, spst space.SpaceStorer) *RestAPI {
	users := user.NewUsers(ust)
	spaces := space.NewSpaces(spst)
	return &RestAPI{
		usp: userspace.NewUserSpace(users, spaces),
	}
}

func (rest *RestAPI) CreateSpaceHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var sp userspace.Space
	if err := dec.Decode(&sp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	nsp, err := rest.usp.CreateSpace(sp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(nsp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *RestAPI) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var usr userspace.User
	if err := dec.Decode(&usr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := rest.usp.RegisterUserInSpaces(&usr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
