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

dusub使用了ByteSize重新定义float类型。以便在输出时转换为易读的文件大小。



### 3 dusub2

dusub2: du the sub directory of given path。统计给定目录下的所有子目录/文件。

dusub2是对dusub的优化版本。将文件大小的单位独立成一个字段。

```shell
# go run ./main.go dusub2  ~/Downloads/jetbrains-goland-2017.3.2
{"name":"goland-2017.3.dmg","size":"193.69","unit":MB,"pathnum":0,"filenum":1}
{"name":"下载说明.txt","size":"4.00","unit":KB,"pathnum":0,"filenum":1}
{"name":"新云软件.url","size":"4.00","unit":KB,"pathnum":0,"filenum":1}

# go run ../main.go dusub2  ~/Downloads/
{"name":"jq","size":"27.98","unit":MB,"pathnum":113,"filenum":859}
{"name":"jt","size":"2752.03","unit":MB,"pathnum":13,"filenum":69}
{"name":"processOn","size":"1923.14","unit":KB,"pathnum":5,"filenum":28}
{"name":"琅琊榜之风起长林47.mp4","size":"1028.39","unit":MB,"pathnum":0,"filenum":1}
{"name":"琅琊榜之风起长林48.mp4","size":"1065.07","unit":MB,"pathnum":0,"filenum":1}
{"name":"琅琊榜之风起长林49.mp4","size":"491.19","unit":MB,"pathnum":0,"filenum":1}
{"name":"琅琊榜之风起长林50.mp4","size":"490.21","unit":MB,"pathnum":0,"filenum":1}
{"name":"网络.pptx","size":"1912.27","unit":KB,"pathnum":0,"filenum":1}
```

pathnum为0且filenum为1时，表示是文件。