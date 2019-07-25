
goWebGuiObj.responseHandler=(msg)=>{
    if(msg.id==""){
        return
    }
    else if(msg.id=="clear"){
        displayBoard([["","",""],["","",""],["","",""]])
        document.getElementById("winText").innerHTML=""

        return
    }
    console.log(msg)
    let json=JSON.parse(msg.extras[0])
    console.log(json)
    displayBoard(json)
    
    if(msg.id=="win"){
        console.log("WIN")
        document.getElementById("winText").innerHTML="Player "+ msg.extras[1] +" has won"
    }
    else if(msg.id=="tie"){
        console.log("TIE")
        document.getElementById("winText").innerHTML="Its a tie"
    }
}

function displayBoard(brd){
    for(let i=0;i<3;i++){
        for(let j=0;j<3;j++){
            document.getElementById(i+"/"+j).innerHTML=brd[j][i]
        }
    }
}