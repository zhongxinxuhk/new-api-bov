package relayconvert

import relaymedia "github.com/zhongxinxuhk/new-api-bov/relaykit/relayconvert/internal/media"

type MediaResolver = relaymedia.MediaResolver

func SetMediaResolver(resolver MediaResolver) {
	relaymedia.SetMediaResolver(resolver)
}
