pushd output
for file in *.pov
do
  povray $file
done
popd