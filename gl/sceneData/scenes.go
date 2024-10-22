package scenedata

import "matwa/caobaEngine/scenes"

var AvailableScenes = map[string]func() scenes.Scene{
	"bot":      GetBotScene,
	"backpack": GetBackpackScene,
	"nada":     GetNadaScene,
}
