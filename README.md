go-web-boilerplate
==================

Just a go martini boilerplate inspired by microscopejs code style.

Requirements
------------

	go and GOPATH

Getting started
---------------

Run the following commands:
	
* Install martini and martini-contrib.

		go get github.com/codegangsta/martini
		go get github.com/codegangsta/martini-contrib/...

* Create a new mysql database.
* Edit ./config/database.json according to your database configuration.

* Build & run application

		go build app.go
		go run app.go

* Open your web browser.

>	[http://localhost:3000](http://localhost:3000)

Roadmap
-------

* Database connection with ORM.
* microscopejs project template grunt tasks.
* bower & npm packages (stylus, backbone, requirejs, grunt, bootstrap).