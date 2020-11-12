#version 330

uniform sampler2D tex;
uniform float time;
uniform mat4 camera;

in vec2 fragTexCoord;
out vec4 outputColor;

float speed = 0.0;

void main() {
  
  float strength = 1 / 100.0;

	if (camera[0].y != 0) {
		speed = (time * camera[0].y / camera[0].x) * 0.05; // 0.02
	}

	vec2 coord = fragTexCoord * 1.0;

	coord.x += sin((coord.x + speed) * 100) * strength;
	coord.y += cos((coord.y + speed) * 100) * strength;

	outputColor = texture(tex, coord) * 0.8;

}
