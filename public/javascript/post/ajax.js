

document.addEventListener('click', getter);
     
function getter(event){
    if(event.target.id === 'batram'){
        var xmlhttp = new XMLHttpRequest();
        xmlhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                document.getElementById("batram"+event.target.getAttribute('data-name') ).innerHTML = this.responseText;
                // alert(event.target.getAttribute('data-name') )
            }
        }  
        xmlhttp.open("post","/edit/like", true);
        var value = event.target.getAttribute('data-value')
        var name = event.target.getAttribute('data-name') 
        xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xmlhttp.send("value=" + value + "&name="+ name);
    }
}

// document.addEventListener("click", function(event){
//     if(event.target.id === "batram"){
//         $.post("/edit/like",
//         {
//             value: document.getElementById("batram"+event.target.getAttribute('data-name')).getAttribute('data-value'),
//             name:event.target.getAttribute('data-name')
//         },
//         function(data, status){
//             // alert(status)
//             // if(status === "success"){
//                 // alert("ya")
//                 // document.getElementById(event.target.id).setAttribute('data-value', "UNLIKE")
//                 document.getElementById("batram"+event.target.getAttribute('data-name')).innerHTML = data
//                 if(event.target.getAttribute('data-value') == "LIKE"){
//                     document.getElementById(event.target.id).setAttribute('data-value', "UNLIKE")
//                 }else if(event.target.getAttribute('data-value') == "UNLIKE"){
//                     document.getElementById(event.target.id).setAttribute('data-value', "LIKE")
//                 }
//             // }
//         });
//     }
// })

//event.target.getAttribute('data-name')
//event.target.getAttribute('data-sender')

//profile-update-display


