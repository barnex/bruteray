all:
	echo package bruteray > html.go
	echo const mainHTML=\` >> html.go
	cat main.html >> html.go
	echo \` >> html.go
	go install
