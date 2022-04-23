#### 👇 [TDD 하기전 개념잡기 [1] [ Test Double] 테스트 더블](https://youngchang.tistory.com/entry/TDD-%ED%95%98%EA%B8%B0%EC%A0%84-%EA%B0%9C%EB%85%90%EC%9E%A1%EA%B8%B0-1-Test-Double-%ED%85%8C%EC%8A%A4%ED%8A%B8-%EB%8D%94%EB%B8%94) 
### 2022-04-03
    - TDD 하기전 개념잡기 [1] [ Test Double] 테스트 더블
    
### 2022-04-05

    - nodejs 3 layer arch docs 
    - proxy_pass nginx 관련 문서 검토
        https://jjeongil.tistory.com/1490
    - 견고한 nodejs 설계 훑어보기
        https://velog.io/@hopsprings2/%EA%B2%AC%EA%B3%A0%ED%95%9C-node.js-%ED%94%84%EB%A1%9C%EC%A0%9D%ED%8A%B8-%EC%95%84%ED%82%A4%ED%85%8D%EC%B3%90-%EC%84%A4%EA%B3%84%ED%95%98%EA%B8%B0

### 2022-04-07

    - 자사 솔루션 대만서비스 준비로 인한 야근...
        - 국내는 co.kr  해외는 com  도메인을 쓴다. 하지만 레포는 하나다.
            - 기본 prod yml 파일에서 수정 작업으로 특정 조건을 달아 분기처리하여 대만서비스 || 국내서비스 배포 작업 예정중이다
            - 그에 따른 검토 중
                - 방법
                  1. release create 시 작동이 아니라 publish 작동 되게 한뒤 master merge 되면 
                     release note 가 template 형식에 맞춰 자동으로 생성되게한뒤 수정작업후 publish 하면 배포되게한다.
                - 참조 사이트
                    - https://stackoverflow.com/questions/56798253/release-template-for-github
                    - https://medium.com/pozalabs/github-repository-release-%EB%B0%98%EC%AF%A4-%EC%9E%90%EB%8F%99%ED%99%94%ED%95%98%EA%B8%B0-30e4738e5283


### 2022-04-08

    - 자사 솔루션 github Action yml 수정 및 테스트
        - 문제점
            - 하나의 repo 에서 대만 || 국내 진행중이다. 이 상황에서 release-note를 작성할때 대만 || 국내 둘다 배포된다. 
        - 해결방법
            - release.title 에 특정조건을 걸어 분기처리를 한다.
                - e.g :
                   title:[TW] 1.23.1 => 대만서비스 배포 || [KR] 0.1.1 => 국내서비스 배포
                   - gitAction 에서 제공하는 startsWith(github.event.release.name, '[TW]')을 이용하여 분기처리!
                    - startWith => 앞글자를 
            - release.title 에 휴먼적err 방지를 위해 master에 push 및 merge 시에 release-note 초안을 생성해준다.
            
#### 👇 [Release-note.title 에 따라 배포 분기처리하기 [ GItHub Action ]](https://youngchang.tistory.com/entry/Release-notetitle-%EC%97%90-%EB%94%B0%EB%9D%BC-%EB%B0%B0%ED%8F%AC-%EB%B6%84%EA%B8%B0%EC%B2%98%EB%A6%AC%ED%95%98%EA%B8%B0-GItHub-Action) 
### 2022-04-10
    - Release-note.title 에 따라 배포 분기처리하기 [ GItHub Action ]

### 2022-04-12
    - 자사 솔루션 github Action yml 수정 및 테스트
        - 해외||국내 동시에 배포하기 
        - 국내만 배포하기
        - 해외만 배포하기
        - 를 분기처리해서 진행 중 
        
#### 👇 [Batch vs Scheduler ( 배치와 스케쥴러 차이)](https://youngchang.tistory.com/entry/Batch-vs-Scheduler-%EB%B0%B0%EC%B9%98%EC%99%80-%EC%8A%A4%EC%BC%80%EC%A5%B4%EB%9F%AC-%EC%B0%A8%EC%9D%B4) 
### 2022-04-22
    - Batch vs Scheduler ( 배치와 스케쥴러 차이)

### to do

