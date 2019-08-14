package main

import "net/http"

func main() {



	//mux := http.NewServeMux()
	//
	//mux.Handle("/group1/M00/", http.StripPrefix("/group1/M00/", http.FileServer(http.Dir("storage_data/data"))))
	//
	//err := http.ListenAndServe(":9999", mux)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	http.Handle("/group1/M00/",http.FileServer(http.Dir("storage_data/data")))
	http.ListenAndServe(":9999", nil)





}
