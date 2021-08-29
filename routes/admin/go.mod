module admin

go 1.16

replace app => ../../app

require (
	app v0.0.0
	db v0.0.0
)

replace db => ../../db
