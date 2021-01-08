package parser

import (
	"GoCrawler/engine"
	"encoding/json"
	"log"
	"regexp"
)

const (
	TitleRe           = `<div class="title_all"><h1><font color=#[0-9a-z]+>([^>]*[^<])</font></h1></div>`                                                                                   // 标题
	CoverRe           = `<img border="0" src="((https|http):\/\/.*\.(jpg|png|jpeg))" alt="" style="MAX-WIDTH: 400px" />`                                                                    // 封面
	TranslatedNameRe  = `<br \/>◎译　　名　([^<]*)<br \/>`                                                                                                                                       // 译名
	MovieNameRe       = `<br \/>◎片　　名　([^<]*)<br \/>`                                                                                                                                       // 片名
	YearRe            = `<br \/>◎年　　代　([^<]*)<br \/>`                                                                                                                                       // 年代
	OriginRe          = `<br \/>◎产　　地　([^<]*)<br \/>`                                                                                                                                       // 产地
	TypeRe            = `<br \/>◎类　　别　([^<]*)<br \/>`                                                                                                                                       // 类别s
	LanguageRe        = `<br \/>◎语　　言　([^<]*)<br \/>`                                                                                                                                       // 语言
	SubtitleRe        = `<br \/>◎字　　幕　([^<]*)<br \/>`                                                                                                                                       // 字幕
	ReleaseDateRe     = `<br \/>◎上映日期　([^<]*)<br \/>`                                                                                                                                       // 上映时间
	ImdbScoreRe       = `<br \/>◎IMDb评分&nbsp;&nbsp;([^<]*)<br \/>`                                                                                                                          // IMDB评分
	DoubanScoreRe     = `<br \/>◎豆瓣评分　([^<]*)<br \/>`                                                                                                                                       // 豆瓣评分
	FileFormatRe      = `<br \/>◎文件格式　([^<]*)<br \/>`                                                                                                                                       // 文件格式
	VideoSizeRe       = `<br \/>◎视频尺寸　([^<]*)<br \/>`                                                                                                                                       // 视频尺寸
	FilmLengthRe      = `<br \/>◎片　　长　([^<]*)<br \/>`                                                                                                                                       // 片长
	DirectorRe        = `<br \/>◎导　　演　([^<]*)<br \/>`                                                                                                                                       // 导演
	ScreenwriterRe    = `<br \/>◎编　　剧　([^<]*)<br \/>`                                                                                                                                       // 编剧
	ToStarRe          = `<br \/>◎主　　演　([^◎]*)<br \/>`                                                                                                                                       // 主演
	LabelRe           = `<br \/>◎标　　签　([^<]*)<br \/>`                                                                                                                                       // 标签
	IntroductionRe    = `<br />◎简　　介(　)?<br /><br />(　)+([^◎]*[^<])(</span>)?<br /><br />`                                                                                                  //简介
	AwardsRe          = `<br \/>◎获奖情况<br /><br />　　([^◎]*)<br /><br /><strong>`                                                                                                             // 获奖情况
	DownloadAddressRe = `<a target="_blank" href="(.*)"><strong><font style="BACKGROUND-COLOR: #ff9966"><font color="#0000ff"><font size="4">([^>]*[^<])</font></font></font></strong></a>` //下载地址
)

