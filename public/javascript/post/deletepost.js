
document.addEventListener('click', DeletePostShow);
     
function DeletePostShow(event){
    if(event.target.id === 'post-delete'){
        var name = event.target.getAttribute('data-name')
        document.getElementById("are-you-sure-delete"+name).style.display="block";

    }else if (event.target.id === "cancel-delete"){
        var name = event.target.getAttribute('data-name')
        document.getElementById("are-you-sure-delete"+name).style.display="none";

    }else if(event.target.id === "proceed-delete"){
       
        var name = event.target.getAttribute('data-delete')

        var xmlhttp = new XMLHttpRequest();
        xmlhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                document.getElementById("vals").innerHTML = "welcome"
            }
        }  
        xmlhttp.open("post","/edit/removestatus", true);
        xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xmlhttp.send("status=" +name);

        document.getElementById("delete-post"+name).style.display="none";
        document.getElementById("are-you-sure-delete"+name).style.display="none";
    }
}