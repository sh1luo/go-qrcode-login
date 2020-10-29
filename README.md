# QRCode-login

Go语言实现的扫码登陆的后端版本，基于Gin框架。目前实现了这些接口：

- /reg - POST - 注册（用户名和密码）
- /login - POST - 登陆（用户名和密码）
- /qrcode - GET- 获取二维码
- /l/:qrcode - POST - 通过二维码id和token登陆
- /status - GET - 获取二维码状态（待添加...）

接下来会打算不借助 **大型的Web框架**（指Gin等），只借助路由等小型库从头实现一遍。