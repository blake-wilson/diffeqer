var HOST = "http://localhost:8080"

function get_results() {
	para = document.getElementById("result");
	xhr = new XMLHttpRequest();
	xhr.onload = function(e) {
		if (xhr.readyState == 4 && xhr.status == 200) {
			document.getElementById("result").innerHTML = xhr.responseText;
		}
	}
	xhr.open("GET", HOST, true);
	xhr.send()
}
