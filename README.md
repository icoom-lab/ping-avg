### Ping-avg

Ping multiple hosts to determine the average over time

```
$ ping-avg -h

Ping multiple hosts to determine the average over time

Usage:
  ping-avg [flags]

Flags:
  -c, --count int   Specifies the number of echo Request messages be sent. The default is 1. (default -1)
  -h, --help        help for ping-avg
  -v, --verbose     verbose output
```

```
ping-avg 8.8.8.8 1.1.1.1

8.8.8.8 	avg=279.012ms 		min=221.133ms 	max=336.891ms 	loss=0
1.1.1.1 	avg=272.1655ms 		min=182.365ms 	max=361.966ms 	loss=0
```