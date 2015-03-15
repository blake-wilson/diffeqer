var HOST = "http://synthetic-verve-88502.appspot.com/"


function updateIters() {
	var iters = document.getElementById("iterations").value;

	getResults(iters.toString());
}

function getResults(numIters) {
	chart = document.getElementById("result");
	xhr = new XMLHttpRequest();
	xhr.onload = function(e) {
		if (xhr.readyState == 4 && xhr.status == 200) {
			info = JSON.parse(xhr.responseText);

			var max = Math.max.apply(Math, info.estimates);
			var min = Math.min.apply(Math, info.estimates);

			var options = {
				title: "Solution for Equation",
				hAxis: {title: "Time", minValue: info.time[0], maxValue: info.time[info.time.length -1]},
				vAxis: {title: "Value", minValue: min, maxValue: max},
				legend: "none"
			};

			draw_chart(info.time, info.estimates, options);
		}
	}
	xhr.open("GET", HOST + "timestep/" + numIters, true);
	xhr.send()
}

function draw_chart(xVals, yVals, options) {
	var arr = [['x', 'y']];
	for (var i = 0; i < xVals.length; i++) {
		arr.push([xVals[i], yVals[i]]);
	}
	var data = google.visualization.arrayToDataTable(arr);

	var chart = new google.visualization.ScatterChart(document.getElementById("result"));

	chart.draw(data, options);
}
