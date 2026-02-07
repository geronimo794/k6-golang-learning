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