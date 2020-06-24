#version 330

uniform sampler2D tex;
uniform float time;

in vec2 fragTexCoord;
out vec4 outputColor;

void main() {
  float speed = time * 1 * 0.05;
  float strength = 1 / 100.0;

  vec2 coord = fragTexCoord * 1;

  coord.x += sin((coord.x + speed) * 60) * strength;
  coord.y += cos((coord.y + speed) * 60) * strength;

  outputColor = texture(tex, coord) * 0.8;

}
