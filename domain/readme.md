

## domain

tools about domains.



### 1 https 

check the expire time of given https domains.

```shell
# go run main.go https
check the expire time for https of given domains

Usage:
  domain https [flags]

Flags:
  -d, --day int32   expire days for domain check (default 60)
  -h, --help        help for https

  
# go run main.go https -d 365 www.yiducloud.com.cn baidu.com 163.com www.sina.com www.weibo.com www.amazon.cn
check if domain is expired within [365] days later:
  www.sina.com error:dial tcp 66.102.251.33:443: connect: connection refused

status	begin-with	expired-at	domain-name
warning	Mar 28，2018	Mar 29，2019	www.amazon.cn
warning	Apr 03，2018	Apr 03，2019	baidu.com
warning	Dec 15，2017	Feb 23，2019	163.com

normal	Nov 16，2017	Nov 16，2020	www.yiducloud.com.cn
normal	Sep 06，2017	Sep 07，2019	www.weibo.com

# go run main.go https www.yiducloud.com.cn baidu.com 163.com www.sina.com www.weibo.com www.amazon.cn
check if domain is expired within [60] days later:
  www.sina.com error:dial tcp 66.102.251.33:443: connect: connection refused

status	begin-with	expired-at	domain-name
normal	Mar 28，2018	Mar 29，2019	www.amazon.cn
normal	Nov 16，2017	Nov 16，2020	www.yiducloud.com.cn
normal	Apr 03，2018	Apr 03，2019	baidu.com
normal	Sep 06，2017	Sep 07，2019	www.weibo.com
normal	Dec 15，2017	Feb 23，2019	163.com

```

