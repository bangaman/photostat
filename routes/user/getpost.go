package user

import (
	"db"
	"fmt"
	"strings"
	"strconv"
)

func Splitter(val string) string {
	filter := []string{}
	brks := strings.Split(val, "@")
	for  _, i := range brks{
		if i != ""{
			filter = append(filter, "<a href='"+i+"' ><img src='"+i+"'/></a>")
		}
	}

	if len(filter) > 0 {
		return strings.Join(filter, "")
	}

	return ""
}



func GetImageLength(val string) string {
	filter := []string{}
	brks := strings.Split(val, "@")
	for  _, i := range brks{
		if i != ""{
			filter = append(filter, "<a href='"+i+"' ><img src='"+i+"'/></a>")
		}
	}

	if len(filter) > 0 {
		if len(filter) > 1 {
			return strconv.Itoa(len(filter))+" items"
		}else{
			return strconv.Itoa(len(filter))+" item"
		}
	}

	return ""
}

func Gettags(a,b,c,aa,bb,cc, statusid string) string {
	mat := []string{}

	if len(a) > 1 {
		fmt.Println(len(aa))
		if len(aa) > 1 {
			mat = append(mat, "<button id='comment-button' data-name='"+statusid+"' data-button='1'>#"+a+" &nbsp; &nbsp;<i class='fas fa-book'></i> <span style='display:none;' id='comment-button-"+statusid+"1'>"+aa+"</span></button>")
		}else{
			mat = append(mat, "<button>#"+a+"</button>")
		}
	}

	if len(b) > 1 {
		if len(bb) > 1 {
			mat = append(mat, "<button id='comment-button' data-name='"+statusid+"' data-button='2'>#"+b+" &nbsp; &nbsp;<i class='fas fa-book'></i> <span style='display:none;' id='comment-button-"+statusid+"2'>"+bb+"</span></button>")
		}else{
			mat = append(mat, "<button>#"+b+"</button>")
		}
	}

	if len(c) > 1 {
		if len(cc) > 1 {
			mat = append(mat, "<button id='comment-button' data-name='"+statusid+"' data-button='3'>#"+c+" &nbsp; &nbsp;<i class='fas fa-book'></i> <span style='display:none;' id='comment-button-"+statusid+"3'>"+cc+"</span></button>")
		}else{
			mat = append(mat, "<button>#"+c+"</button>")
		}
	}

	if len(mat) > 0 {
		return strings.Join(mat, "")
	}

	return ""
}

func GetAllpost(id string) []string {
	gtpostslice := []string{}

	type GetUser struct{
		username string `json:"username"`
		profilepic string `json:"profilepic"`
		statusid string `json:"statusid"`
		image string `json:"image"`	
		tag1 string `json:"tag1"`
		text1 string `json:"text1"`
		tag2 string `json:"tag2"`
		text2 string `json:"text2"`
		tag3 string `json:"tag3"`
		text3 string `json:"text3"`
		statusvalue string `json:"statusvalue"`
	}

	
	dbs, err := db.Conn()
	
	if err != nil {
		panic(err.Error())
	}
	
	//defer the close till after all connections are cloed
	defer dbs.Close()
	
	result, err := dbs.Query("select username, profilepic, statusid, image, tag1, text1, tag2, text2, tag3, text3, statusvalue from allusers JOIN status ON allusers.usersid = status.usersid ORDER BY statusdate DESC")
	
	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {
		var tag  GetUser
	
		err = result.Scan(&tag.username, &tag.profilepic, &tag.statusid, &tag.image, &tag.tag1, &tag.text1, &tag.tag2, &tag.text2, &tag.tag3, &tag.text3, &tag.statusvalue )
		if err != nil {
			panic(err.Error())
		}

		if tag.statusvalue == "post"{
			gtpostslice = append(gtpostslice, "<div class='post-div'><div id='dscrption'><img class='avatar' src='"+tag.profilepic+"' /><span>"+tag.username+"</span></div>")
		}else{
			gtpostslice = append(gtpostslice, "<div class='post-div'><div id='dscrption'><img class='avatar' src='"+tag.profilepic+"' /><span>"+tag.username+"</span> <span id='for-sale-post' style='font-size:12px;margin-left:20px;color:green;'><i class='fa fa-shopping-cart'></i> for sale</span> <span style='font-size:12px;margin-left:20px;'>"+GetImageLength(tag.image)+"</span></div>")
		}
		gtpostslice = append(gtpostslice, "<div class='post-description'><div id='image'>"+Splitter(tag.image)+"</div>")
		gtpostslice = append(gtpostslice, "<div id='tags'><div id='tag-button'>"+Gettags(tag.tag1, tag.tag2, tag.tag3, tag.text1, tag.text2, tag.text3, tag.statusid)+"</div><p id='comment-button-display-"+tag.statusid+"'></p></div>")
		gtpostslice = append(gtpostslice, "	<div id='comment-like'><div id='batram"+tag.statusid+"'><button id='batram' data-name='"+tag.statusid+"' data-value='"+Getlike(tag.statusid, id, "like")+"'><span style='pointer-events:none;'> <i style='color:"+Getlike(tag.statusid, id, "color")+";' class='fas fa-heart'></i> "+Getlike(tag.statusid, id, "count")+"</span></button><button><span><i style='color:#ccc;' class='fas fa-comments'></i> <b>"+CountCommentPerPost(tag.statusid)+"</b></span></button></div></div>")
		gtpostslice = append(gtpostslice, "<div id='show-liked-details' style='margin-left:10px;'><span style='font-size:12px;'>"+Getlike(tag.statusid, id, "countwithname")+"</span></div>")
		gtpostslice = append(gtpostslice, "<div class='all-comments'><div id='all-comments-inputs'> <input type='text' id='input-text"+tag.statusid+"' placeholder='Comment'/><button id='getcommentinput' data-name='"+tag.statusid+"'>Send</button></div>")
		gtpostslice = append(gtpostslice, "<div class='all-comments-details' id='all-comments-details"+tag.statusid+"'>"+GetStatusComment(tag.statusid)+"</div></div>")
		gtpostslice = append(gtpostslice, "</div><div id='space-breaker'></div></div>")
	}

	if len(gtpostslice) > 0{
		return gtpostslice
	}

	return []string{}

}




