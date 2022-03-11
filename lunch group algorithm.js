function randomDigitCharactersSpecialCharacterslength(v){
    var text='[';
    var qwe =''

    var possible = ["우영창","김수민","김효선","이규형","기모띠"];
    const toto = v?.replaceAll(",","")
    console.log(toto+4)

    for( var i=0; i < possible.length; i++ ){
        qwe += randomPW(possible[toto[i]], 3)

        if(qwe.length === i*6 ){
           text += qwe += 'w'
        }
    }
    
    console.log(qwe)
    return qwe
}

function randomPW(srt ,lenth) {
    let test ='';

    test+= srt

    return test
}

function randomInt(){
    let randomIndexArray = []
    for (let i=0; i <= 3; i++) {
        console.log("2")
      let randomNum= Math.floor(Math.random() * 20)
      if (randomIndexArray.indexOf(randomNum.toString()) === -1) {
        randomIndexArray.push(randomNum.toString())
      } else {
        i--
      }
    }
    return randomIndexArray.toString()
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