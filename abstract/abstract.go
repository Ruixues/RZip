package abstract

const (
	FileTypeDirectory = 0
	FileTypeFile      = iota
)

type FileMessage struct {
	Name string
	Type int
	Size int64 //字节
}

type FileTreeNode struct {
	FileInfo interface{}
	Sons     map[string]*FileTreeNode //根据名称 因为名称是唯一的
	Message  FileMessage
}
type Dealer interface {
	GetNowPath() string                   //获取当前的路径
	GetNowFiles() []*FileTreeNode         //获取当前路径下的文件列表
	GetIntoRelativeDirectory(string) bool //进入目录
	Init(string) error                    //打开特定的文件
}
