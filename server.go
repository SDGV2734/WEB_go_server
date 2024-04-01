// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
// )
// // func main() {
// //     fmt.Printf("Starting server at port 8080\n")
// //     if err := http.ListenAndServe(":8080", nil); err != nil {
// //         log.Fatal(err)
// //     }
// // }
// func main() {

// 	http.HandleFunc("/login",login_)

//     http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
//         fmt.Fprintf(w, "Hello!")
//     })


//     fmt.Printf("Starting server at port 8080\n")
//     if err := http.ListenAndServe(":8080", nil); err != nil {
//         log.Fatal(err)
//     }
// }
// func login_(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "this is my login ")
// }

package main

import (
 "fmt"
 "log"
 "net/http"
)

func main() {
 // Serve static files
 http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

 fmt.Printf("Sever listening at port 8080\n")

 // ? Create a home page handler
 http.HandleFunc("/", homePageHandler)
 // ? Create a login page handler
 http.HandleFunc("/login", loginPageHandler)

 // ? Start the server, if any error, console log it
 if err := http.ListenAndServe(":8080", nil); err != nil {
  log.Fatal(err)
 }
}

// ? Handler for the home route
func homePageHandler(w http.ResponseWriter, r *http.Request) {
 http.ServeFile(w, r, "./static/homepage.html")
}

// ? Handler for the home route
func loginPageHandler(w http.ResponseWriter, r *http.Request) {
 http.ServeFile(w, r, "./static/login.html")
}