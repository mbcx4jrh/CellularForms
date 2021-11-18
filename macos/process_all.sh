pushd output
for file in *.cf
do
  ../form/form -h ../povray/default_render.pov -i $file
  povray $file.pov
done
popd