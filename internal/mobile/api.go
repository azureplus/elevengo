package mobile

const (
	apiSpaceGetInfo = "https://proapi.115.com/android/1.0/user/space_info"

	apiFileList   = "https://proapi.115.com/android/2.0/ufile/files"
	apiFileDelete = "https://proapi.115.com/android/1.0/rb/delete"

	apiOfflineSign = "https://proapi.115.com/android/files/offlinesign"

	apiUploadGetKey   = "https://proapi.115.com/android/2.0/user/upload_key"
	apiUploadGetInfo  = "https://uplb.115.com/3.0/getuploadinfo.php"
	apiUploadInit     = "https://uplb.115.com/3.0/initupload.php"
	apiUploadGetToken = "https://uplb.115.com/3.0/gettoken.php"
)

type SpaceInfo struct {
}

type UploadInfo struct {
	Endpoint    string `json:"endpoint"`
	GetTokenUrl string `json:"gettokenurl"`
}