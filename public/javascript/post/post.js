
function checkExt(file, prev){
    if (/\.(jpeg|jpg|png|gif)$/i.test(file.name)){
        var reader = new FileReader();
        reader.addEventListener("load", function () {
          var image = new Image();
          image.title = file.name;
          image.id = "themage"
          image.src = this.result;
          prev.appendChild( image );
        }, false);
        reader.readAsDataURL(file);
    }
}

var m = {}
m[4] = {place:"kaduna"}
console.log(m)
document.addEventListener("click", function(event){
    if(event.target.id === "image-label"){
        document.getElementById(event.target.id).onchange = function(){
            var gtval = document.querySelector("#files")
            var gtfile = document.querySelector('input[type=file]').files;

            var cutter = {}

            
            

            if(gtfile.length > 3){
                cutter[0] = gtfile[0]
                cutter[1] = gtfile[1]
                cutter[2] = gtfile[2]  
                console.log("na true") 
            }

            console.log(cutter[0])

            if(cutter[0] != undefined){
                console.log("process")
                for(var getall in cutter){
                    checkExt(cutter[getall], gtval)  
                }
            }else{
                for(var getall in gtfile){
                    checkExt(gtfile[getall], gtval)  
                }
            }
       }
    }
})

// import {jax} from "../bundle/ajax.js"

// function checkExt(file, prev){
//     if (/\.(jpeg|jpg|png|gif)$/i.test(file.name)){
//         var reader = new FileReader();
//         reader.addEventListener("load", function () {
//           var image = new Image();
//           image.title = file.name;
//           image.id = "themage"
//           image.src = this.result;
//           prev.appendChild( image );
//           document.getElementById("file").style.display="block"
//         }, false);
//         reader.readAsDataURL(file);
//     }
// }


// jax.event({
//     event:"click",
//     run: function (event){
//         if(event.target.id === "userpostreview"){
//            document.getElementById(event.target.id).onchange = function(){
//             var gtval = document.querySelector("#files")
//             var gtfile = document.querySelector('input[type=file]').files;
//             for(var getall in gtfile){
//                 checkExt(gtfile[getall], gtval)
//             }
//            }
//         }
//     }
// })


// jax.event({
//     event:"click",
//     run: function(event){
//         if(event.target.id === "reviewimage"){
//             document.getElementById(event.target.id).onchange = function(){
//              var gtval = document.querySelector("#files")
//              var gtfile = document.querySelector('input[type=file]').files[0];
//              checkExt(gtfile, gtval)
//             }
//         }
//     }
// })



// function show_message()
// {
//     try{
//         if(document.getElementById("upload").innerHTML.length > 5){
//             setTimeout(() => {
//                 document.getElementById("upload").style.display="none";
//             }, 3000);
//         }
//     }catch(err){
//         console.log("something went wrong")
//     }
// }
// show_message()

