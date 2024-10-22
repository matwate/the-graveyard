#version 430 core
out vec4 FragColor;

in vec2 TexCoords;
in vec3 Normal;
in vec3 FragPos;

uniform sampler2D texture_diffuse1;
uniform sampler2D texture_specular1;
uniform vec4 lightColor;
uniform vec3 viewPos;

vec3 CalculateLight(vec3 lightPos);

void main()
{    
   
   vec3 result = CalculateLight(vec3(0.0, 0.0, 2.0));
   vec3 result2 = CalculateLight(vec3(0.0, 0.0, -2.0));

   vec3 light = result + result2;

   FragColor = vec4(light, 1.0);
}

vec3 CalculateLight(vec3 lightPos){
float ambientStrength = 0.1;
   vec3 ambient = ambientStrength * lightColor.rgb;

   vec3 norm = normalize(Normal);
   vec3 lightDir = normalize(lightPos - FragPos);
   float diff = max(dot(norm, lightDir), 0.0);
   vec3 diffuse = diff * lightColor.rgb;


   vec4 tex_Specular = texture(texture_specular1, TexCoords);
   vec3 viewDir = normalize(viewPos - FragPos);
   vec3 reflectDir = reflect(-lightDir, norm);
   float spec = pow(max(dot(viewDir, reflectDir), 0.0), 32);
   vec3 specular = tex_Specular.x * spec * lightColor.rgb;

   vec3 result = (ambient + diffuse + specular) * texture(texture_diffuse1, TexCoords).rgb;

   return result;
}