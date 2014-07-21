sort|uniq -c|awk '{ x[$2] += $1 } END { for (w in x) { print x[w] " " w } }'|sort -rn|sed 20q
