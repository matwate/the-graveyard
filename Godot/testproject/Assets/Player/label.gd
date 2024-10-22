extends Label

@onready var player: CharacterBody2D = $".."

func _process(delta: float) -> void:
    var velstring = str(player.velocity)
    set_text(velstring)