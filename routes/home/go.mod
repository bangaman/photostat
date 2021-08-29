module home

go 1.12

require (
	db v0.0.0-00010101000000-000000000000
	templates v0.0.0
)

replace templates => ../templates

replace db => ../../db
