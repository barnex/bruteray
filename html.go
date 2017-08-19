package bruteray
const mainHTML=`
<html>

<head>
	<title>bruteray</title>

	<script>

function refresh(){
	document.getElementById("preview").src = "/preview?w=600&h=400&cachebreak=" + Math.random();
	document.getElementById("render").src = "/render?cachebreak=" + Math.random();
}
window.setInterval(refresh, 1000)

	</script>
</head>

<body>
	<img id="preview" width=600 height=400></img>
	<img id="render" width=600 height=400></img>
</body>

</html>
`
