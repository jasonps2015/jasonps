# jasonps

# GO Linux配置
/etc/profile.d/go.sh -  新建一个 GO.sh 的文件 里面放入下内容

export GOPATH=/home/go_cord 
export GOROOT=/usr/local/go 
export GOBIN=$GOPATH/bin 
export PATH=$PATH:$GOBIN:/usr/local/go/bin

立即生效环境变量
source /etc/profile.d/go.sh


#后台 layuiAdmin
layuiAdmin  GO  后台模板整合 2019-06-21 05:09:05

后台地址   URL/layuiadmingo