// ParseMovieInfo 解析影视信息
func ParseMovieInfo(contents []byte, url string) engine.ParserResult {
	var movieinfo = MovieInfo{}

	titleRe := regexp.MustCompile(TitleRe)
	titleM := titleRe.FindSubmatch(contents)
	if len(titleM) == 0 {
		movieinfo.Title = ""
	} else {
		movieinfo.Title = string(titleM[1])
	}

	coverRe := regexp.MustCompile(CoverRe)
	coverM := coverRe.FindSubmatch(contents)
	if len(coverM) == 0 {
		movieinfo.Cover = ""
	} else {
		movieinfo.Cover = string(coverM[1])
	}

	translatedNameRe := regexp.MustCompile(TranslatedNameRe)
	translatedNameM := translatedNameRe.FindSubmatch(contents)
	if len(translatedNameM) == 0 {
		movieinfo.TranslatedName = ""
	} else {
		movieinfo.TranslatedName = string(translatedNameM[1])
	}

	movieNameRe := regexp.MustCompile(MovieNameRe)
	movieNameM := movieNameRe.FindSubmatch(contents)
	if len(movieNameM) == 0 {
		movieinfo.MovieName = ""
	} else {
		movieinfo.MovieName = string(movieNameM[1])
	}

	yearRe := regexp.MustCompile(YearRe)
	yearM := yearRe.FindSubmatch(contents)
	if len(yearM) == 0 {
		movieinfo.Year = ""
	} else {
		movieinfo.Year = string(yearM[1])
	}

	originRe := regexp.MustCompile(OriginRe)
	originM := originRe.FindSubmatch(contents)
	if len(originM) == 0 {
		movieinfo.Origin = ""
	} else {
		movieinfo.Origin = string(originM[1])
	}

	typeRe := regexp.MustCompile(TypeRe)
	typeM := typeRe.FindSubmatch(contents)
	if len(typeM) == 0 {
		movieinfo.Type = ""
	} else {
		movieinfo.Type = string(typeM[1])
	}

	languageRe := regexp.MustCompile(LanguageRe)
	languageM := languageRe.FindSubmatch(contents)
	if len(languageM) == 0 {
		movieinfo.Language = ""
	} else {
		movieinfo.Language = string(languageM[1])
	}

	subtitleRe := regexp.MustCompile(SubtitleRe)
	subtitleM := subtitleRe.FindSubmatch(contents)
	if len(subtitleM) == 0 {
		movieinfo.Subtitle = ""
	} else {
		movieinfo.Subtitle = string(subtitleM[1])
	}

	releaseDateRe := regexp.MustCompile(ReleaseDateRe)
	releaseDateM := releaseDateRe.FindSubmatch(contents)
	if len(releaseDateM) == 0 {
		movieinfo.ReleaseDate = ""
	} else {
		movieinfo.ReleaseDate = string(releaseDateM[1])
	}

	imdbScoreRe := regexp.MustCompile(ImdbScoreRe)
	imdbScoreM := imdbScoreRe.FindSubmatch(contents)
	if len(imdbScoreM) == 0 {
		movieinfo.ImdbScore = ""
	} else {
		movieinfo.ImdbScore = string(imdbScoreM[1])
	}

	doubanScoreRe := regexp.MustCompile(DoubanScoreRe)
	doubanScoreM := doubanScoreRe.FindSubmatch(contents)
	if len(doubanScoreM) == 0 {
		movieinfo.DoubanScore = ""
	} else {
		movieinfo.DoubanScore = string(doubanScoreM[1])
	}

	fileFormatRe := regexp.MustCompile(FileFormatRe)
	fileFormatM := fileFormatRe.FindSubmatch(contents)
	if len(fileFormatM) == 0 {
		movieinfo.FileFormat = ""
	} else {
		movieinfo.FileFormat = string(fileFormatM[1])
	}

	videoSizeRe := regexp.MustCompile(VideoSizeRe)
	videoSizeM := videoSizeRe.FindSubmatch(contents)
	if len(videoSizeM) == 0 {
		movieinfo.VideoSize = ""
	} else {
		movieinfo.VideoSize = string(videoSizeM[1])
	}

	filmLengthRe := regexp.MustCompile(FilmLengthRe)
	filmLengthM := filmLengthRe.FindSubmatch(contents)
	if len(filmLengthM) == 0 {
		movieinfo.FilmLength = ""
	} else {
		movieinfo.FilmLength = string(filmLengthM[1])
	}

	directorRe := regexp.MustCompile(DirectorRe)
	directorM := directorRe.FindSubmatch(contents)
	if len(directorM) == 0 {
		movieinfo.Director = ""
	} else {
		movieinfo.Director = string(directorM[1])
	}

	screenwriterRe := regexp.MustCompile(ScreenwriterRe)
	screenwriterM := screenwriterRe.FindSubmatch(contents)
	if len(screenwriterM) == 0 {
		movieinfo.Screenwriter = ""
	} else {
		movieinfo.Screenwriter = string(screenwriterM[1])
	}

	toStarRe := regexp.MustCompile(ToStarRe)
	toStarM := toStarRe.FindSubmatch(contents)
	if len(toStarM) == 0 {
		movieinfo.ToStar = ""
	} else {
		movieinfo.ToStar = string(toStarM[1])
	}

	labelRe := regexp.MustCompile(LabelRe)
	labelM := labelRe.FindSubmatch(contents)
	if len(labelM) == 0 {
		movieinfo.Label = ""
	} else {
		movieinfo.Label = string(labelM[1])
	}

	introductionRe := regexp.MustCompile(IntroductionRe)
	introductionM := introductionRe.FindSubmatch(contents)
	if len(introductionM) == 0 {
		movieinfo.Introduction = ""
	} else {
		movieinfo.Introduction = string(introductionM[3])
	}

	awardsRe := regexp.MustCompile(AwardsRe)
	awardsM := awardsRe.FindSubmatch(contents)
	if len(awardsM) == 0 {
		movieinfo.Awards = ""
	} else {
		movieinfo.Awards = string(awardsM[1])
	}

	downloadAddressRe := regexp.MustCompile(DownloadAddressRe)
	downloadAddressM := downloadAddressRe.FindSubmatch(contents)
	if len(downloadAddressM) == 0 {
		movieinfo.DownloadAddressTitle = ""
		movieinfo.DownloadAddress = ""
	} else {
		movieinfo.DownloadAddressTitle = string(downloadAddressM[2])
		movieinfo.DownloadAddress = string(downloadAddressM[1])
	}

	result := engine.ParserResult{}
	jsonData, err := json.Marshal(&movieinfo)
	if err != nil {
		log.Println("json序列化出错：", err.Error())
	}
	result.Items = append(result.Items, string(jsonData))

	return result
}
