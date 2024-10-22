extends Node2D

@export var PLAYER_SPEED = 300
@export var JUMP_SPEED = -350
var jump_frames = 0
var gravity_scale = 1
var can_jump = true
func movement(body: CharacterBody2D, delta: float) -> void:
    var xDir := Input.get_axis("move_left", "move_right")
    var yDir := Input.get_axis("move_up", "move_down")
    var dir = Vector2(xDir, yDir)

    walk(body, dir)
    handle_jump(body, delta)
    body.move_and_slide()




func handle_jump(body: CharacterBody2D, delta: float) -> void:
    var gravity = body.get_gravity()

    if Input.is_action_pressed("jump"):
        if jump_frames < 20:
            gravity_scale = 0.7
            jump_frames += 1
            if body.is_on_floor():
                body.velocity.y = JUMP_SPEED
            body.velocity.y -= jump_frames
    if Input.is_action_just_released("jump"):
        jump_frames = 0
        gravity_scale = 1

    # Disable jump buffering
    if body.is_on_floor():
        if jump_frames > 0:
            can_jump = false
        else:
            can_jump = true


    if not body.is_on_floor():
        if jump_frames < 20 and jump_frames > 0:
            gravity_scale = 1

        # Fast fall
        if Input.is_action_pressed("move_down"):
            gravity_scale = 3
            can_jump = false

    body.velocity += gravity * delta * gravity_scale

func walk(body :CharacterBody2D, dir : Vector2) -> void:
    body.velocity = Vector2(dir.x * PLAYER_SPEED, body.velocity.y) 