// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"jinlingonlineapi/restapi/operations"
	"jinlingonlineapi/restapi/operations/activity"
	"jinlingonlineapi/restapi/operations/category"
	"jinlingonlineapi/restapi/operations/member"
	"jinlingonlineapi/restapi/operations/search"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.json

func configureFlags(api *operations.JinlinonlineAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.JinlinonlineAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.ActivityNrActivityDetailHandler = activity.NrActivityDetailHandlerFunc(func(params activity.NrActivityDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation activity.NrActivityDetail has not yet been implemented")
	})
	api.ActivityNrActivityFavHandler = activity.NrActivityFavHandlerFunc(func(params activity.NrActivityFavParams) middleware.Responder {
		return middleware.NotImplemented("operation activity.NrActivityFav has not yet been implemented")
	})
	api.ActivityNrActivityFieldListHandler = activity.NrActivityFieldListHandlerFunc(func(params activity.NrActivityFieldListParams) middleware.Responder {
		return middleware.NotImplemented("operation activity.NrActivityFieldList has not yet been implemented")
	})
	api.ActivityNrActivityJoinHandler = activity.NrActivityJoinHandlerFunc(func(params activity.NrActivityJoinParams) middleware.Responder {
		return middleware.NotImplemented("operation activity.NrActivityJoin has not yet been implemented")
	})
	api.ActivityNrActivityListHandler = activity.NrActivityListHandlerFunc(func(params activity.NrActivityListParams) middleware.Responder {
		return middleware.NotImplemented("operation activity.NrActivityList has not yet been implemented")
	})
	api.ActivityNrActivityListFavHandler = activity.NrActivityListFavHandlerFunc(func(params activity.NrActivityListFavParams) middleware.Responder {
		return middleware.NotImplemented("operation activity.NrActivityListFav has not yet been implemented")
	})
	api.ActivityNrActivityListJoinedHandler = activity.NrActivityListJoinedHandlerFunc(func(params activity.NrActivityListJoinedParams) middleware.Responder {
		return middleware.NotImplemented("operation activity.NrActivityListJoined has not yet been implemented")
	})
	api.ActivityNrActivityListSponsorHandler = activity.NrActivityListSponsorHandlerFunc(func(params activity.NrActivityListSponsorParams) middleware.Responder {
		return middleware.NotImplemented("operation activity.NrActivityListSponsor has not yet been implemented")
	})
	api.ActivityNrActivitySponsorsListHandler = activity.NrActivitySponsorsListHandlerFunc(func(params activity.NrActivitySponsorsListParams) middleware.Responder {
		return middleware.NotImplemented("operation activity.NrActivitySponsorsList has not yet been implemented")
	})
	api.CategoryNrCategoryListHandler = category.NrCategoryListHandlerFunc(func(params category.NrCategoryListParams) middleware.Responder {
		return middleware.NotImplemented("operation category.NrCategoryList has not yet been implemented")
	})
	api.CategoryNrCategorySubListHandler = category.NrCategorySubListHandlerFunc(func(params category.NrCategorySubListParams) middleware.Responder {
		return middleware.NotImplemented("operation category.NrCategorySubList has not yet been implemented")
	})
	api.MemberNrMemberCheckRechargeHandler = member.NrMemberCheckRechargeHandlerFunc(func(params member.NrMemberCheckRechargeParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberCheckRecharge has not yet been implemented")
	})
	api.MemberNrMemberInitHandler = member.NrMemberInitHandlerFunc(func(params member.NrMemberInitParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberInit has not yet been implemented")
	})
	api.MemberNrMemberRegisterSendSmsHandler = member.NrMemberRegisterSendSmsHandlerFunc(func(params member.NrMemberRegisterSendSmsParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberRegisterSendSms has not yet been implemented")
	})
	api.MemberNrOrderSerialNuberHandler = member.NrOrderSerialNuberHandlerFunc(func(params member.NrOrderSerialNuberParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrOrderSerialNuber has not yet been implemented")
	})
	api.SearchNrSearchHandler = search.NrSearchHandlerFunc(func(params search.NrSearchParams) middleware.Responder {
		return middleware.NotImplemented("operation search.NrSearch has not yet been implemented")
	})
	api.MemberFindPassEditPassHandler = member.FindPassEditPassHandlerFunc(func(params member.FindPassEditPassParams) middleware.Responder {
		return middleware.NotImplemented("operation member.FindPassEditPass has not yet been implemented")
	})
	api.MemberFindPassSendSmsHandler = member.FindPassSendSmsHandlerFunc(func(params member.FindPassSendSmsParams) middleware.Responder {
		return middleware.NotImplemented("operation member.FindPassSendSms has not yet been implemented")
	})
	api.MemberLoginHandler = member.LoginHandlerFunc(func(params member.LoginParams) middleware.Responder {
		return middleware.NotImplemented("operation member.Login has not yet been implemented")
	})
	api.MemberOrderListHandler = member.OrderListHandlerFunc(func(params member.OrderListParams) middleware.Responder {
		return middleware.NotImplemented("operation member.OrderList has not yet been implemented")
	})
	api.MemberRegisterHandler = member.RegisterHandlerFunc(func(params member.RegisterParams) middleware.Responder {
		return middleware.NotImplemented("operation member.Register has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}