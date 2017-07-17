#version 330 core

in vec3 vp;
uniform mat4 camera;
uniform mat4 transform;
uniform mat4 perspective;

void main(){
  gl_Position = perspective * camera * vec4(vp, 1.0);
}
