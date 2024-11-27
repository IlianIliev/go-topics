package generics

import "net/http"

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
		"data": [
			{
				"id": 1,
				"name": "John Doe",
				"email": "john_doe@example.org"
			},
			{
				"id": 2,
				"name": "Jane Doe",
				"email": "jane_doe@example.org"
			}
		],
		"paging": {
			"total": 2,
			"limit": 100
		}
	}`))
}

func companiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
		"data": [
			{
				"id": 1,
				"name": "Green Tech LLC",
				"address": "123 Main Street"
			},
			{
				"id": 2,
				"name": "Blue Tech LLC",
				"address": "456 Main Street"
			}
		],
		"paging": {
			"total": 2,
			"limit": 100
		}
	}`))
}

func startServer() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/companies", companiesHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
