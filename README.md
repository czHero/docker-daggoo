nginx.conf:

	需要修改对外服务器ip

start.sh:

	需要修改工程名


docker-compose.yml:

	1、需要配置数据库对外端口

	2、需要配置web服务对外端口

oracle:

	1、docker exec -it dockerId bash
	
	2、su oracle
	
	3、修改oracle编码，http://note.youdao.com/noteshare?id=a9fc8140a8bf4ba3c7dab69a3c155cb2

	4、创建用户，授权


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
