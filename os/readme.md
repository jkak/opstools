## os 



### 1 du

```shell
# go run main.go du ./
{"name":".","size":20389,"pathnum":4,"filenum":7}

# go run main.go du ~/Downloads/softBakup/
{"name":"~/Downloads/softBakup/","size":3624903384,"pathnum":405,"filenum":4339}
```

结果说明：

* name 本次查询的目录或文件名
* size 被查询对象的总大小，包括目录及软链本身的占用，单位是Byte
* pathnum 被查询对象下包含的目录数
* filenum 被查询对象下包含的文件数