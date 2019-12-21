package main

import "time"

type Users struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
}

type Movies struct {
	Items []struct {
		Name         string `json:"Name"`
		ServerID     string `json:"ServerId"`
		ID           string `json:"Id"`
		RunTimeTicks int64  `json:"RunTimeTicks"`
		IsFolder     bool   `json:"IsFolder"`
		Type         string `json:"Type"`
		UserData     struct {
			PlaybackPositionTicks int       `json:"PlaybackPositionTicks"`
			PlayCount             int       `json:"PlayCount"`
			IsFavorite            bool      `json:"IsFavorite"`
			LastPlayedDate        time.Time `json:"LastPlayedDate"`
			Played                bool      `json:"Played"`
		} `json:"UserData"`
		ImageTags struct {
			Primary string `json:"Primary"`
			Logo    string `json:"Logo"`
			Thumb   string `json:"Thumb"`
		} `json:"ImageTags"`
		BackdropImageTags []string `json:"BackdropImageTags"`
		MediaType         string   `json:"MediaType"`
	} `json:"Items"`
}

type Movie struct {
	Name                  string    `json:"Name"`
	OriginalTitle         string    `json:"OriginalTitle"`
	ServerID              string    `json:"ServerId"`
	ID                    string    `json:"Id"`
	Etag                  string    `json:"Etag"`
	DateCreated           time.Time `json:"DateCreated"`
	CanDelete             bool      `json:"CanDelete"`
	CanDownload           bool      `json:"CanDownload"`
	PresentationUniqueKey string    `json:"PresentationUniqueKey"`
	SupportsSync          bool      `json:"SupportsSync"`
	Container             string    `json:"Container"`
	SortName              string    `json:"SortName"`
	PremiereDate          time.Time `json:"PremiereDate"`
	ExternalUrls          []struct {
		Name string `json:"Name"`
		URL  string `json:"Url"`
	} `json:"ExternalUrls"`
	MediaSources []struct {
		Protocol             string `json:"Protocol"`
		ID                   string `json:"Id"`
		Path                 string `json:"Path"`
		Type                 string `json:"Type"`
		Container            string `json:"Container"`
		Size                 int64  `json:"Size"`
		Name                 string `json:"Name"`
		IsRemote             bool   `json:"IsRemote"`
		RunTimeTicks         int64  `json:"RunTimeTicks"`
		SupportsTranscoding  bool   `json:"SupportsTranscoding"`
		SupportsDirectStream bool   `json:"SupportsDirectStream"`
		SupportsDirectPlay   bool   `json:"SupportsDirectPlay"`
		IsInfiniteStream     bool   `json:"IsInfiniteStream"`
		RequiresOpening      bool   `json:"RequiresOpening"`
		RequiresClosing      bool   `json:"RequiresClosing"`
		RequiresLooping      bool   `json:"RequiresLooping"`
		SupportsProbing      bool   `json:"SupportsProbing"`
		MediaStreams         []struct {
			Codec                  string  `json:"Codec"`
			TimeBase               string  `json:"TimeBase"`
			CodecTimeBase          string  `json:"CodecTimeBase"`
			VideoRange             string  `json:"VideoRange,omitempty"`
			DisplayTitle           string  `json:"DisplayTitle"`
			IsInterlaced           bool    `json:"IsInterlaced"`
			BitRate                int     `json:"BitRate,omitempty"`
			RefFrames              int     `json:"RefFrames,omitempty"`
			IsDefault              bool    `json:"IsDefault"`
			IsForced               bool    `json:"IsForced"`
			Height                 int     `json:"Height,omitempty"`
			Width                  int     `json:"Width,omitempty"`
			AverageFrameRate       float64 `json:"AverageFrameRate,omitempty"`
			RealFrameRate          float64 `json:"RealFrameRate,omitempty"`
			Profile                string  `json:"Profile,omitempty"`
			Type                   string  `json:"Type"`
			AspectRatio            string  `json:"AspectRatio,omitempty"`
			Index                  int     `json:"Index"`
			IsExternal             bool    `json:"IsExternal"`
			IsTextSubtitleStream   bool    `json:"IsTextSubtitleStream"`
			SupportsExternalStream bool    `json:"SupportsExternalStream"`
			Protocol               string  `json:"Protocol"`
			PixelFormat            string  `json:"PixelFormat,omitempty"`
			Level                  int     `json:"Level"`
			IsAnamorphic           bool    `json:"IsAnamorphic,omitempty"`
			Language               string  `json:"Language,omitempty"`
			DisplayLanguage        string  `json:"DisplayLanguage,omitempty"`
			Channels               int     `json:"Channels,omitempty"`
			SampleRate             int     `json:"SampleRate,omitempty"`
			ChannelLayout          string  `json:"ChannelLayout,omitempty"`
		} `json:"MediaStreams"`
		Formats             []interface{} `json:"Formats"`
		Bitrate             int           `json:"Bitrate"`
		RequiredHTTPHeaders struct {
		} `json:"RequiredHttpHeaders"`
		ReadAtNativeFramerate   bool `json:"ReadAtNativeFramerate"`
		DefaultAudioStreamIndex int  `json:"DefaultAudioStreamIndex"`
	} `json:"MediaSources"`
	CriticRating        int      `json:"CriticRating"`
	ProductionLocations []string `json:"ProductionLocations"`
	Path                string   `json:"Path"`
	OfficialRating      string   `json:"OfficialRating"`
	Overview            string   `json:"Overview"`
	Taglines            []string `json:"Taglines"`
	Genres              []string `json:"Genres"`
	CommunityRating     int      `json:"CommunityRating"`
	RunTimeTicks        int64    `json:"RunTimeTicks"`
	PlayAccess          string   `json:"PlayAccess"`
	ProductionYear      int      `json:"ProductionYear"`
	RemoteTrailers      []struct {
		URL  string `json:"Url"`
		Name string `json:"Name"`
	} `json:"RemoteTrailers"`
	ProviderIds struct {
		Tmdb           string `json:"Tmdb"`
		Imdb           string `json:"Imdb"`
		TmdbCollection string `json:"TmdbCollection"`
	} `json:"ProviderIds"`
	IsFolder bool   `json:"IsFolder"`
	ParentID string `json:"ParentId"`
	Type     string `json:"Type"`
	People   []struct {
		Name            string `json:"Name"`
		ID              string `json:"Id"`
		Role            string `json:"Role"`
		Type            string `json:"Type"`
		PrimaryImageTag string `json:"PrimaryImageTag,omitempty"`
	} `json:"People"`
	Studios []struct {
		Name string `json:"Name"`
		ID   int    `json:"Id"`
	} `json:"Studios"`
	GenreItems []struct {
		Name string `json:"Name"`
		ID   int    `json:"Id"`
	} `json:"GenreItems"`
	TagItems []interface{} `json:"TagItems"`
	UserData struct {
		PlaybackPositionTicks int  `json:"PlaybackPositionTicks"`
		PlayCount             int  `json:"PlayCount"`
		IsFavorite            bool `json:"IsFavorite"`
		Played                bool `json:"Played"`
	} `json:"UserData"`
	DisplayPreferencesID    string        `json:"DisplayPreferencesId"`
	Tags                    []interface{} `json:"Tags"`
	PrimaryImageAspectRatio float64       `json:"PrimaryImageAspectRatio"`
	MediaStreams            []struct {
		Codec                  string  `json:"Codec"`
		TimeBase               string  `json:"TimeBase"`
		CodecTimeBase          string  `json:"CodecTimeBase"`
		VideoRange             string  `json:"VideoRange,omitempty"`
		DisplayTitle           string  `json:"DisplayTitle"`
		IsInterlaced           bool    `json:"IsInterlaced"`
		BitRate                int     `json:"BitRate,omitempty"`
		RefFrames              int     `json:"RefFrames,omitempty"`
		IsDefault              bool    `json:"IsDefault"`
		IsForced               bool    `json:"IsForced"`
		Height                 int     `json:"Height,omitempty"`
		Width                  int     `json:"Width,omitempty"`
		AverageFrameRate       float64 `json:"AverageFrameRate,omitempty"`
		RealFrameRate          float64 `json:"RealFrameRate,omitempty"`
		Profile                string  `json:"Profile,omitempty"`
		Type                   string  `json:"Type"`
		AspectRatio            string  `json:"AspectRatio,omitempty"`
		Index                  int     `json:"Index"`
		IsExternal             bool    `json:"IsExternal"`
		IsTextSubtitleStream   bool    `json:"IsTextSubtitleStream"`
		SupportsExternalStream bool    `json:"SupportsExternalStream"`
		Protocol               string  `json:"Protocol"`
		PixelFormat            string  `json:"PixelFormat,omitempty"`
		Level                  int     `json:"Level"`
		IsAnamorphic           bool    `json:"IsAnamorphic,omitempty"`
		Language               string  `json:"Language,omitempty"`
		DisplayLanguage        string  `json:"DisplayLanguage,omitempty"`
		Channels               int     `json:"Channels,omitempty"`
		SampleRate             int     `json:"SampleRate,omitempty"`
		ChannelLayout          string  `json:"ChannelLayout,omitempty"`
	} `json:"MediaStreams"`
	ImageTags struct {
		Primary string `json:"Primary"`
		Logo    string `json:"Logo"`
		Thumb   string `json:"Thumb"`
	} `json:"ImageTags"`
	BackdropImageTags []string `json:"BackdropImageTags"`
	Chapters          []struct {
		StartPositionTicks int    `json:"StartPositionTicks"`
		Name               string `json:"Name"`
	} `json:"Chapters"`
	MediaType    string        `json:"MediaType"`
	LockedFields []interface{} `json:"LockedFields"`
	LockData     bool          `json:"LockData"`
	Width        int           `json:"Width"`
	Height       int           `json:"Height"`
}
