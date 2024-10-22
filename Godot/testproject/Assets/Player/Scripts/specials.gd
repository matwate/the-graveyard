extends Node2D
@onready var player = $"../.."

func _physics_process(delta: float) -> void:
    if Input.is_action_just_pressed("special"):
        player.special()
