package unpack

// FileMagicNumbers 文件类型及其对应的魔数（文件头标识）
var FileMagicNumbers = map[string][]byte{
	// 图像文件
	".png":  {0x89, 'P', 'N', 'G', '\r', '\n', 0x1A, '\n'},
	".jpg":  {0xFF, 0xD8, 0xFF},
	".jpeg": {0xFF, 0xD8, 0xFF},
	".gif":  {'G', 'I', 'F', '8'},
	".bmp":  {'B', 'M'},
	".webp": {'R', 'I', 'F', 'F'},
	".tiff": {0x49, 0x49, 0x2A, 0x00},
	".cr2":  {0x49, 0x49, 0x2A, 0x00, 0x10, 0x00, 0x00, 0x00, 0x43, 0x52},
	".ico":  {0x00, 0x00, 0x01, 0x00},
	".heic": {'f', 't', 'y', 'p', 'h', 'e', 'i', 'c'},

	// 文档文件
	".pdf":  {'%', 'P', 'D', 'F', '-'},
	".doc":  {0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1},
	".docx": {'P', 'K', 0x03, 0x04},
	".xls":  {0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1},
	".xlsx": {'P', 'K', 0x03, 0x04},
	".ppt":  {0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1},
	".pptx": {'P', 'K', 0x03, 0x04},
	".odt":  {'P', 'K', 0x03, 0x04},
	".ods":  {'P', 'K', 0x03, 0x04},
	".odp":  {'P', 'K', 0x03, 0x04},
	".rtf":  {'{', '\\', 'r', 't', 'f', '1'},
	".epub": {'P', 'K', 0x03, 0x04},
	".mobi": {'M', 'O', 'B', 'I'},

	// 压缩文件
	".zip": {'P', 'K', 0x03, 0x04},
	".rar": {'R', 'a', 'r', '!', 0x1A, 0x07, 0x00},
	".7z":  {'7', 'z', 0xBC, 0xAF, 0x27, 0x1C},
	".gz":  {0x1F, 0x8B, 0x08},
	".tar": {'u', 's', 't', 'a', 'r'},
	".bz2": {'B', 'Z', 'h'},
	".xz":  {0xFD, '7', 'z', 'X', 'Z', 0x00},

	// 音频文件
	".mp3":  {0xFF, 0xFB},
	".wav":  {'R', 'I', 'F', 'F'},
	".ogg":  {'O', 'g', 'g', 'S'},
	".flac": {'f', 'L', 'a', 'C'},
	".aac":  {0xFF, 0xF1},
	".m4a":  {'f', 't', 'y', 'p', 'M', '4', 'A', ' '},
	".mid":  {'M', 'T', 'h', 'd'},
	".aiff": {'F', 'O', 'R', 'M'},

	// 视频文件
	".mp4":  {0x00, 0x00, 0x00, 0x18, 'f', 't', 'y', 'p'},
	".avi":  {'R', 'I', 'F', 'F'},
	".mkv":  {0x1A, 0x45, 0xDF, 0xA3},
	".flv":  {'F', 'L', 'V'},
	".mov":  {'f', 't', 'y', 'p', 'q', 't', ' ', ' '},
	".webm": {0x1A, 0x45, 0xDF, 0xA3},
	".mpg":  {0x00, 0x00, 0x01, 0xBA},
	".wmv":  {0x30, 0x26, 0xB2, 0x75, 0x8E, 0x66, 0xCF, 0x11},

	// 可执行文件和库文件
	".exe": {'M', 'Z'},
	".elf": {0x7F, 'E', 'L', 'F'},
	".so":  {0x7F, 'E', 'L', 'F'},
	".dll": {'M', 'Z'},
	".app": {0xCA, 0xFE, 0xBA, 0xBE}, // Mach-O 文件格式 (macOS)

	// 脚本和标记语言文件
	".json": {'{'}, // 假设 JSON 文件以 { 开头
	".xml":  {'<', '?', 'x', 'm', 'l'},
	".html": {'<', '!', 'D', 'O', 'C', 'T', 'Y', 'P', 'E'},

	// 字体文件
	".ttf":   {0x00, 0x01, 0x00, 0x00, 0x00},
	".otf":   {'O', 'T', 'T', 'O'},
	".woff":  {'w', 'O', 'F', 'F'},
	".woff2": {'w', 'O', 'F', '2'},

	// 数据库文件
	".sqlite": {'S', 'Q', 'L', 'i', 't', 'e', ' ', 'f', 'o', 'r', 'm', 'a', 't', ' ', '3', 0x00},
	".db":     {0x00, 0x06, 0x15, 0x61, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00},

	// 其他
	".swf":     {'F', 'W', 'S'},
	".class":   {0xCA, 0xFE, 0xBA, 0xBE},
	".psd":     {'8', 'B', 'P', 'S'},
	".torrent": {'d', '8', ':', 'a', 'n', 'n', 'o', 'u', 'n', 'c', 'e'},
	".blend":   {'B', 'L', 'E', 'N', 'D', 'E', 'R'},
	".pcap":    {0xD4, 0xC3, 0xB2, 0xA1},
	".dwg":     {0x41, 0x43, 0x31, 0x30},
	".iso":     {0x43, 0x44, 0x30, 0x30, 0x31},
	".vsd":     {0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1},
	".mdb":     {0x00, 0x01, 0x00, 0x00, 0x53, 0x74, 0x61, 0x6E, 0x64, 0x61, 0x72, 0x64, 0x20, 0x4A, 0x65, 0x74},
}
