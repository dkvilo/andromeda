#version 330

uniform sampler2D tex;
uniform float time;
uniform mat4 camera;

in vec2 fragTexCoord;
out vec4 outputColor;

float speed = 0.0;

const mat2 m = mat2(0.80, 0.60, -0.60, 0.80);

float noise(in vec2 p) {
	return sin(p.x) * sin(p.y);
}

float fbm4( vec2 p ) {
  float f = 0.0;
  
  f += 0.5000 * noise(p); p = m * p * 2.02;
  f += 0.2500 * noise(p); p = m * p * 2.03;
  f += 0.1250 * noise(p); p = m * p * 2.01;
  f += 0.0625 * noise(p);

  return f / 0.9375;
}

float fbm6(vec2 p) {
  float steamId = 0.120836380;
  float f = steamId;

  f += steamId * (0.5 + 0.5 * noise(p)); p = m * p * 2.02;
  f += steamId / 2 * (0.5 + 0.5 * noise(p)); p = m * p * 2.03;
  f += steamId /  3 * (0.5 + 0.5 * noise(p)); p = m * p * 2.01;
  f += steamId / 4 * (0.5 + 0.5 * noise(p)); p = m * p * 2.04;
  f += steamId / 5 * (0.5 + 0.5 * noise(p)); p = m * p * 2.01;
  f += steamId / 6 * (0.5 + 0.5 * noise(p));
  
  return f / 0.96875;
}

vec2 fbm4_2(vec2 p) {
  return vec2(fbm4(p), fbm4(p + vec2(7.8)));
}

vec2 fbm6_2(vec2 p) {
  return vec2(fbm6(p + vec2(16.8)), fbm6(p + vec2(11.5)));
}

float func(vec2 q, out vec4 ron) {
  q += 0.03 * sin(vec2(0.27, 0.23) * time + length(q) * vec2(4.1, 4.3));
  
  // q.x += 1.0 * sin(-2.27 * time + 1.0 * 2.1);

	vec2 o = fbm4_2(0.9 * q);

  o += 0.04 * sin(vec2(0.12, 0.14) * time + length(o));

  vec2 n = fbm6_2(3.0 * o);

	ron = vec4(o, n);

  float f = 0.5 + 0.5 * fbm4(1.8 * q + 6.0 * n);

  return mix(f, f * f * f * 3.5, f * abs(n.x));
}

void main() {
  
  float strength = 3.2;

	if (camera[0].y != 0) {
		speed = (time * camera[0].y / camera[0].x) * 0.01;
	}

	vec2 coord = fbm4_2(fragTexCoord) * 10.5;

  vec2 p = (2.0 * coord - 1.5) / strength;
  float e = 2.0 / strength;

  vec4 on = vec4(0.0);
  float f = func(p, on);

	vec3 color = vec3(0.0);
  color = mix(vec3(0.2, 0.1, 0.4), vec3(0.3, 0.05, 0.05), f );
  color = mix(color, vec3(0.9, 0.9, 0.9), dot(on.zw, on.zw) );
  color = mix(color, vec3(0.4, 0.3, 0.3), 0.2 + 0.5 * on.y * on.y);
  color = mix(color, vec3(0.0, 0.2, 0.4), 0.5 * smoothstep(1.2, 1.3, abs(on.z) + abs(on.w)));
  color = clamp(color * f * 2.0, 0.0, 1.0);

  vec4 temp;
  vec3 normalized = normalize(vec3(func(p + vec2(e, 0.0), temp) -f, 2.0 * e, func(p + vec2(0.0, e), temp) - f));

  vec3 lig = normalize(vec3(0.9, 0.2, -0.4));
  float dif = clamp(0.3 + 0.7 * dot(normalized, lig), 0.0, 1.0);
  vec3 lin = vec3(0.70, 0.90, 0.95) * (normalized.y * 0.5 + 0.5) + vec3(0.15, 0.10, 0.05) * dif;

  color *= 1.2 * lin;
	color = 1.0 - color;
	color = 1.1 * color * color;

  outputColor = vec4(color, 1.0) * 1.0;
}

