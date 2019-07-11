all:
	echo package serve > html.go
	echo const mainHTML=\` >> html.go
	cat ui.html >> html.go
	echo \` >> html.go
	go install
