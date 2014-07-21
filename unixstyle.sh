export LC_ALL='C'
cat sample_data/small_NASA_log_Jul95 | \
  awk '{print $7}' | sort | uniq -c | \
  awk '{ x[$2] += $1 } END { for (w in x) { print x[w] " " w } }' | \
  sort -rn | sed 20q

