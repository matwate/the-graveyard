extends Node

@export 
var startingState: State
var currentState: State


func init_state(parent: CharacterBody2D) -> void:
    for child in get_children():
        child.parent = parent
    change_state(startingState)

func change_state(state: State) -> void:
    if currentState:
        currentState.exit()
    currentState = state
    currentState.enter()

func process_physics(delta: float) -> void:
    var new_state = currentState.process_physics(delta)
    if new_state:
        change_state(new_state)

func process_input(event: InputEvent) -> void:
    var new_state = currentState.process_input(event)
    if new_state:
        change_state(new_state)

func process_frame(delta: float) -> void:
    var new_state = currentState.process_frame(delta)
    if new_state:
        change_state(new_state)