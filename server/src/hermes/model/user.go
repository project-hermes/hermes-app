package model

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/appengine/log"
	"strings"
)

type User struct {
	auth.UserInfo
}

type UserPublic struct {
	auth.UserInfo
}

type UserInterface interface {
	Get(uid string) (User, error)
	List() ([]User, error)
	Public(user User) UserPublic
	Private(user UserPublic) User
	CreateUpdateFromFirebase(uid string) (User, error)
}

type UserImplementation struct {
	appCtx      context.Context
	firebaseApp *firebase.App
}

func NewUserImplementation(appCtx context.Context, firebaseApp *firebase.App) UserInterface {
	return UserImplementation{appCtx: appCtx, firebaseApp: firebaseApp}
}

func loadUserInfoFromFirebase(appCtx context.Context, firebaseApp *firebase.App, uid string) (auth.UserInfo, error) {
	if authClient, err := firebaseApp.Auth(appCtx); err != nil {
		log.Errorf(appCtx, err.Error())
		return auth.UserInfo{}, err
	} else if user, err := authClient.GetUser(appCtx, uid); err != nil {
		log.Errorf(appCtx, err.Error())
		return auth.UserInfo{}, err
	} else {
		return *user.UserInfo, nil
	}
}

func (u UserImplementation) Get(uid string) (User, error) {
	if db, err := u.firebaseApp.Firestore(u.appCtx); err != nil {
		log.Errorf(u.appCtx, err.Error())
		return User{}, err
	} else if doc, err := db.Collection("User").Doc(uid).Get(u.appCtx); err != nil && !doc.Exists() {
		if firebaseUserInfo, err := loadUserInfoFromFirebase(u.appCtx, u.firebaseApp, uid); err != nil {
			return User{}, err
		} else {
			return User{firebaseUserInfo}, nil
		}
	} else {
		var tempUser User
		if err := doc.DataTo(&tempUser); err != nil {
			log.Errorf(u.appCtx, err.Error())
			return User{}, err
		}
		return tempUser, nil
	}
}

func (u UserImplementation) List() ([]User, error) {
	return nil, nil
}

func (u UserImplementation) Public(user User) UserPublic {
	return UserPublic{user.UserInfo}
}

func (u UserImplementation) Private(user UserPublic) User {
	return User{user.UserInfo}
}

func (u UserImplementation) CreateUpdateFromFirebase(uid string) (User, error) {
	var tempUser User
	db, err := u.firebaseApp.Firestore(u.appCtx)
	defer db.Close()
	if err != nil {
		log.Errorf(u.appCtx, err.Error())
		return User{}, err
	}
	docRef := db.Collection("User").Doc(uid)
	firebaseUserInfo, err := loadUserInfoFromFirebase(u.appCtx, u.firebaseApp, uid)
	if err != nil {
		return User{}, err
	}

	if doc, err := docRef.Get(u.appCtx); err != nil && !strings.Contains(err.Error(), "rpc error: code = NotFound") {
		log.Errorf(u.appCtx, err.Error())
		return User{}, err
	} else if err == nil {
		if err := doc.DataTo(&tempUser); err != nil {
			log.Errorf(u.appCtx, err.Error())
			return User{}, err
		}
	}

	if tempUser.UserInfo != firebaseUserInfo {
		tempUser.UserInfo = firebaseUserInfo
		if _, err := docRef.Set(u.appCtx, tempUser); err != nil {
			log.Errorf(u.appCtx, err.Error())
			return User{}, err
		}
	}

	return tempUser, nil
}
