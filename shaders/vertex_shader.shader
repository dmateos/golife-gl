#version 330 core

in vec3 vp;
in vec3 nm;
out vec3 nm_out;

uniform mat4 camera;
uniform mat4 transform;
uniform mat4 projection;

void main(){
  nm_out = nm;
  gl_Position = projection * camera * transform * vec4(vp, 1.0);
}
