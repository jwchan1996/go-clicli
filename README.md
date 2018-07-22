# acgzone-server
和谐区的后端重构，语言从 node 换 golang

### 依赖
* httprouter
* fasthttp
* go-sql-driver

### TO DO……
* 争取重构所有 api ，但是无法移除 koa ，koa 将会作为中间层只负责 ssr
* 数据库 mongoDB 换 mysql
* 争取抽象出一个 go web 框架

#### p.s.
这是一次充满勇气的决定，经过这一波，会抽出更多的最佳实践
但是同时增加了我个人的维护成本，我需要同时维护前端（vue、react）、node中间层（koa）、后端（go）

但是这不是我没事找事，我也是经过很久的思想斗争的::>_<::
