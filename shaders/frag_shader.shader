#version 330 core

// Ouput data
in vec3 nm_out;
out vec3 color;

void main()
{
  // Output color = red 
  color = nm_out;
}
