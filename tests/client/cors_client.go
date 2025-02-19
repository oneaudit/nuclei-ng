package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", corsRequest)
	_ = http.ListenAndServe(":5001", nil)
}

func corsRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>CORS Client</title>
    <script>
        function makeRequest() {
            fetch('http://localhost:5000/cors', {
                method: 'GET',
                credentials: 'include'
            })
            .then(response => response.text())
            .then(data => {
                document.getElementById('response').innerHTML = data;
            })
            .catch(error => {
                document.getElementById('response').innerHTML = 'Error:' + error;
            });
        }
    </script>
</head>
<body>
    <button onclick="makeRequest()">Request CORS Page</button>
    <div id="response"></div>
</body>
</html>
`))
	return
}
