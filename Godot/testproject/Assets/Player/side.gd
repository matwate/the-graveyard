extends Node2D

@export var speed = 1000

func dash(body: PlayerManager) -> void:
	body.velocity = Vector2(0, 0)
	body.velocity += body.velocity.normalized() * speed
	
	
