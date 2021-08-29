module login

go 1.12

require (
	app v0.0.0
	auth v0.0.0
	db v0.0.0-00010101000000-000000000000 // indirect
)

replace app => ../../app

replace auth => ../../auth

replace db => ../../db

