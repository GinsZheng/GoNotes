package main

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func uploadImage() {
	http.HandleFunc("/upload/", uploadHandle)                         // 上传图片(重命名) 访问地址：127.0.0.1:8104/upload/
	http.HandleFunc("/uploadOriginalName/", uploadOriginalNameHandle) // 上传图片(原名)
	http.HandleFunc("/uploaded/", showPicHandle)                      // 显示图片
	err := http.ListenAndServe("localhost:8104", nil)
	fmt.Println(err)
}

// 上传图像接口：重命名
func uploadHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	req.ParseForm()
	if req.Method != "POST" {
		w.Write([]byte(html))
		fmt.Println("Get")
	} else {
		// 接收图片
		uploadFile, handle, err := req.FormFile("image")
		errorHandle(err, w)

		// 检查图片后缀
		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			errorHandle(errors.New("只支持jpg/png图片上传"), w)
			return
			//defer os.Exit(2)
		}

		// 保存图片
		os.Mkdir("./uploaded/", 0777)
		uuid1, err := uuid.NewV1()
		uuid1String := uuid1.String() // uuid转为string
		saveFile, err := os.OpenFile("./uploaded/"+uuid1String, os.O_WRONLY|os.O_CREATE, 0666)
		errorHandle(err, w)
		io.Copy(saveFile, uploadFile)

		defer uploadFile.Close()
		defer saveFile.Close()
		// 上传图片成功
		w.Write([]byte("查看上传图片: <a target='_blank' href='/uploaded/" + uuid1String + "'>" + handle.Filename + "</a>"))
	}
}

// 上传图像接口：原图片名
func uploadOriginalNameHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	req.ParseForm()
	if req.Method != "POST" {
		w.Write([]byte(html))
		fmt.Println("Get")
	} else {
		// 接收图片
		uploadFile, handle, err := req.FormFile("image")
		errorHandle(err, w)

		// 检查图片后缀
		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			errorHandle(errors.New("只支持jpg/png图片上传"), w)
			return
			//defer os.Exit(2)
		}

		// 保存图片
		os.Mkdir("./uploaded/", 0777)
		saveFile, err := os.OpenFile("./uploaded/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		errorHandle(err, w)
		io.Copy(saveFile, uploadFile)

		defer uploadFile.Close()
		defer saveFile.Close()
		// 上传图片成功
		w.Write([]byte("查看上传图片: <a target='_blank' href='/uploaded/" + handle.Filename + "'>" + handle.Filename + "</a>"))
	}
}

// 显示图片接口
func showPicHandle(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("." + req.URL.Path)
	errorHandle(err, w)

	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorHandle(err, w)
	w.Write(buff)
}

// 统一错误输出接口(错误可以在H5页面上显示)
func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

// 当方法为Get时，给一个H5页面
const html = `<html> 
  <head></head>
  <body>
      <form method="post" enctype="multipart/form-data">
          <input type="file" name="image" />
          <input type="submit" />
      </form>
  </body>
</html>`

// 参考资料：https://zhuanlan.zhihu.com/p/571936444
