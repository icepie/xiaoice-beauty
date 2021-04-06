package model

// Content 分析图片接口提交结构中的内容
type Content struct {
	Imageurl string `json:"imageUrl"`
}

// AnalyzeBody 分析图片接口提交结构
type AnalyzeImgBody struct {
	Msgid      int64   `json:"MsgId"`
	Createtime int64   `json:"CreateTime"`
	Traceid    string  `json:"TraceId"`
	Content    Content `json:"Content"`
}

// AnalyzeBody 分析图片接口返回结构
type AnalyzeImgRte struct {
	Msgid      string      `json:"msgId"`
	Timestamp  int64       `json:"timestamp"`
	Receiverid interface{} `json:"receiverId"`
	Content    struct {
		Text     string `json:"text"`
		Imageurl string `json:"imageUrl"`
		Metadata struct {
			FbrID0             string      `json:"FBR_Id0"`
			FbrKey0            string      `json:"FBR_Key0"`
			FbrScore0          float64     `json:"FBR_Score0"`
			FbrID1             string      `json:"FBR_Id1"`
			FbrKey1            string      `json:"FBR_Key1"`
			FbrScore1          float64     `json:"FBR_Score1"`
			FbrID2             string      `json:"FBR_Id2"`
			FbrKey2            string      `json:"FBR_Key2"`
			FbrScore2          float64     `json:"FBR_Score2"`
			FbrCnt             int         `json:"FBR_Cnt"`
			FbrMode            string      `json:"FBR_Mode"`
			Gender             string      `json:"gender"`
			Score              float64     `json:"score"`
			CommentPrefix      interface{} `json:"Comment_Prefix"`
			Isceleb            string      `json:"isCeleb"`
			Enableuserfeedback bool        `json:"enableUserFeedback"`
			Isemoji            string      `json:"isEmoji"`
			FaceNumber         int         `json:"face_number"`
			Reportimgurl       string      `json:"reportImgUrl"`
			FacePoints         string      `json:"face_points"`
			Answerfeed         string      `json:"AnswerFeed"`
			W                  string      `json:"w"`
			Aid                string      `json:"aid"`
		} `json:"metadata"`
	} `json:"content"`
}
