# detailcron
cron configuration details


To run : go run detailCronTab.go <input argument>
Ex : go run detailCronTab.go "*/15 0 */2 1-10 1,5 /usr/bin/find"
O/p : 
minute   0 15 30 45 
hour     0
day of month     1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31
month    1 2 3 4 5 6 7 8 9 10
day of week      1 5
command  /usr/bin/find

Following Expressions are valid : 

- "*"	any value
- ,	value list separator
- -	range of values
- /	step values
