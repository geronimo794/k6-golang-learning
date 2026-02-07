# k6 Load Test Results: PGBouncer Comparison

## Summary

| Scenario | Iterations | Avg duration | Failed % | Throughput (req/s) |
|----------|------------|--------------|----------|--------------------|
| Without PGBouncer | 5,000 | 7.13 ms | **29.76%** | 139.0 |
| With PGBouncer | 5,000 | 3.25 ms | **0.00%** | **302.3** |
| Without PGBouncer | 10,000 | 6.16 ms | **64.89%** | 160.7 |
| With PGBouncer | 10,000 | 3.44 ms | **0.00%** | 286.6 |

---

## HTTP metrics

| Metric | Without (5k) | With (5k) | Without (10k) | With (10k) |
|--------|--------------|-----------|---------------|------------|
| **http_req_duration** | | | | |
| avg | 7.13 ms | 3.25 ms | 6.16 ms | 3.44 ms |
| min | 4.08 ms | 2.62 ms | 3.88 ms | 2.59 ms |
| med | 7.04 ms | 3.04 ms | 5.62 ms | 3.08 ms |
| max | 43.31 ms | 20.09 ms | 42.17 ms | 24.96 ms |
| p(90) | 8.83 ms | 3.56 ms | 8.16 ms | 4.62 ms |
| p(95) | 10.00 ms | 4.25 ms | 8.94 ms | 5.31 ms |
| **http_req_failed** | 29.76% (1488) | 0.00% (0) | 64.89% (6489) | 0.00% (0) |
| **http_reqs** | 5000 @ 139.0/s | 5000 @ 302.3/s | 10000 @ 160.7/s | 10000 @ 286.6/s |

---

## Execution metrics

| Metric | Without (5k) | With (5k) | Without (10k) | With (10k) |
|--------|--------------|-----------|---------------|------------|
| **iteration_duration** | | | | |
| avg | 7.18 ms | 3.29 ms | 6.21 ms | 3.48 ms |
| min | 4.10 ms | 2.64 ms | 3.91 ms | 2.62 ms |
| med | 7.09 ms | 3.08 ms | 5.66 ms | 3.12 ms |
| max | 43.37 ms | 20.15 ms | 42.22 ms | 25.01 ms |
| p(90) | 8.88 ms | 3.62 ms | 8.21 ms | 4.68 ms |
| p(95) | 10.08 ms | 4.31 ms | 8.99 ms | 5.36 ms |
| **iterations** | 5000 @ 139.0/s | 5000 @ 302.3/s | 10000 @ 160.7/s | 10000 @ 286.6/s |

---

## Network

| Metric | Without (5k) | With (5k) | Without (10k) | With (10k) |
|--------|--------------|-----------|---------------|------------|
| data_received | 1.2 MB (34 kB/s) | 774 kB (47 kB/s) | 3.5 MB (57 kB/s) | 1.5 MB (44 kB/s) |
| data_sent | 405 kB (11 kB/s) | 405 kB (25 kB/s) | 810 kB (13 kB/s) | 810 kB (23 kB/s) |

---

## Head-to-head (5k iterations)

| Metric | Without PGBouncer | With PGBouncer | Change |
|--------|-------------------|----------------|--------|
| Avg request duration | 7.13 ms | 3.25 ms | **−54%** |
| Failed requests | 29.76% | 0.00% | **−100%** |
| Throughput | 139.0/s | 302.3/s | **+117%** |

## Head-to-head (10k iterations)

| Metric | Without PGBouncer | With PGBouncer | Change |
|--------|-------------------|----------------|--------|
| Avg request duration | 6.16 ms | 3.44 ms | **−44%** |
| Failed requests | 64.89% | 0.00% | **−100%** |
| Throughput | 160.7/s | 286.6/s | **+78%** |

---

## Raw output

<details>
<summary>Without PGBouncer — 5000 iterations</summary>

