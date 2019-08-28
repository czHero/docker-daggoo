source ~/.bashrc
rm -fr /data/var/run/celery_task/worker.pid
/etc/init.d/celeryd start
/bin/bash