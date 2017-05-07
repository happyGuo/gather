package agent

import "bytes"

type Sender struct {
	id                   int
	sBuffer              chan bytes.Buffer
	file_mem_folder_name string
	memBuffer            chan bytes.Buffer //sender自己的chan，用于保证sBuffer不阻塞
	connection           Connection
	status               *int
	sendToAddress        string
}

func NewSender() {

}
