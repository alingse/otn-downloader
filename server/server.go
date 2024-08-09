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
<h1>识别</h1>
<h1>识别进度/h1>
<div id="text-process"></div>
<h2>识别到的数据</h2>
<div id="text-result"></div>
<script src="https://unpkg.com/html5-qrcode@2.0.9/dist/html5-qrcode.min.js"></script>
<script src="https://unpkg.com/jquery@3.6.0/dist/jquery.min.js"></script>
<div style="width: 500px" id="reader"></div>
<script>
	var html5QrcodeScanner = new Html5QrcodeScanner(
		"reader", { fps: 10, qrbox: 250 });

	var request = {};
	function onScanSuccess(decodedText, decodedResult) {
		console.log(decodedText, decodedResult);
		let dataArray = decodedText.split(":", 3);
		if (dataArray.length <= 3) {
			return;
		}
		let index = dataArray[0];
		let total = dataArray[1];
		let text = dataArray[2];
		if (request[index]) {
			return;
		}
		request[index] = text;
		let lastText = $('#text-result').text();
		$('#text-result').text(lastText+text);
		$('#text-process).text(index + ' : ' + total);
	}
	html5QrcodeScanner.render(onScanSuccess);
</script>
</body>
</html>`
