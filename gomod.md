上面这个不推荐用了

推荐设置代理

export GOPROXY=https://athens.azurefd.net

或

export GOPROXY=https://goproxy.io

然后直接使用 go mod tidy

关于 GOPROXY 推荐看下 < https://github.com/gomods/athens > 这个项目，另外官方的 issues 里也有同学在提议官方考虑下大陆被墙的状况，希望能提供更直接的解决方案，不过有待观察

另外顺便推荐下 < https://github.com/hqpko/gomod > ，主要是让 go mod graph 输出类似于 tree 的层级结构，没事写点小工具还是挺有意思的 :bowtie: