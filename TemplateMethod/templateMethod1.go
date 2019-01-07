package TemplateMethod

import "fmt"

type DownloadApi interface {
	serverDownload()
	localSave()
	Download()
}
type Template struct {
	DownloadApi
	Uri string
	LocalPath string
}
func (t *Template)serverDownload()  {
	//默认http方法下载
	fmt.Printf("download %s via http\n",t.Uri)
}
func (t *Template)localSave()  {
	fmt.Printf("默认保存在:%s\n",t.LocalPath)
}
func (t *Template)Download()  {
	fmt.Println("prepare downloading...")
	t.DownloadApi.serverDownload()
	t.DownloadApi.localSave()
	fmt.Println("finish download...")
}

type HttpDownloader struct {
	*Template
}
//覆盖默认的serverDownload
func (hd *HttpDownloader)serverDownload()  {
	fmt.Printf("download %s via http\n", hd.Uri)
}
//覆盖默认的localSave
func (hd *HttpDownloader)localSave()  {
	fmt.Printf("用户选择save在:%s\n",hd.LocalPath)
}
//HttpDownloader默认实现
func NewHttpDownload(uri string,localPath string) DownloadApi  {
	//子类覆盖了父类中部分实现,同时继承了其他实现
	hd:=&HttpDownloader{}
	t:=Template{
		hd,
		uri,
		localPath,
	}
	//用组合来实现继承 一定会将父类赋值给子类中的属性
	hd.Template=&t
	return &t
}

type FtpDownloader struct {
	*Template
}
//覆盖默认的serverDownload
func (fd *FtpDownloader)serverDownload()  {
	fmt.Printf("download %s via ftp\n", fd.Uri)
}
//注意没有覆盖默认的localSave

//FtpDownloader 构造函数
func NewFtpDownload(uri string,localPath string) DownloadApi {
	//子类覆盖了父类中部分实现,同时继承了其他实现
	fd:=&FtpDownloader{}
	t:=Template{
		fd,
		uri,
		localPath,
	}
	//用组合来实现继承 一定会将父类赋值给子类中的属性
	fd.Template=&t
	return &t
}
