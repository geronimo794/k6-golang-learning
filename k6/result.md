# Without PGBouncer
## 5000 iterations
```
  █ TOTAL RESULTS 

    HTTP
    http_req_duration..............: avg=5.64ms min=3.73ms med=5.25ms max=53.36ms p(90)=6.98ms p(95)=7.91ms
      { expected_response:true }...: avg=5.73ms min=3.73ms med=5.33ms max=53.36ms p(90)=7.12ms p(95)=8.1ms 
    http_req_failed................: 29.79% 1490 out of 5000
    http_reqs......................: 5000   175.172983/s

    EXECUTION
    iteration_duration.............: avg=5.69ms min=3.77ms med=5.31ms max=53.47ms p(90)=7.04ms p(95)=7.99ms
    iterations.....................: 5000   175.172983/s
    vus............................: 1      min=1            max=1
    vus_max........................: 1      min=1            max=1

    NETWORK
    data_received..................: 1.2 MB 43 kB/s
    data_sent......................: 405 kB 14 kB/s
```
## 10000 iterations
```
  █ TOTAL RESULTS 

    HTTP
    http_req_duration..............: avg=5.6ms  min=3.97ms med=5.28ms max=47.4ms  p(90)=6.88ms p(95)=7.57ms
      { expected_response:true }...: avg=6.25ms min=4.37ms med=5.94ms max=33.46ms p(90)=7.27ms p(95)=7.98ms
    http_req_failed................: 64.87% 6487 out of 10000
    http_reqs......................: 10000  176.692148/s

    EXECUTION
    iteration_duration.............: avg=5.65ms min=4.03ms med=5.33ms max=47.45ms p(90)=6.94ms p(95)=7.63ms
    iterations.....................: 10000  176.692148/s
    vus............................: 1      min=1             max=1
    vus_max........................: 1      min=1             max=1

    NETWORK
    data_received..................: 3.5 MB 62 kB/s
    data_sent......................: 810 kB 14 kB/s




running (00m56.6s), 0/1 VUs, 10000 complete and 0 interrupted iterations
```

# With PGBouncer
## 5000 iterations
```
      █ TOTAL RESULTS 

    HTTP
    http_req_duration..............: avg=2.31ms min=1.62ms med=2.2ms  max=34.91ms p(90)=2.51ms p(95)=2.84ms
      { expected_response:true }...: avg=2.31ms min=1.62ms med=2.2ms  max=34.91ms p(90)=2.51ms p(95)=2.84ms
    http_req_failed................: 0.00%  0 out of 5000
    http_reqs......................: 5000   424.210475/s

    EXECUTION
    iteration_duration.............: avg=2.35ms min=1.65ms med=2.23ms max=35.31ms p(90)=2.55ms p(95)=2.9ms 
    iterations.....................: 5000   424.210475/s
    vus............................: 1      min=1         max=1
    vus_max........................: 1      min=1         max=1

    NETWORK
    data_received..................: 774 kB 66 kB/s
    data_sent......................: 405 kB 34 kB/s
```
## 10000 iterations
```
  █ TOTAL RESULTS 

    HTTP
    http_req_duration..............: avg=3.15ms min=1.86ms med=2.58ms max=30.89ms p(90)=5.25ms p(95)=7.3ms 
      { expected_response:true }...: avg=3.15ms min=1.86ms med=2.58ms max=30.89ms p(90)=5.25ms p(95)=7.3ms 
    http_req_failed................: 0.00%  0 out of 10000
    http_reqs......................: 10000  311.229477/s

    EXECUTION
    iteration_duration.............: avg=3.2ms  min=1.88ms med=2.62ms max=31.12ms p(90)=5.35ms p(95)=7.42ms
    iterations.....................: 10000  311.229477/s
    vus............................: 1      min=1          max=1
    vus_max........................: 1      min=1          max=1

    NETWORK
    data_received..................: 1.5 MB 48 kB/s
    data_sent......................: 810 kB 25 kB/s




running (00m32.1s), 0/1 VUs, 10000 complete and 0 interrupted iterations
```