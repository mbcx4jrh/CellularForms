
#include "colors.inc"

#declare GROWTH_T = texture { pigment {White}}

camera {
    location <0, 2, -3>
    look_at <0, 0, 1>
}

light_source {
    <2, 4, -3> colour White
}

#include "foo.inc"