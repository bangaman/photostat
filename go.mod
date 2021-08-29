module main

go 1.12

require (
	admin v0.0.0
	app v0.0.0
	callback v0.0.0
	edit v0.0.0-00010101000000-000000000000
	github.com/codegangsta/negroni v1.0.0
	github.com/coreos/go-oidc v2.2.1+incompatible // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	home v0.0.0
	login v0.0.0
	logout v0.0.0
	middlewares v0.0.0
	notification v0.0.0-00010101000000-000000000000
	post v0.0.0-00010101000000-000000000000
	user v0.0.0
)

replace app => ./app

replace auth => ./auth

replace callback => ./routes/callback

replace home => ./routes/home

replace login => ./routes/login

replace logout => ./routes/logout

replace middlewares => ./routes/middlewares

replace user => ./routes/user

replace templates => ./routes/templates

replace admin => ./routes/admin

replace db => ./db

replace post => ./routes/post

replace edit => ./routes/edit

replace notification => ./routes/notification
