package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func (h *Handler) InitImportsRoutes(api *gin.RouterGroup){
	importsGroup := api.Group("imports")
	{
		importsGroup.Use(h.userIdentity)
		importsGroup.Use(h.authorizeRole(adminRole))
		importsGroup.POST("/import-xmls",h.ImportFiles)
		importsGroup.POST("/import-picture",h.ImportPicture)
	}
}
// @Summary      Import image
// @Description  This endpoint allows you to upload image file
// @Tags         import
// @Accept       multipart/form-data
// @Produce      json
// @Param        image formData file true "Image file (jpg, png, jpg)"
// @Success      200 {string} string 
// @Failure      400 {string} string 
// @Router       /imports/import-picture [post]
func (h *Handler)ImportPicture(c *gin.Context){
	picture, err := c.FormFile("picture")
  if err!= nil {
    newResponse(c, http.StatusBadRequest, "Failed to load picture")
    return
  }
  pictureStream, err := picture.Open()
  if err!= nil {
    newResponse(c, http.StatusInternalServerError, "Failed to read picture")
    return
  }
  defer pictureStream.Close()
	err =	h.services.Imports().ImportPicture(&pictureStream)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Failed to load picture")
    return
	}
  newResponse(c, http.StatusOK, "Picture uploaded successfully")
}


// @Summary      Import products and offers
// @Description  This endpoint allows you to upload two XML files: import.xml for product data and offers.xml for offer data.
// @Tags         import
// @Accept       multipart/form-data
// @Produce      json
// @Param        import.xml formData file true "XML file containing product data"
// @Param        offers.xml formData file true "XML file containing offer data"
// @Success      200 {string} string "Successful import"
// @Failure      400 {string} string "Error in file upload"
// @Router       /imports/import-xmls [post]
func (h *Handler) ImportFiles(c *gin.Context ){
	importXml, err := c.FormFile("importXml")
	
  if err!= nil {
    newResponse(c, http.StatusInternalServerError, "Failed to load importXML")
    return
  }
	importStream, err := importXml.Open()
	if err!= nil {
		newResponse(c, http.StatusInternalServerError, "Failed to read importXML")
    return
	}
	defer importStream.Close()
	offersXml, err := c.FormFile("importXml")
	  if err!= nil {
    newResponse(c, http.StatusInternalServerError, "Failed to load offerstXML")
    return
  }
	offersStream, err := offersXml.Open()
	if err!= nil {
		newResponse(c, http.StatusInternalServerError, "Failed to read offersXML")
    return
	}
	defer offersStream.Close()
	err = h.services.Imports().Parse(&importStream, &offersStream)
	if err!= nil {
		newResponse(c, http.StatusInternalServerError, "Failed to parse files")
		return
	}
}