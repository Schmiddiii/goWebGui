
let obj = {
    responsePort: document.URL + "backend/respond",
    getExtras: getDataByClassName,
    responseHandler: handleResponse
}

let views = document.getElementsByClassName("click")

for (i = 0; i < views.length; i++) {
    setUpOnClick(views.item(i))
}

function setUpOnClick(view) {
    view.onclick = () => {
        let xhl = new XMLHttpRequest()
        xhl.open('PUT', obj.responsePort, true)
        xhl.onreadystatechange = () => {
            if (xhl.readyState == 4 && xhl.status == 200) {
                let response = xhl.response
                obj.responseHandler(JSON.parse(response))
            }
        }

        let extraData = obj.getExtras(view.id)

        xhl.send(JSON.stringify({ id: view.id, extras: extraData }))
    }
}

function getDataByClassName(classname) {
    let views = document.getElementsByClassName("data_" + classname)

    let result = []

    for (i = 0; i < views.length; i++) {
        let v = views.item(i)
        if (v.nodeName == "INPUT") {
            result.push(v.value)
        }
    }

    return result
}

function handleResponse(message){
    console.log(message)
    console.log("Id: "+message.id)
}