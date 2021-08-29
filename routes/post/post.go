package post

import (
	"net/http"
	"context"
    "github.com/cloudinary/cloudinary-go"
    "github.com/cloudinary/cloudinary-go/api/uploader"
	"app"
	"templates"
	"db"
	"fmt"
	"io/ioutil"
	"strings"
	"html/template"
	"os"
	
)

type ErrorData struct{
	TextErr1 template.HTML
	TextErr2 template.HTML
	TextErr3 template.HTML
	Tag1 string
	Text1 string
	Tag2 string
	Text2 string
	Tag3 string
	Text3 string
}


func Breakall(ma string) string {
	bat := strings.Split(ma, "-")
	fresh := []string{bat[len(bat)-1]}
	return strings.Join(fresh, "")
}

func Makepost(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")
	var ctx = context.Background()
	cld, _ := cloudinary.NewFromParams(os.Getenv("Cloud_name"),os.Getenv("Cloud_api_secret"),os.Getenv("Cloud_sercret"))
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	usersid := session.Values["usersid"].(string)
	title := db.Router(r.URL.Path[len("/post/"):], "/")
	mersh := ErrorData{}
	grabdetails := make(map[string]string)

	if title[0] == "make-post"{
		if strings.ToLower(r.Method) == "post"{
			r.ParseMultipartForm(10 << 15)
			verifyer := []string{}
			cloud := []string{}

			file := r.MultipartForm
			files := file.File["myfile"]

			for i, _ := range files{
				filer, _ := files[i].Open()
				defer filer.Close()
				fmt.Println(filer)
				fmt.Println(files[i].Filename)
				
				tempfile, _ := ioutil.TempFile("", "image-*")
				
				resp, _:= cld.Upload.Upload(ctx, filer, uploader.UploadParams{PublicID: "image"+Breakall(tempfile.Name())});
				
				cloud = append(cloud, "@"+resp.SecureURL)
			}


			mersh.Tag1 =  r.FormValue("tag1")
			mersh.Text1 =  r.FormValue("text1")
			mersh.Tag2 =  r.FormValue("tag2")
			mersh.Text2 =  r.FormValue("text2")
			mersh.Tag3 =  r.FormValue("tag3")
			mersh.Text3 =  r.FormValue("text3")

			if r.FormValue("tag1") == ""{
				if r.FormValue("text1") != ""{
					mersh.TextErr1 = template.HTML("<div id='button-error'><button>#Tags cant be blank</button></div>")
					verifyer = append(verifyer, "error")
				}
			}

			if r.FormValue("tag2") == ""{
				if r.FormValue("text2") != ""{
					mersh.TextErr2 = template.HTML("<div id='button-error'><button>#Tags cant be blank</button></div>")
					verifyer = append(verifyer, "error")
				}
			}

			if r.FormValue("tag3") == ""{
				if r.FormValue("text3") != ""{
					mersh.TextErr3 = template.HTML("<div id='button-error'><button>#Tags cant be blank</button></div>")
					verifyer = append(verifyer, "error")
				}
			}
			
			if len(verifyer) == 0 {
				
				if r.FormValue("tag1") != ""{
					grabdetails["tag1"] = r.FormValue("tag1")
				}else{
					grabdetails["tag1"] = "0"
				}

				if r.FormValue("text1") != ""{
					grabdetails["text1"] = r.FormValue("text1")
				}else{
					grabdetails["text1"] = "0"
				}

				if r.FormValue("tag2") != ""{
					grabdetails["tag2"] = r.FormValue("tag2")
				}else{
					grabdetails["tag2"] = "0"
				}

				if r.FormValue("text2") != ""{
					grabdetails["text2"] = r.FormValue("text2")
				}else{
					grabdetails["text2"] = "0"
				}

				if r.FormValue("tag3") != ""{
					grabdetails["tag3"] = r.FormValue("tag3")
				}else{
					grabdetails["tag3"] = "0"
				}

				if r.FormValue("text3") != ""{
					grabdetails["text3"] = r.FormValue("text3")
				}else{
					grabdetails["text3"] = "0"
				}

				if  r.FormValue("type") == "post"{
					grabdetails["statusvalue"] = "post"
				}else{
					grabdetails["statusvalue"] = "sell"
				}

				fmt.Println(r.FormValue("type"))

				if len(cloud) > 0 {
					grabdetails["image"] = strings.Join(cloud, "")
				}

				grabdetails["usersid"] = usersid

				Add(grabdetails)
				http.Redirect(w, r, "/user/", http.StatusSeeOther)
			}
	
		}

		templates.RenderTemplate(w, []string{"post", "post"}, mersh)
	}
}

