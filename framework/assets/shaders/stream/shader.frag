#version 330

uniform sampler2D tex;
uniform float time;
uniform mat4 camera;

in vec2 fragTexCoord;
out vec4 outputColor;

float random (in float x) {
  return fract(sin(x)*1e4);
}

float random (in vec2 st) {
  return fract(sin(dot(st.xy, vec2(12.9898, 78.233))) * 43758.5453123);
}

float pattern(vec2 st, vec2 v, float t) {
  vec2 p = floor(st + v);
  return step(t, random(100. + p * .000001) + random(p.x) * 0.5 );
}

void main() {
  
  vec2 st = gl_FragCoord.xy / vec2(400, 400);
  st.x *= 400.0 / 400.0;

  vec2 grid = vec2(100., 50.);
  st *= grid;

  vec2 ipos = floor(st);  // integer
  vec2 fpos = fract(st);  // fraction

  vec2 vel = vec2(time * 2. * max(grid.x, grid.y)); // time
  vel *= vec2(-1., 0.0) * random(1.0 + ipos.y); // direction

  // Assign a random value base on the integer coord
  vec2 offset = vec2(100.1, 0.);

  vec3 color = vec3(0.);
  color.r = pattern(st + offset, vel, 0.5);
  color.g = pattern(st, vel, 0.5);
  color.b = pattern(st - offset, vel, 0.5);

  // Margins
  // color *= step(0.2, step(fpos.y, fpos.x));
  color *= step(0.2, fpos.y);
  outputColor = vec4(1.0 -color, 1.0);
}



