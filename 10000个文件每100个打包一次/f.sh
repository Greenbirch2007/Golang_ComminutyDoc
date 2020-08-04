#!/bin/bash

ls  > /tmp/1.txt
declare -i totle_hang=`ls  |wc -l`
declare -i n_mode=`echo "scale=0; $totle_hang/100" |bc -l`


declare -i i=1
until [ "$i" -gt "$n_mode" ]
do
declare -i min_hang=`echo "$i*100-99" |bc -l`
file_name=`less /tmp/1.txt | awk "NR>=$min_hang" |awk "NR<=100"`
         tar -zcvf $i.tar.gz $file_name

        ((i++))
echo $i
done

