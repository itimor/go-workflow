# iris-ticket

## api文档生成
```
npm install apidoc -g
cd backend
apidoc -i controllers/ -o apidoc/
```
`controllers` 为api目录，执行命令后会在 `apidoc` 生成一个页面文件，可以直接访问， 文档参考 [apidocjs](http://apidocjs.com/)。


## 初始化sql
init.sql是运行程序化，需要执行的一个初始化sql，他会导入账号、角色和基础菜单信息。
