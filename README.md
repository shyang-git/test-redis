# test-redis

#### redis 실행(docker)
redis와 redis-cli 실행
```
docker network create some-network
docker network ls
docker run --network some-network --name some-redis -d redis
docker run -it --network some-network --rm redis redis-cli -h some-redis
```

#### 테스트 방법
* settings.env에 서버 설정
  - settings-sample.env 참고

* test-lpush.go -k ncurion
  - ncurion key에 1초에 1개씩 데이터를 추가함
  - test-lpush.g -k ncurion

* test-rpop.go -k ncurion
  - ncurion key에서 데이터가 있을 때마다 가져옴
  - 개발안됨

* test-mon.go -k ncurion -e 10 -t 1
  - nucion key monitoring
  - test-mon.go -k ncurion 까지 개발

#### 출력 예제
time(시간) [keys] len(개수) ttl(초) expire(10초 설정)
```
2021-02-26T08:05:39Z [ncurion key] 151 9 true
2021-02-26T08:05:40Z [ncurion key] 152 9 true
2021-02-26T08:05:41Z [ncurion key] 153 9 true
2021-02-26T08:05:42Z [ncurion key] 154 9 true
2021-02-26T08:05:42Z [ncurion key] 154 9 true
```
