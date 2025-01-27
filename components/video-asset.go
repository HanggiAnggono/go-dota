package components

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func VideoAsset(videoPath string, children ...Node) Node {
	return Video(
		Poster(fmt.Sprintf("https://cdn.akamai.steamstatic.com/apps/dota2/videos/dota_react/%s.png", videoPath)),
		AutoPlay(),
		Preload("auto"),
		Loop(),
		PlaysInline(),
		Source(
			Type("video/mp4; codecs=\"hvc1\""),
			Src(fmt.Sprintf("https://cdn.akamai.steamstatic.com/apps/dota2/videos/dota_react/%s.mov", videoPath)),
		),
		Source(
			Type("video/webm"),
			Src(fmt.Sprintf("https://cdn.akamai.steamstatic.com/apps/dota2/videos/dota_react/%s.webm?undefined", videoPath)),
		),
		Img(
			Src(fmt.Sprintf("https://cdn.akamai.steamstatic.com/apps/dota2/videos/dota_react/%s.png", videoPath)),
		),
		Group(children),
	)
}
