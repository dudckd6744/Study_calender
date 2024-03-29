#### 👇 [EC2 볼룸 축소](https://youngchang.tistory.com/entry/EC2-volume-%EC%B6%95%EC%86%8C-EBS) 
### 2022-03-01
    - EC2 볼룸 축소 블로깅
    
### 2022-03-02
    - 사내 프로젝트 리펙토링 관련 검토 ( 금일 : 로직, 테이블, 결제사 사이트 훑어보면서 흐름 이해)
        - 결제사 마다 테이블을 따로 관리중이다. 
            - 확장성 문제로 인해 테이블을 하나로 합치고 필수적인 컬럼만 지정하여 관리해준다.
            - 세세한 정보들은 로깅으로 남겨 준다.
            - 이러한 문제를 해결 함으로써 유지보수 , 확장성, 가독성 다 연관되어 좀 더 퀄리티 있는 코드를 짜낼 수 있을것같다.
    
### 2022-03-03
    - 사내 프로젝트 리펙토링 관련 검토 ( 테이블 설계, 매핑 함수 로직 검토)
        - 결제사 마다 return해주는 키가 다르니깐 매핑할때 하나하나 매핑을 해줘야된다.
            - 결제사가 늘어나며 늘어날수록 매핑해주는 코드가 늘어나며 다소 가독성이 떨어질수도있다고 생각이든다.
            - 함수화를해서 삼항연산자로 2개의 매핑코드를 1개로 줄일 수 있는 방법 ( 조금의 개발생산성 향상?)
            - 국내 | 국외 함수를 만들어서 함수내에 메서드형태로 결제사마다 매핑코드를 만들어 가독성을 높이는 방법?
            
### 2022-03-04
    - 사내 자사 솔루션 결제 관련 리펙토링
        - DB 설계 및 리뉴얼
        - 로직 검토 및 수정 작업

### 2022-03-05
    - 사내 자사 솔루션 결제 관련 리펙토링
        - 어제와 동일 작업
        
### 2022-03-06
    - RoYal 프론트 부분 리펙토링을 위한 github All repo 검색 후 코드 분석

### 2022-03-11
    - 사내 팀원 점심 식사 이벤트 그룹 랜덤배치 알고리즘 설계 

### 2022-03-14

    - nestjs mongodb || typeorm Testing 강의 study
    - 사내 hotfix 이슈 || toss payments 연동 이슈로 인한 야근으로 개인공부시간 부족...

### 2022-03-18
    - 사내 프로젝트 toss 페이먼츠 결제사 배포전 이슈 해결을 위한 야근...

### 2022-03-22
    - 사내 프로젝트 toss 페이먼츠 prod 배포후 모니터링 으로 인한 야근..

### 2022-03-23
    - 사내 프로젝트 카카오 페이 긴급 도입으로 인한 야근
        - 결제 준비 API
        - 결제 승인 API
    - ECS 블로깅 || 해시 함수들 속도 블로깅 복습

### 2022-03-24
    - 카카오페이 (3/23~3/24) 연동
        - 결제 준비 , 결제 승인 API 는 내가 구현하고 취소, 환불은 cto님이 구현하였다.
        - 결제 승인 부분에서 고려해야될 부분과 데이터에 저장전 카카오페이에서 주는 response 값을 매핑시켜야되서 시간을 조금 잡아먹었다.
        - 모두 구현 후에도 err 난 부분들에대해서 디버깅하느라 고생했다.
            - 가맹점 코드 (CID) permission denine  Error
                - 결제 준비 , 결제 승인 까지 구현한 상태에서 실결제 테스트를 해볼려고 cid 키를 발급받은 키로 교체하여 테스트 하였는데 permission denine 이라는 err 가 자꾸 떴다. 
                하지만 TEST CID 로하면 또 잘된다.
                이러한 부분에서 상당히 해맸다.
                구글링을 통해 알아본 결과 카카오페이에 네이티브앱 키 와 CID , 어드민 키는 1:1:1 로 매핑되어있다는 글을 보았다.
                그래서 브레이크 앤컴퍼니 다른 애플리케이션 어드민키로 요청해보았더니 잘돼서 다행히 error 해결 하였다.
            - 부분 취소 요청시 결제 가능한 금액과 취소금액이 매치되지않는 Error
                - 결제취소시에는 아무런 이슈가 없었다.
                부분취소시에 결제취소랑 다를바 없이 금액만 다르게해서 요청했지만 not match 에러가 났다.
                구글링 통해 알아본 결과 결제 취소시 부가세 금액도 같이 넣어주지않아 카카오페이 내에서 자동으로 계산을 하는 로직을 돌리는 경우가있는데 이런경우에 정상적으로 취소가 안된다고 나와있어
                취소금액 부가세를 같이 넣어서 요청을 하니깐 부분취소가 정상적으로 작동되었다.
                
