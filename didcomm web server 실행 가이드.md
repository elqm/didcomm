# 환경 구성
VSCode 설치 - https://code.visualstudio.com/Download

go 설치 - https://go.dev/doc/install

<br/> 
<br/>

# 실행 순서
##### 1. didcomm 프로젝트 폴더 다운 (didcomm_web / didcomm-demo-main 폴더 포함)
   
##### 2. cmd창에서 didcomm_web 폴더 위치로 이동

![2  cmd 폴더 위치 이동](https://github.com/elqm/didcomm/assets/121847260/3ce12d3b-96de-41d4-9f71-579dbed32625)


##### 3. go run main -port `웹서버 포트 번호` -backurl `백엔드 서버 URL` 입력

``` 
ex)
go run main -port 8080 -backurl 192.168.20.10:8090
```

![3  go run](https://github.com/elqm/didcomm/assets/121847260/7a0c0ce3-9b08-4466-a336-b62d21dc3ec3)


##### 4. 웹 브라우저에 localhost:웹서버 포트번호 입력
``` 
ex)
localhost:8080
```

![4  브라우저 입력](https://github.com/elqm/didcomm/assets/121847260/c23b7277-76c4-43b6-acbb-534dee0f8cc4)


##### 5. 페이지 접속 완료

![5  페이지 띄우기](https://github.com/elqm/didcomm/assets/121847260/faba8a2d-06a2-40be-8f2a-86ea3b9f0c8c)
