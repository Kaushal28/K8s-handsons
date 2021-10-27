package main

import(
	"os"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(os.Getenv("MOUNT_DIR"))))
	http.ListenAndServe(":8000", nil)
}
