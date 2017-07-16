#version 330 core

in vec3 vp;
uniform mat4 camera;
uniform mat4 transform;

void main(){
  gl_Position = camera * vec4(vp, 1.0);
}
