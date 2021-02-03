
var sayHello = function () {
    alert("Hello users");
}

var saySomething = function (val) {
    alert(val);
}

var getUserEmail = function () {
    var rtn = "";
    var cname = "email";
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            rtn = c.substring(name.length, c.length);
        }
    }
    return rtn;
}

var setUserEmail = function (val) {
    var d = new Date();
    d.setTime(d.getTime() + (365 * 24 * 60 * 60 * 1000));
    var expires = "expires=" + d.toUTCString();
    document.cookie = "email" + "=" + val + ";" + expires + ";path=/";
}

var deleteNoteConfirm = function () {
    var r = confirm("Delete this Note!");
    if (r == true) {
        // txt = "You pressed OK!";
        deleteNote();
    } 
}


var resetPasswordConfirm = function () {
    var r = confirm("This will reset your password!");
    if (r == true) {
        // txt = "You pressed OK!";
        resetPassword();
    } 
}

var saveLocalStorage = function(key, jsonVal){
    localStorage.setItem(key, jsonVal);
}

var getLocalStorage = function(key){
    return  localStorage.getItem(key);
}