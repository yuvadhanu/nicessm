package handlers

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"nicessm-api-service/constants"
	"nicessm-api-service/response"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

//DocumentUpload : ""
func (h *Handler) DocumentUpload(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")

	fmt.Println("method:", r.Method)
	vars := mux.Vars(r)
	scenario := vars["scenario"]
	log.Println("Document uploaded for ", scenario)
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer file.Close()
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	docStart := h.ConfigReader.GetString(h.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOC)

	fileuri := docStart + scenario + "/" + s + handler.Filename
	responseURI := constants.DEFAULTFILEURL + scenario + "/" + s + handler.Filename
	fmt.Println("fileuri=", fileuri)
	f, err := os.OpenFile(fileuri, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	//w.Write([]byte(fileuri))
	// response.With200mV2(w, fileuri, platform)
	m := make(map[string]interface{})
	m["uri"] = responseURI
	response.With200V2(w, fileuri, m, platform)
}

//DocumentsUpload : ""
func (h *Handler) DocumentsUpload(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	fhs := r.MultipartForm.File["uploadfiles"]
	vars := mux.Vars(r)
	scenario := vars["scenario"]
	fileURIs := []string{}
	for k, fh := range fhs {
		fi, err := fh.Open()
		if err != nil {
			response.With400V2(w, err.Error(), platform)
			return
		}
		// f is one of the files
		defer fi.Close()
		n := 5
		b := make([]byte, n)
		if _, err := rand.Read(b); err != nil {
			panic(err)
		}
		s := fmt.Sprintf("%X", b)
		docStart := h.ConfigReader.GetString(h.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOC)

		fileuri := docStart + scenario + "/" + s + fhs[k].Filename
		fmt.Println("fileuri=", fileuri)
		responseURI := constants.DEFAULTFILEURL + scenario + "/" + s + fhs[k].Filename

		f, err := os.OpenFile(fileuri, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, fi)
		if err != nil {
			response.With400V2(w, "Error while copying "+fhs[k].Filename+" -"+err.Error(), platform)
			return
		}
		fileURIs = append(fileURIs, responseURI)
	}
	m := make(map[string]interface{})
	m["FileURIs"] = fileURIs
	response.With200V2(w, "Success", m, platform)
}

type Base64Doc struct {
	FileName string `json:"fileName"`
	Data     string `json:"data"`
}

//DocumentUploadBase64 : ""
func (h *Handler) DocumentUploadBase64(w http.ResponseWriter, r *http.Request) {
	var base64Doc Base64Doc
	platform := r.URL.Query().Get("platform")
	if err := json.NewDecoder(r.Body).Decode(&base64Doc); err != nil {
		fmt.Println(err)
		response.With400V2(w, err.Error(), platform)
		return
	}

	idx := strings.Index(base64Doc.Data, ";base64,")
	if idx < 0 {
		panic("InvalidImage")
	}
	ImageType := base64Doc.Data[11:idx]
	log.Println(ImageType)
	unbased, err := base64.StdEncoding.DecodeString(base64Doc.Data[idx+8:])
	if err != nil {
		panic("Cannot decode b64")
	}
	re := bytes.NewReader(unbased)
	fmt.Println("method:", r.Method)
	vars := mux.Vars(r)
	scenario := vars["scenario"]

	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	fileuri := "/documents/" + scenario + "/" + s + base64Doc.FileName
	switch ImageType {
	case "png":
		im, err := png.Decode(re)
		if err != nil {
			panic("Bad png")
		}
		f, err := os.OpenFile("."+fileuri, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic("Cannot open file")
		}

		png.Encode(f, im)
		defer f.Close()
	case "jpeg":
		im, err := jpeg.Decode(re)
		if err != nil {
			panic("Bad jpeg")
		}

		f, err := os.OpenFile("."+fileuri, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic("Cannot open file")
		}

		jpeg.Encode(f, im, nil)
		defer f.Close()
	case "default":
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -Invalid Format", platform)
		return

	}

	m := make(map[string]interface{})
	m["uri"] = fileuri
	response.With200V2(w, fileuri, m, platform)
}
