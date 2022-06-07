### 20220-06-04  
  - MSA, 추가 기능 확장을 고려한 DB 테이블 설계
    [box] table 

### 20220-06-07
  - 시퀄라이즈 bulkCreate
    - ignoreDuplicated 
      - pk 중복된 값이 있으면 insert 하지않는 옵션
      - default = false
  - 시퀄라이즈 include insert 이슈
    - 쿼리 한번에 자식 여러개를 insert 하는게 아니라 insert 쿼리를 자식수만큼 요청
  - orm 단에서 unique 를 복수로 설정해도 실제 db 에서 적용이 안돼고 orm 단에서 적용
    - string fk 설정도 마찬가지로 orm 단에서 적용되는 형태


## todo
