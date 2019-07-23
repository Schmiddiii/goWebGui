goWebGuiObj.responseHandler=(msg)=>{
    console.log("Modified Response")
    console.log(msg)
    document.getElementById("result").innerHTML=msg.extras[0]
}