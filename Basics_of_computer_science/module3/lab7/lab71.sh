#!/bin/bash

first_param=$(date +%s)
$1
echo "poshla zaderzka"
last_param=$(date +%s)
diff=$(($last_param - $first_param))
wait_time=$(($2 - $diff))
i=0
while [ $i -lt 10 ]; do
    $1
    if [ $wait_time -gt 0 ]; then
        echo "poshla zaderzka"
        sleep $wait_time
    fi
done

#sudo ln -fs /home/sevastyan/.my_script/lab71.sh /usr/local/bin/lab71.sh