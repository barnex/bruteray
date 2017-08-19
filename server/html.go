package server
const mainHTML=`
<html>

<head>
	<script>
function refresh(){
	document.getElementById("preview").src = "/render?w=600&h=400&cachebreak=" + Math.random();
}
window.setInterval(refresh, 500)

	</script>
</head>

<body>
	<img id="preview" width=600 height=400></img>
</body>

</html>
`
