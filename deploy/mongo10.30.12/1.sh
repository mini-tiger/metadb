#!/bin/sh
master() {
 echo `ps aux|grep -c mongo`
}

host_count() {
 echo `netstat -ntlp|grep -c 8080`
}

i=80
while [ "$i" -le 101 ]; do
	m=$( master )
        h=$( host_count )
	if [ "$m" -eq 1 -a "$h" -eq 2 ]
	then
     		echo "directory not created"
                break
	else
     		echo "directory already created"
	fi
        sleep 15
        i=$(( i + 1 ))
done
