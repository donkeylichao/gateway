#!/bin/sh

case $1 in 
	start)
		nohup ./gateway 2>&1 >> info.log 2>&1 /dev/null &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
		killall gateway
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		killall gateway
		sleep 1
		nohup ./gateway 2>&1 >> info.log 2>&1 /dev/null &
		echo "服务已重启..."
		sleep 1
	;;
	*) 
		echo "$0 {start|stop|restart}"
		exit 4
	;;
esac

