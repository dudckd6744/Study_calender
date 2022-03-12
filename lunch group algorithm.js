 function randomDigitCharactersSpecialCharacterslength(v){
  var qwe =''

  var 원팀 = ["",];
  var 투팀 = ["",];
  var 쓰리팀 = [""];
  var possible = [...원팀,...투팀,...쓰리팀]
  const toto = v

  for( var i=0; i < possible.length; i++ ){
      qwe += possible[toto[i]]
  }
  
  var str2 = qwe.replace(/(.{12})/g,"$1\n- ")
  return '- '+ str2

}


function randomInt(){
  let randomIndexArray = []
  for (let i=0; i < 19; i++) {
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



let qqq = 4
test=()=>{
  qqq = qqq - 1;
    if(qqq === 0){
      console.log('---- 식사조 ----');
      clearInterval(playTimer)
      console.log(res);
      return true
    }
  console.log(qqq +"초")

}
let playTimer= setInterval(test , 1000)



const res = pipe(
  test,
randomInt,
randomDigitCharactersSpecialCharacterslength
)(0);


