#version 330

uniform float clickCoordY;
uniform float clickCoordX;
uniform vec2 u_resolution;

uniform sampler2D tex;
uniform float time;

in vec2 fragTexCoord;
out vec4 outputColor;

const mat2 m = mat2(1.80, 1.60, 1.2, 1.0);

float noise(in vec2 p) {
	return sin(p.x) + sin(p.y);
}

float fbm4( vec2 p ) {
  float f = 0.0;
  
  f += 0.5000 * noise(p); p = m * p * 2.02;
  f += 0.2500 * noise(p); p = m * p * 2.03;
  f += 0.1250 * noise(p); p = m * p * 2.01;
  f += 0.0625 * noise(p);

  return f / 0.9375;
}

void main() {
  
	float speed = time * 1 * 0.05;
	float strength = 1 / 100.0;

	vec2 coord = fragTexCoord * 1.0;

	coord.x += sin((coord.x + speed) * fbm4(coord) * 20) * strength;
	coord.y += cos((coord.y + speed) * fbm4(coord) * 20) * strength;

	outputColor = texture(tex, coord) * 0.8; // 0.8 for transparency
}
