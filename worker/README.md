daggoo,

docker run:

	docker run  -v /data:/data -p 8000:8000 --sysctl net.core.somaxconn=4096 -w /data -dit chizhang/daggoo /bin/bash start.sh –D

docker-compose.yml:

	version: "3"
	services:

	    worker:
	        image: chizhang/daggoo
	        # build: ./worker
	        command: /bin/bash start.sh
	        tty: true
	        volumes:
	            - /data:/data
	        sysctls:
	            net.core.somaxconn: 4096


