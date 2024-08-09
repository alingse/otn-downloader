package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	log.Println("Starting our simple http server.")

	// Registering our handler functions, and creating paths.
	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	log.Println("Started on port", 8001)
	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server.
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexHtml)
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info page")
}

const indexHtml = `
<!DOCTYPE html>
<html>
<head>
	<title>OTN Downloader Server</title>
</head>
<body>
<h1>访问</h1>
<span>chrome://flags/#unsafely-treat-insecure-origin-as-secure</span>
<script src="https://unpkg.com/html5-qrcode@2.0.9/dist/html5-qrcode.min.js"></script>
<div style="width: 500px" id="reader"></div>
<script>
	function onScanSuccess(decodedText, decodedResult) {
		console.log(decodedText, decodedResult);
	}

	var html5QrcodeScanner = new Html5QrcodeScanner(
		"reader", { fps: 10, qrbox: 250 });
	html5QrcodeScanner.render(onScanSuccess);
</script>
</body>
</html>`
