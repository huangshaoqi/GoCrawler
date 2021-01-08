package parser

type MovieInfo struct {
	Title                string `json:"title"`                  // 标题
	Cover                string `json:"cover"`                  // 封面
	TranslatedName       string `json:"translated_name"`        // 译名
	MovieName            string `json:"movie_name"`             // 片名
	Year                 string `json:"year"`                   // 年代
	Origin               string `json:"origin"`                 // 产地
	Type                 string `json:"type"`                   // 类别
	Language             string `json:"language"`               // 语言
	Subtitle             string `json:"subtitle"`               // 字幕
	ReleaseDate          string `json:"release_date"`           // 上映时间
	ImdbScore            string `json:"imdb_score"`             // IMDB评分
	DoubanScore          string `json:"douban_score"`           // 豆瓣评分
	FileFormat           string `json:"file_format"`            // 文件格式
	VideoSize            string `json:"video_size"`             // 视频尺寸
	FilmLength           string `json:"film_length"`            // 片长
	Director             string `json:"director"`               // 导演
	Screenwriter         string `json:"screenwriter"`           // 编剧
	ToStar               string `json:"to_star"`                // 主演
	Label                string `json:"label"`                  // 标签
	Introduction         string `json:"introduction"`           //简介
	Awards               string `json:"awards"`                 // 获奖情况
	DownloadAddressTitle string `json:"download_address_title"` // 下载地址标题
	DownloadAddress      string `json:"download_address"`       //下载地址

}
