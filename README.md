nginx.conf:
	需要修改对外服务器ip
docker-compose.yml:
	1、需要配置数据库对外端口
	2、需要配置web服务对外端口
oracle:
	1、修改oracle编码
	4、创建用户


/data目录情况:
├── oracle
│   └── data
├── PyWorkspace
│   ├── DGNebula
│   ├── DShare
│   └── mysite
├── start.sh
├── var
│   └── log
│   	├── nginx
├── www
│   └── static



配置完之后，执行
启动容器：
docker-compose up -d
关闭容器：
docker-compose down
重启容器：
docker-compose restart
