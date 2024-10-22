class_name State
extends Node2D


@export var anim_name : String = ""

var parent : CharacterBody2D


func enter() -> void:
	if anim_name != "":
		parent.animations.play(anim_name)

func exit() -> void:
	pass 

func process_input(event: InputEvent) -> State:
	return null
func process_frame(delta: float) -> State:
	return null

func process_physics(delta: float) -> State:
	return null
