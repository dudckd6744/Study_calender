function randomDigitCharactersSpecialCharacterslength(v){
    var qwe =''
    
    var 원팀 = [...]
    var 투팀 = [...]
    var 쓰리팀 = [...]

    var possible = [...원팀,...투팀,...쓰리팀]
    const toto = v
  
    for( var i=0; i < possible.length; i++ ){
        qwe += randomPW(possible[toto[i]])
  
    }
    
    var str2 = qwe.replace(/(.{12})/g,"$1<=식사조=>")

    return str2
  }
  
  function randomPW(srt ,lenth) {
    let test ='';
  
    test+= srt
  
    return test
  }
  
  
  function randomInt(){
    let randomIndexArray = []
    for (let i=0; i < 19; i++) {
        console.log("2")
      let randomNum= Math.floor(Math.random() * 19)
      if (randomIndexArray.indexOf(randomNum.toString()) === -1) {
        randomIndexArray.push(randomNum.toString())
      } else {
        i--
      }
    }
    return randomIndexArray
  }
  
  
  const pipe = (...funcs) => v => {
    return funcs.reduce((res, func) => {
      return func(res);
    }, v);
  };
  
  const res = pipe(
  randomInt,
  randomDigitCharactersSpecialCharacterslength
  )(0);
  
  
  
  console.log(res); // 10. 잘 작동한다 !
