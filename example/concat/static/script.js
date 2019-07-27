goWebGuiObj.responseHandler=(msg)=>{
    console.log(msg)
    document.getElementById("result").innerHTML=msg.extras[0]
}

goWebGuiObj.init()