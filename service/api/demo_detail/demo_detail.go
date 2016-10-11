package demo_detail

import (
	faceslandmarkspicinfo "ManageCenter/service/model/faces_landmarks_picinfo"
	facespicinfo "ManageCenter/service/model/faces_picinfo"
	picinfo "ManageCenter/service/model/picinfo"
	vars "ManageCenter/service/vars"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"math"
	"strconv"
	"time"
)

type RepIdentify struct {
	Res       string  `json:"res"`
	Rate      float32 `json:"rate"`
	Url       string  `json:"url"`
	Review    bool    `json:"review"`
	CreatedAt string  `json:"created_at"`
}
type RepFaceLandmark struct {
	Url        string                                       `json:"url"`
	CreatedAt  string                                       `json:"created_at"`
	FaceBounds []*faceslandmarkspicinfo.FaceLandmarksStruct `json:"face_bounds"`
}

type RepFaceDetection struct {
	Url        string                           `json:"url"`
	CreatedAt  string                           `json:"created_at"`
	FaceBounds []*facespicinfo.FaceBoundsStruct `json:"face_bounds"`
}

func GetDemoDetail(c *gin.Context) {
	reqTime, err := time.Parse("2006-01-02", c.Query("time"))
	reqTpye := c.Query("type")
	if err != nil {
		log.Error(fmt.Sprintf("time Parse err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	endTime := reqTime.Add(24 * time.Hour)
	reqSord := c.Query("sord")
	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	if err != nil {
		log.Error(fmt.Sprintf("strconv Atoi err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	var (
		sord int
	)
	if reqSord == "asc" {
		sord = 1
	} else {
		sord = -1
	}
	if reqTpye == "faces" {
		detail, records, err := facespicinfo.GetDetail(reqTime, endTime, sord, pageIndex, pageSize)
		if err != nil {
			log.Error(fmt.Sprintf(" redis GetDetailCount err ", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrDemoNotFound.Code,
				"message": vars.ErrDemoNotFound.Msg,
			})
			return
		}
		repList := make([]*RepFaceDetection, 0, len(detail))
		for _, res := range detail {
			rep := &RepFaceDetection{
				Url:        res.OssUrl,
				CreatedAt:  res.CreatedAt.Format("2006-01-02 15:04:05"),
				FaceBounds: res.FaceBounds,
			}
			repList = append(repList, rep)
		}
		total := int(math.Ceil(float64(records) / float64(pageSize)))

		c.JSON(200, gin.H{
			"code":    0,
			"result":  repList,
			"page":    pageIndex,
			"total":   total,
			"records": records,
		})
	} else if reqTpye == "faces_landmarks" {
		detail, records, err := faceslandmarkspicinfo.GetDetail(reqTime, endTime, sord, pageIndex, pageSize)
		if err != nil {
			log.Error(fmt.Sprintf(" redis GetDetailCount err ", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrDemoNotFound.Code,
				"message": vars.ErrDemoNotFound.Msg,
			})
			return
		}
		repList := make([]*RepFaceLandmark, 0, len(detail))
		for _, res := range detail {
			rep := &RepFaceLandmark{
				Url:        res.OssUrl,
				CreatedAt:  res.CreatedAt.Format("2006-01-02 15:04:05"),
				FaceBounds: res.FaceLandmarks,
			}
			repList = append(repList, rep)
		}
		total := int(math.Ceil(float64(records) / float64(pageSize)))

		c.JSON(200, gin.H{
			"code":    0,
			"result":  repList,
			"page":    pageIndex,
			"total":   total,
			"records": records,
		})
	} else {
		detail, records, err := picinfo.GetDetail(reqTime, endTime, sord, pageIndex, pageSize)
		if err != nil {
			log.Error(fmt.Sprintf(" redis GetDetailCount err ", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrDemoNotFound.Code,
				"message": vars.ErrDemoNotFound.Msg,
			})
			return
		}
		repList := make([]*RepIdentify, 0, len(detail))
		for _, res := range detail {
			if res.Result.Porn == nil {
				continue
			}
			rep := &RepIdentify{
				Res:       res.Result.Porn.Res,
				Rate:      res.Result.Porn.Rate,
				Review:    res.Result.Porn.Review,
				Url:       res.OssUrl,
				CreatedAt: res.CreatedAt.Format("2006-01-02 15:04:05"),
			}
			repList = append(repList, rep)
		}
		total := int(math.Ceil(float64(records) / float64(pageSize)))

		c.JSON(200, gin.H{
			"code":    0,
			"result":  repList,
			"page":    pageIndex,
			"total":   total,
			"records": records,
		})
	}

}
