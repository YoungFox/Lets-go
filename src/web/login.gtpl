<html>
<head>
<title></title>
</head>
<body>
<form action="/upload" enctype="multipart/form-data" method="post">
  <input type="file" name="uploadfile" />
	用户名:<input type="text" name="username">
	密码:<input type="password" name="password">
	<input type="submit" value="上传">
  <input type="hidden" name="token" value="{{.}}">
</form>
</body>
</html>