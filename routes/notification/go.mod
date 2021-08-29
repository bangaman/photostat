module notification

go 1.16

replace db => ../../db

replace app => ../../app

replace templates => ../templates

require (
	app v0.0.0-00010101000000-000000000000
	db v0.0.0-00010101000000-000000000000
	templates v0.0.0-00010101000000-000000000000
)