```
  █ TOTAL RESULTS 

    HTTP
    http_req_duration..............: avg=7.13ms min=4.08ms med=7.04ms max=43.31ms p(90)=8.83ms p(95)=10ms   
      { expected_response:true }...: avg=7.81ms min=5.55ms med=7.41ms max=39.07ms p(90)=9.24ms p(95)=10.68ms
    http_req_failed................: 29.76% 1488 out of 5000
    http_reqs......................: 5000   139.004581/s

    EXECUTION
    iteration_duration.............: avg=7.18ms min=4.1ms  med=7.09ms max=43.37ms p(90)=8.88ms p(95)=10.08ms
    iterations.....................: 5000   139.004581/s
    vus............................: 1      min=1            max=1
    vus_max........................: 1      min=1            max=1

    NETWORK
    data_received..................: 1.2 MB 34 kB/s
    data_sent......................: 405 kB 11 kB/s




running (00m36.0s), 0/1 VUs, 5000 complete and 0 interrupted iterations
default ✓ [ 100% ] 1 VUs  00m36.0s/10m0s  5000/5000 shared iters
```

</details>

<details>
<summary>Without PGBouncer — 10000 iterations</summary>

```
  █ TOTAL RESULTS 

    HTTP
    http_req_duration..............: avg=6.16ms min=3.88ms med=5.62ms max=42.17ms p(90)=8.16ms p(95)=8.94ms 
      { expected_response:true }...: avg=7.81ms min=5.74ms med=7.45ms max=42.17ms p(90)=9ms    p(95)=10.21ms
    http_req_failed................: 64.89% 6489 out of 10000
    http_reqs......................: 10000  160.72444/s

    EXECUTION
    iteration_duration.............: avg=6.21ms min=3.91ms med=5.66ms max=42.22ms p(90)=8.21ms p(95)=8.99ms 
    iterations.....................: 10000  160.72444/s
    vus............................: 1      min=1             max=1
    vus_max........................: 1      min=1             max=1

    NETWORK
    data_received..................: 3.5 MB 57 kB/s
    data_sent......................: 810 kB 13 kB/s




running (01m02.2s), 0/1 VUs, 10000 complete and 0 interrupted iterations
default ✓ [ 100% ] 1 VUs  01m02.2s/10m0s  10000/10000 shared iters
```

</details>

<details>
<summary>With PGBouncer — 5000 iterations</summary>

```
  █ TOTAL RESULTS 

    HTTP
    http_req_duration..............: avg=3.25ms min=2.62ms med=3.04ms max=20.09ms p(90)=3.56ms p(95)=4.25ms
      { expected_response:true }...: avg=3.25ms min=2.62ms med=3.04ms max=20.09ms p(90)=3.56ms p(95)=4.25ms
    http_req_failed................: 0.00%  0 out of 5000
    http_reqs......................: 5000   302.316777/s

    EXECUTION
    iteration_duration.............: avg=3.29ms min=2.64ms med=3.08ms max=20.15ms p(90)=3.62ms p(95)=4.31ms
    iterations.....................: 5000   302.316777/s
    vus............................: 1      min=1         max=1
    vus_max........................: 1      min=1         max=1

    NETWORK
    data_received..................: 774 kB 47 kB/s
    data_sent......................: 405 kB 25 kB/s




running (00m16.5s), 0/1 VUs, 5000 complete and 0 interrupted iterations
default ✓ [ 100% ] 1 VUs  00m16.5s/10m0s  5000/5000 shared iters
```

</details>

<details>
<summary>With PGBouncer — 10000 iterations</summary>

```
  █ TOTAL RESULTS 

    HTTP
    http_req_duration..............: avg=3.44ms min=2.59ms med=3.08ms max=24.96ms p(90)=4.62ms p(95)=5.31ms
      { expected_response:true }...: avg=3.44ms min=2.59ms med=3.08ms max=24.96ms p(90)=4.62ms p(95)=5.31ms
    http_req_failed................: 0.00%  0 out of 10000
    http_reqs......................: 10000  286.648595/s

    EXECUTION
    iteration_duration.............: avg=3.48ms min=2.62ms med=3.12ms max=25.01ms p(90)=4.68ms p(95)=5.36ms
    iterations.....................: 10000  286.648595/s
    vus............................: 1      min=1          max=1
    vus_max........................: 1      min=1          max=1

    NETWORK
    data_received..................: 1.5 MB 44 kB/s
    data_sent......................: 810 kB 23 kB/s




running (00m34.9s), 0/1 VUs, 10000 complete and 0 interrupted iterations
default ✓ [ 100% ] 1 VUs  00m34.9s/10m0s  10000/10000 shared iters
```

</details>
