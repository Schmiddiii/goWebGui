
# goWebGui

goWebGui is a tool for having a web-interface for your Go-programms.

## Usage

Install goWebGui by running `go get github.com/Schmiddiii/goWebGui`.

### Golang

To set up your go-code import `goweb "github.com/Schmiddiii/goWebGui"` to your file and set up the handler recieving the message from Javascript. This is a function with the input and the output of `goweb.Message`. This message consists of an ID from which HTML-element it was send from and some additional data as a string array named Extras. After handling the message in your code you can send a response to your Javascript by returning another `goweb.Message`. In your main-function start goWebGui by using `goweb.SetUpCode("3000", handler)`, where `3000` is your port and `handler` is your message handler you just made.

### HTML and Javascript

Your HTML and Javascript and optionaly your CSS have to be in a folder named `static`. Your folder structure should therefor look similar to this:

```plain
    yourProject
    +-- static
    |   +-- index.html
    |   +-- script.js
    |   +-- style.css
    +-- main.go
```

In your HTML-file you have to reference the `jsWebGui.js` script by using `<script src="https://raw.githack.com/Schmiddiii/goWebGui/master/jsWebGui.js"></script>`. You should import this in front of all your other scripts. This is using [Githack](https://raw.githack.com/) to use the script from this repository. If you want this file to never update and potentialy break your scripts you should copy the script into your static folder and reference this from your HTML. To have a button send messages to your Go-code assign the class `click` and an id to it. For transferring additional data use the class `data_id`, where `id` is the id of your button, to your inputs with type text or checkbox.

In your Javascript you just have to override the `goWebGuiObj.responseHandler` function, which recieves an object with an id and extras. Additionally you can also override the `getExtras` function to manipulate the extras being sent to the Go-code. Finally call `goWebGuiObj.init()` to activate the message-handler.

## More information

To learn more about the usage of this programm see the examples implementing string-concatination or a tick-tack-toe example.
