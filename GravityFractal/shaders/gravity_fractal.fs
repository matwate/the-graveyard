#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Output fragment color
out vec4 finalColor;

// Uniform inputs
uniform vec2 attractorPositions[10];  // Assuming max 10 attractors
uniform float attractorMasses[10];
uniform int attractorCount;

const int MAX_ITERATIONS = 100;
const float G = 1;
const float TIME_STEP = 0.1;

vec2 calculateForce(vec2 position, vec2 attractorPos, float attractorMass) {
    vec2 direction = attractorPos - position;
    float distanceSquared = max(dot(direction, direction), 1e-6); // Avoid division by zero
    float forceMagnitude = G * attractorMass / distanceSquared;
    return normalize(direction) * forceMagnitude;
}

void main() {
    vec2 position = fragTexCoord * vec2(1600.0, 900.0);  // Scale to window size
    vec2 velocity = vec2(0.0, 0.0);
    
    float closestDistance = 1e10;
    int closestAttractor = 0;
    
    for (int i = 0; i < MAX_ITERATIONS; i++) {
        vec2 totalForce = vec2(0.0, 0.0);
        
        for (int j = 0; j < attractorCount; j++) {
            totalForce += calculateForce(position, attractorPositions[j], attractorMasses[j]);
            
            float dist = distance(position, attractorPositions[j]);
            if (dist < closestDistance) {
                closestDistance = dist;
                closestAttractor = j;
            }
        }
        
        vec2 acceleration = totalForce;  // Assuming unit mass for simplicity
        velocity += acceleration * TIME_STEP;
        position += velocity * TIME_STEP;
    }
    
    // Color based on the closest attractor
    float hue = float(closestAttractor) / float(attractorCount);
    vec3 color = vec3(
        abs(fract(hue + 1.0) * 6.0 - 3.0) - 1.0,
        abs(fract(hue + 2.0/3.0) * 6.0 - 3.0) - 1.0,
        abs(fract(hue + 1.0/3.0) * 6.0 - 3.0) - 1.0
    );
    finalColor = vec4(clamp(color, 0.0, 1.0), 1.0);
}