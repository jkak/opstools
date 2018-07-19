## os 



### 1 du

```shell
# go run ../main.go du

# go run main.go du ./
{"name":".","size":20389,"pathnum":4,"filenum":7}

# go run main.go du ~/Downloads/softBakup/
{"name":"~/Downloads/softBakup/","size":"3464.28MB","pathnum":405,"filenum":4339}
```

结果说明：

* name 本次查询的目录或文件名
* size 被查询对象的总大小
* pathnum 被查询对象下包含的目录数
* filenum 被查询对象下包含的文件数