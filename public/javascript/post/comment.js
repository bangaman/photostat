
document.addEventListener('click', getter);
     
function getter(event){
    if(event.target.id === 'getcommentinput'){
        var val = document.getElementById("input-text"+event.target.getAttribute('data-name')).value.trim()
        var statusid = event.target.getAttribute('data-name')
        if(val.length > 0 ){
            var xmlhttp = new XMLHttpRequest();
            xmlhttp.onreadystatechange = function() {
                if (this.readyState == 4 && this.status == 200) {
                    document.getElementById("all-comments-details"+event.target.getAttribute('data-name')).innerHTML = this.responseText;
                }
            }  
            xmlhttp.open("post","/edit/comment", true);
            xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
            xmlhttp.send("value=" + val + "&statusid="+ statusid);
            document.getElementById("input-text"+event.target.getAttribute('data-name')).value = ""
        }else{
            alert("Comment box cant be empty")
        }
    }
}


document.addEventListener('click', ShowMessage);
     
function ShowMessage(event){
    if(event.target.id === 'comment-button'){
        var name = event.target.getAttribute('data-name')
        var databutton = event.target.getAttribute('data-button')
        var message = document.getElementById("comment-button-"+name+databutton).innerHTML
        document.getElementById("comment-button-display-"+name).innerHTML = message
    }

}