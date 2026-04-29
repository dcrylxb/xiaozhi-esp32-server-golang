package request

type V3Request struct {
	User      V3User      `json:"user"`
	ReqParams V3ReqParams `json:"req_params"`
	Additions V3Additions `json:"Additions"`
}

type V3User struct {
	UID string `json:"uid"`
}

type V3AudioParams struct {
	Format       string `json:"format"`
	SampleRate   int    `json:"sample_rate"`
	SpeechRate   int    `json:"speech_rate,omitempty"`
	LoudnessRate int    `json:"loudness_rate,omitempty"`
}

type V3ReqParams struct {
	Text        string        `json:"text"`
	Speaker     string        `json:"speaker"`
	AudioParams V3AudioParams `json:"audio_params"`
	Model       string        `json:"model,omitempty"`
}

type V3Additions struct {
	DisableMarkdownFilter bool `json:"disable_markdown_filter"`
}

func NewV3Request(text, speaker, audioFormat string, sampleRate int, requestModel string, uid string) V3Request {
	return V3Request{
		User: V3User{UID: uid},
		ReqParams: V3ReqParams{
			Text:    text,
			Speaker: speaker,
			AudioParams: V3AudioParams{
				Format:     audioFormat,
				SampleRate: sampleRate,
			},
			Model: requestModel,
		},
		Additions: V3Additions{
			DisableMarkdownFilter: true,
		},
	}
}
