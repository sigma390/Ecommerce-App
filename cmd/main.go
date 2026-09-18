package main

import "log"

func main() {

	cgf := config{
		addr: ":8080",
		db: dbconfig{
			dsn: "",
		},
	}

	api := application{
		config: cgf,
	}

	if err := api.run(); err != nil {
		log.Fatal(err)
	}

}
