extends State

@export var move_state: State

func enter() -> void:
    super()
    parent.velocity = Vector2.ZERO

func process_input(event: InputEvent) -> State:
    if Input.is_action_pressed("move_right") or Input.is_action_pressed("move_left"):
        return move_state
    return null

func process_physics(delta: float) -> State:
    if parent.is_on_floor():
        parent.velocity.y = 0
    else:
        parent.velocity.y += gravity * delta
    return null
