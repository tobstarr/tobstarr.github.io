# C3

	<!DOCTYPE html>
	<html>
			<head>
					<meta charset="utf-8">
					<meta name="viewport" content="width=device-width">
					<title>Chart</title>
					<link href="https://cdnjs.cloudflare.com/ajax/libs/c3/0.4.10/c3.min.css" rel="stylesheet" type="text/css">
					<script src="https://cdnjs.cloudflare.com/ajax/libs/d3/3.5.6/d3.min.js" charset="utf-8"></script>
					<script src="https://cdnjs.cloudflare.com/ajax/libs/c3/0.4.10/c3.min.js"></script>
			</head>
			<body>
				<div id="chart"></div>

				<script type="text/javascript" charset="utf-8">
					var chart = c3.generate({
						bindto: "#chart",
						data: {
							x: 'x',
							columns: [
								['x', '2001-01-01', '2001-01-02', '2001-01-03', '2001-01-04'],
								['data1', 1, 3, 2, 4]
							],
							types: {
								data1: "bar"
							}
						},
						axis: {
							x: {
								type: 'timeseries',
								tick: {
									format: "%Y-%m-%d"
								}
							}
						}
					});
				</script>
			</body>
	</html>
