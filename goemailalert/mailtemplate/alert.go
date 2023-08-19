package mailtemplate

const (
	AlertTemplate = `<html>
		<head>
		<title>Threshold Cross Alert</title>
		</head>
		<body>
			<p>This is to inform you that threshold value of the device is crossed. Please take necessary actions.
			</p>
			Threshold Value: %d </br>
			Current Value: %d
		</body>
	</html>`
)
