all:
	echo package server > html.go
	echo const mainHTML=\` >> html.go
	cat main.html >> html.go
	echo \` >> html.go
	go install
