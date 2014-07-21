# That's not a monolithic system ;)

tr -cs A-Za-z '\n' | \
tr A-Z a-z | \
sort | \
uniq -c | \
sort -rn | \
sed ${1}q

# If you are not a UNIX adept, you may need a little explanation, 
# but not much, to understand this pipeline of processes. 
# The plan is easy:
# 
# 1. Make one-word lines by transliterating the complement (-c) 
#    of the alphabet into newlines (note the quoted newline), 
#    and squeezing out (-s) multiple newlines.
# 2. Transliterate upper case to lower case.
# 3. Sort to bring identical words together.
# 4. Replace each run of duplicate words with a single representative 
#    and include a count (-c).
# 5. Sort in reverse (-r) numeric (-n) order.
# 6. Pass through a stream editor; quit (q) after printing the 
#    number of lines designated by the script’s first parameter (${1}).
