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
### to do
