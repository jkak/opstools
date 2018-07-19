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



### 2 dusub

dusub: du the sub directory of given path。统计给定目录下的所有子目录/文件。

```shell
# go run main.go dusub

# ls cmd/
du       du.go    dusub    dusub.go root.go

# go run main.go dusub  cmd readme.md -i "du,dusub"
{"name":"du.go","size":4.00KB,"path":"cmd/du.go""pathnum":0,"filenum":1}
{"name":"dusub.go","size":4.00KB,"path":"cmd/dusub.go""pathnum":0,"filenum":1}
{"name":"root.go","size":4.00KB,"path":"cmd/root.go""pathnum":0,"filenum":1}

```

