#version 330 core

in vec3 vp;
uniform mat4 camera;
uniform mat4 transform;
uniform mat4 projection;

void main(){
  gl_Position = projection * camera * transform * vec4(vp, 1.0);
}
