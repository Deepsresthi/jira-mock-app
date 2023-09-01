// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}


type tAuth struct {}
var Auth tAuth


func (_ tAuth) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.Login", args).URL
}

func (_ tAuth) DoLogin(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Auth.DoLogin", args).URL
}

func (_ tAuth) Signup(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.Signup", args).URL
}

func (_ tAuth) DoSignup(
		username string,
		mailID string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "mailID", mailID)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Auth.DoSignup", args).URL
}


type tSprints struct {}
var Sprints tSprints


func (_ tSprints) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sprints.Index", args).URL
}

func (_ tSprints) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sprints.List", args).URL
}

func (_ tSprints) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sprints.New", args).URL
}

func (_ tSprints) Create(
		name string,
		startDate interface{},
		endDate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "startDate", startDate)
	revel.Unbind(args, "endDate", endDate)
	return revel.MainRouter.Reverse("Sprints.Create", args).URL
}

func (_ tSprints) Show(
		id uint,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Sprints.Show", args).URL
}

func (_ tSprints) Edit(
		id uint,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Sprints.Edit", args).URL
}

func (_ tSprints) Update(
		id uint,
		name string,
		startDate interface{},
		endDate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "startDate", startDate)
	revel.Unbind(args, "endDate", endDate)
	return revel.MainRouter.Reverse("Sprints.Update", args).URL
}

func (_ tSprints) Delete(
		id uint,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Sprints.Delete", args).URL
}


