module user

go 1.12

require (
	app v0.0.0
	db v0.0.0-00010101000000-000000000000
	templates v0.0.0
)

replace app => ../../app

replace templates => ../templates

replace db => ../../db
