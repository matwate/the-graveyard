extends Node2D
class_name CelestialBody

const G = 6.67430e-11

var mass: float
var pos: Vector2
var vel: Vector2
var accel: Vector2

func _init(mass: float, pos: Vector2, vel: Vector2) -> void:
    self.mass = mass
    self.pos = pos
    self.vel = vel
    self.accel = Vector2.ZERO

func update(delta: float, bodies: Array):
    # RK4 integration
    var k1 = calculate_derivatives(position, velocity)
    var k2 = calculate_derivatives(position + k1.x * delta / 2, velocity + k1.y * delta / 2)
    var k3 = calculate_derivatives(position + k2.x * delta / 2, velocity + k2.y * delta / 2)
    var k4 = calculate_derivatives(position + k3.x * delta, velocity + k3.y * delta)

    position += (k1.x + 2*k2.x + 2*k3.x + k4.x) * delta / 6
    velocity += (k1.y + 2*k2.y + 2*k3.y + k4.y) * delta / 6

    # Update acceleration for the next frame
    acceleration = calculate_acceleration(bodies)

func calculate_derivatives(pos: Vector2, vel: Vector2) -> Array:
    return [vel, acceleration]

func calculate_acceleration(bodies: Array) -> Vector2:
    var total_acceleration = Vector2.ZERO
    for body in bodies:
        if body != self:
            var r = body.position - position
            var force = G * mass * body.mass / r.length_squared()
            total_acceleration += force * r.normalized() / mass
    return total_acceleration

func set_circular_orbit(center: CelestialBody, orbit_radius: float):
    position = center.position + Vector2(orbit_radius, 0)
    var orbit_speed = sqrt(G * center.mass / orbit_radius)
    velocity = Vector2(0, -orbit_speed)