## Brief :
The point of this repo, is to show the difference between sending multiple concurrent requests using worker pool vs the normal way of sending multiple requests (loop and sequence exec).


## Dependencies:
[Golang](https://medium.com/@patdhlk/how-to-install-go-1-9-1-on-ubuntu-16-04-ee64c073cd79) ,
[Endpoints server](https://github.com/amrHassanAbdallah/endpoints_availability)


##  Up & Running

`go run main.go`

### Under The hood

The difference between normal seq exec and multiple workers (worker pool).
for the following URLs.
```
["http://localhost:8010/api/slow/test/22","http://localhost:8010/api/fast/test/1","http://localhost:8010/api/fast/test/2","http://localhost:8010/api/fast/test/3","http://localhost:8010/api/slow/test/66"]
```

The endpoints that under `fast/` are fast in the reponse meaning there is no delaying in response return, but `slow/` on the other hand delay the response for 6 seconds.


#### Normal (master branch)
```
testing 6.008153452 http://localhost:8010/api/slow/test/22 400
testing 0.002283575 http://localhost:8010/api/fast/test/1 200
testing 0.002024825 http://localhost:8010/api/fast/test/2 402
testing 0.001916271 http://localhost:8010/api/fast/test/3 400
testing 6.002096375 http://localhost:8010/api/slow/test/66 203
12.02s elapsed

```

#### Multiple workers (feat-multiple-worker branch)
```

0.00 elapsed with domain: http://localhost:8010/api/fast/test/2 with response status 505
0.00 elapsed with domain: http://localhost:8010/api/fast/test/1 with response status 200
0.00 elapsed with domain: http://localhost:8010/api/fast/test/3 with response status 500
6.00 elapsed with domain: http://localhost:8010/api/slow/test/22 with response status 203
6.00 elapsed with domain: http://localhost:8010/api/slow/test/66 with response status 505
6.00s elapsed
```