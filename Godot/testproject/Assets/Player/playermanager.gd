class_name Player
extends Node2D

@onready var animations = $animations
@onready var stateMachine = $stateMachine


func _ready( ) -> void:
	stateMachine.init_state(self)

func _unhandled_input(event: InputEvent) -> void:
	stateMachine.process_input(event)

func _physics_process(delta: float) -> void:
	stateMachine.process_physics(delta)

func _process(delta: float) -> void:
	stateMachine.process_frame(delta)
