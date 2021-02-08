#!/bin/bash

#counter_str() {
 #   output=0
  #  for items in $1; do 
   #     output=$(($output + $items))
    #    echo $output
    #done
    #echo $output
#}

counter() {
    echo $1 
    strings=$( cat $1 | sed '/^\s*$/d' | wc -l )
    echo $strings
    echo $strings >> filewithstrings
}

touch filewithstrings

rec() {
    if [ -d "$1" ]; then
        ls "$1" | while read name; do
            rec "$1/$name"
        done
    else 
        if [[ -f "$1" ]]; then
            #echo "ITS A FILE"
            if [[ $1 == *.c ]]; then
                #echo "I FOUND A C FILE UUUUUUUU"
                counter $1
            fi
            if [[ $1 == *.h ]]; then 
                #echo "I FOUND A .h FILE YEA"
                counter $1
            fi
        fi
    fi
}

rec $1 "filewithstrings"
awk '{n += $1}; END{print n}' filewithstrings
rm filewithstrings

#/home/sevastyan/* запускаемая папка

#sudo ln -fs /home/sevastyan/.my_script/lab72.sh /usr/local/bin/lab72.sh