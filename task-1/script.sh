#!/bin/sh

dir_marker=".prune-enable"
dump_marker="crash.dump"
log_size="+1M"
log_keep_lines="20000"


## detect marked dirs
dirs=$(find /opt -name $dir_marker -type f | sed -r 's|/[^/]+$||' | sort | uniq)

echo "\n#### Marked directories ####"
echo "$dirs"


## prune dumps
echo "\n#### Pruning the dumps ####"

for dir in $dirs
do
  find $dir -name $dump_marker -delete -type f -exec echo {} "was pruned" \;
done

## cut down logs files 
echo "\n#### Cleaning-up the logs ####"

for dir in $dirs
do
  find $dir -name "*.log" -type f  -size $log_size -exec sh -c 'echo "cleaning-up: {}" && echo "$(tail -n '$log_keep_lines' {})" > "{}" ' \;

done

