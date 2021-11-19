pushd $1
for file in *.cf
do
  ../form/form -v -traits -h ../povray/default_render.pov -i $file -s 3 -slice -1000
  povray $file.pov
done
popd