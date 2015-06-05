var HOST = "http://synthetic-verve-88502.appspot.com/"

function requestUpdate() {
	chart = document.getElementById("result");
	xhr = new XMLHttpRequest();
	xhr.onload = function(e) {
		if (xhr.readyState == 4 && xhr.status == 200) {
			info = JSON.parse(xhr.responseText);

			var max = Math.max.apply(Math, info.estimates);
			var min = Math.min.apply(Math, info.estimates);

			var options = {
				hAxis: {title: "Time", minValue: info.time[0], maxValue: info.time[info.time.length -1],
				       gridlines: {count: 8}},
				vAxis: {title: "Value", minValue: min, maxValue: max, gridlines: {count: 8}},
				legend: "none"
			};

			draw_chart(info.time, info.estimates, options);
		}
	}
	xhr.open("POST", HOST, true);

	var request = {
		timestep: Number(document.getElementById("timestep").value) * 0.001,
		initial_time: Number(document.getElementById("tinit").value),
		initial_value: Number(document.getElementById("yinit").value),
		final_time: Number(document.getElementById("range").value),

		method: document.getElementById("method").value,
		expression: document.getElementById("expression").value
	}

	xhr.send(JSON.stringify(request));
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

requestUpdate()
