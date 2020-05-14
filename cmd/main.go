package main

func main() {
	app.InfoLog.Println("Listening on tcp://localhost:8080")
	err := server.ListenAndServe()
	app.ErrorLog.Fatalln(err)
}