func GetMainUserPosts(id, name, imge string) []string {
	gtpostslice := []string{}

	type GetUser struct{
		username string `json:"username"`
		profilepic string `json:"profilepic"`
		statusid string `json:"statusid"`
		image string `json:"image"`	
		tag1 string `json:"tag1"`
		text1 string `json:"text1"`
		tag2 string `json:"tag2"`
		text2 string `json:"text2"`
		tag3 string `json:"tag3"`
		text3 string `json:"text3"`
	}
	
	dbs, err := db.Conn()
	
	if err != nil {
		panic(err.Error())
	}
	
	//defer the close till after all connections are cloed
	defer dbs.Close()
	
	result, err := dbs.Query("select username, profilepic, statusid, image, tag1, text1, tag2, text2, tag3, text3 from allusers JOIN status ON allusers.usersid = status.usersid ORDER BY statusdate DESC")
	
	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {
		var tag  GetUser
	
		err = result.Scan(&tag.username, &tag.profilepic, &tag.statusid, &tag.image, &tag.tag1, &tag.text1, &tag.tag2, &tag.text2, &tag.tag3, &tag.text3)
		if err != nil {
			panic(err.Error())
		}

		if tag.username == name{
			gtpostslice = append(gtpostslice, "<div class='are-you-sure' id='are-you-sure-delete"+tag.statusid+"' style='display:none;'><div id='sure'><p>are you sure you want to delete this status</div> <div id='sure-button'><button id='proceed-delete' data-delete='"+tag.statusid+"'><i class='fas fa-trash-alt'></i> Delete</button><button id='cancel-delete' data-name='"+tag.statusid+"'>Cancel</button></div></div>")
			gtpostslice = append(gtpostslice, "<div class='post-div' id='delete-post"+tag.statusid+"'><div id='dscrption'>"+imge+"<span>"+tag.username+"</span> <button id='post-delete' data-name='"+tag.statusid+"' ><i style='pointer-events:none;' class='fas fa-trash-alt'></i></button></div>")
			gtpostslice = append(gtpostslice, "<div class='post-description'><div id='image'>"+Splitter(tag.image)+"</div>")
			gtpostslice = append(gtpostslice, "<div id='tags'>"+Gettags(tag.tag1, tag.tag2, tag.tag3, tag.text1, tag.text2, tag.text3, tag.statusid)+"<p id='comment-button-display-"+tag.statusid+"'></p></div>")
			gtpostslice = append(gtpostslice, "	<div id='comment-like'><div id='batram"+tag.statusid+"'><button id='batram' data-name='"+tag.statusid+"' data-value='"+Getlike(tag.statusid, id, "like")+"'><span style='pointer-events:none;'> <i style='color:"+Getlike(tag.statusid, id, "color")+";' class='fas fa-heart'></i> "+Getlike(tag.statusid, id, "count")+"</span></button><button><span><i style='color:#ccc;' class='fas fa-comments'></i> <b>"+CountCommentPerPost(tag.statusid)+"</b></span></button></div></div>")
			gtpostslice = append(gtpostslice, "<div class='all-comments'><div id='all-comments-inputs'> <input type='text' id='input-text"+tag.statusid+"' placeholder='Comment'/><button id='getcommentinput' data-name='"+tag.statusid+"'>Send</button></div>")
		    gtpostslice = append(gtpostslice, "<div class='all-comments-details' id='all-comments-details"+tag.statusid+"'>"+GetStatusComment(tag.statusid)+"</div></div>")
			gtpostslice = append(gtpostslice, "</div><div id='space-breaker'></div></div>")
		}
	}

	if len(gtpostslice) > 0{
		return gtpostslice
	}else{
		gtpostslice = append(gtpostslice, "<div class='empty-post-showcase'><div id='breaker'></div><p>@"+name+" Share your first <b>STATUS <i class='fas fa-smile'></i></b></p><span id='empty-box'><i class='fas fa-box-open'></i></span>")
		gtpostslice = append(gtpostslice, "<div id='breaker'></div></div>")
		return gtpostslice
	}

	return []string{}

}






func CountMainUserPosts(id string) []string {
	gtpostslice := []string{}

	type GetUser struct{
		usersid string `json:"usersid"`
	}
	
	dbs, err := db.Conn()
	
	if err != nil {
		panic(err.Error())
	}
	
	//defer the close till after all connections are cloed
	defer dbs.Close()
	
	result, err := dbs.Query("select usersid from status")
	
	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {
		var tag  GetUser
	
		err = result.Scan(&tag.usersid )
		if err != nil {
			panic(err.Error())
		}

		if tag.usersid == id{
			gtpostslice = append(gtpostslice, id)
		}
	}

	if len(gtpostslice) > 0{
		return gtpostslice
	}

	return []string{}

}

