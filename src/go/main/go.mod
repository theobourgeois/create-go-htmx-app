module main

go 1.21.6

require server v0.0.0

require db v0.0.0

require github.com/go-sql-driver/mysql v1.7.1 // indirect

replace server => ../server

replace db => ../db
