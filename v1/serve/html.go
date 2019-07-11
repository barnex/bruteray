package serve

const mainHTML = `
<html>

<head>
	<title>bruteray</title>

	<script>

function refresh(){
	document.getElementById("render").src = "/render?cachebreak=" + Math.random();
}
window.setInterval(refresh, 2000)

	</script>
</head>

<body>
	<img id="render"></img>
</body>

</html>
`
