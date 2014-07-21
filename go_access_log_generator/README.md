#What is a big daily logfile?

14,007 Megabyte or 13.68 Gigabyte ~ 2000 page views per second

http://blog.cloudflare.com/2000-page-views-per-second

Log-Structure for estimation:
```
[IP ADRESS ] - - [dd/mmm/yyyy:hh:mm:ss -0h00] "GET /pathtofileorsim HTTP/1.0" [HTTP-Response-Code] [Bytes-Transfered]
199.72.81.55 - - [01/Jul/1995:00:00:01 -0h00] "GET /history/apollo/ HTTP/1.0" 200 6245
```
