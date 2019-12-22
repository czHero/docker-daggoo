#source ~/.bash_profile
source ~/.bashrc
cd /data/PyWorkspace/DGNebula
./project.sh start
daphne -b 0.0.0.0 -p 18001 DgEt.asgi:application
#rm -fr /data/var/log/celery/run/worker.pid
#/etc/init.d/celeryd start
/bin/bash
