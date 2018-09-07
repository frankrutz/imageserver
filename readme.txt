apt-get update
apt-get install golang
apt-get install imagemagick imagemagick-doc 


convert -size 1280x800 canvas:white backwhite.png
convert backwhite.png -gravity Center -font Times-Bold -pointsize 144 -draw "text 0,0 'Barbaras Life'" inputname_tmp.png
convert inputname_tmp.png  logo.jpg  -gravity east -composite inputname.png
rm inputname_tmp.png