#### 👇 [js[자바스크립트] Scope 종류[렉시컬 환경, 실행 컨텍스트, 스코프체인]](https://youngchang.tistory.com/entry/js%EC%9E%90%EB%B0%94%EC%8A%A4%ED%81%AC%EB%A6%BD%ED%8A%B8-Scope-%EC%A2%85%EB%A5%98%EB%A0%89%EC%8B%9C%EC%BB%AC-%ED%99%98%EA%B2%BD-%EC%8B%A4%ED%96%89-%EC%BB%A8%ED%85%8D%EC%8A%A4%ED%8A%B8-%EC%8A%A4%EC%BD%94%ED%94%84%EC%B2%B4%EC%9D%B8), [package-Lock.json 을 굳이 깃에 같이 저장해야 될까??](https://youngchang.tistory.com/entry/package-Lockjson-%EC%9D%84-%EA%B5%B3%EC%9D%B4-%EA%B9%83%EC%97%90-%EA%B0%99%EC%9D%B4-%EC%A0%80%EC%9E%A5%ED%95%B4%EC%95%BC-%EB%90%A0%EA%B9%8C) 
### 2022-03-26
    - js 기초 개념들 블로깅!!
    - package-Lock.json 을 굳이 깃에 같이 저장해야 될까?? 블로깅
    
#### 👇 [package.json [ tilde, caret ],[ major, minor, patch ]](https://youngchang.tistory.com/entry/packagejson-tilde-caret-major-minor-patch) 
### 2022-03-27
    - package.json [ tilde, caret ],[ major, minor, patch ] 블로깅

### 2022-03-28
    - 사내 프로젝트 토스페이먼츠 가상계좌 부분환불 취소 시 err 로 인한 야근
        - 토스 페이먼츠 가상계좌 이외 로직에 데이터 update 쿼리 부분은 transaction 이 걸려있는 상태이다
        - 토스 페이먼츠 유효성 검사 err 날시 payment 데이터를 변경하는데 await을 안걸었을땐 transaction 걸려있는 상태와 무방하므로 상태 업데이트가 잘됐는데 await을 걸어 놓은후 trasaction 걸려있는 상태에서 함수가 리턴이 되어야 commit 되고 await 이 걸린 쿼리가 날아가는데 해당 부분은 유효성 검사후 err가 나면 payment를 업데이트하고 throw를 날리는데 throw 리턴되면 rollback 이된다. 그럼 rollback 이 끝난후 await 이 실행되어야하는데 그땐 이미 함수가 실행이 끝난 상태이다. 해당 문제로 mysql lock 걸리고 timeout 이 떠서 상태 update 가 안되고 mysql err가 났다.
        근데 throw로 err 를 안보내고 return 으로 성공 시켜도 mysql lock timeout err가 나는 로직이었다.
        - 해결
           - 해당 문제에선 payment update 하는 부분에 transaction 을 걸어도 안되는 부분이었고 await을 걸 이유도 딱히 없었다.. 
           - 유효성검사 후 payment update 하는부분에 await 을 없앤후엔 잘 update 되었다.  이제 리뷰할때 신중히 해야겠다..
        - 환불 계좌 유효성 체크 err 날때 알림톡으로 보내는 환불 금액이 0 으로 나타나는 err 수정도 필요하다...

### 2022-03-29
    - 자사 솔루션 해외서비스 준비로인한 야근...
        - 2주 스프린트로 업무를 배정받고 작업을 시작하는데 업무가 많이 몰려있어 야근으로 조금 조정하고있는 중이다.
            - 환불 API 가 따로 없고 다른 API 에서 로직으로 구성되어있었는데 로직을 빼고 환불 API 를 개발
            - 토스페이먼츠 추가 이슈 해결
                - 알림톡 관련 err
                - 가상계좌 부분환불 취소 시 mysql lock timeout err
                - 결제승인 요청시 프론트 amount 백 amount 비교해서 다르면 err 

### 2022-03-30
    - 자사 솔루션 해외서비스
        - 환불 API 재검토 및 코드 정리 , 스웨거 문서화
        - 백오피스에 필요한 사유 작성 API 개발
        - git convention 작성
        
### to do

