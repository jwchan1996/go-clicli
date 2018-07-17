#### 用户
```shell
user/ #所有用户
user/:uname #单一用户信息
user/role/:role #根据权限索引用户
user/:uname/post/ #单一用户下所有文章

user/create/:uinfo #创建单一用户
user/delete/:uid #删除用户
user/update/:uid/:uinfo 
```

#### 文章
```shell
post/ #所有文章
post/:pid #单一文章信息
post/sort/:sname #根据分类索引文章
post/tag/:tname #根据标签索引文章
post/:pid/comment/ #单一文章下所有评论

post/create/:pinfo
post/delete/:pid
post/update/:pid/:pinfo
```

