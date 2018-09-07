apt-get update
apt-get install golang
apt-get install imagemagick imagemagick-doc 


#background for the card - total canvas
convert -size 1280x800 canvas:white background.png 
#background for the name
convert -size 900x250 canvas:white backwhite.png 

#print text onto backwhite.png, output to xyz_tmp.png
convert backwhite.png -gravity Center  -font Times-Bold -pointsize 144 -draw "text 0,0 'Barbaras Life'" Barbara_tmp.png

#put logo.jpg right to xyz_png, output to xyz_tmp-2.png
montage -background '#FFFFFF' -geometry +4+4 Barbara_tmp.png  logo.jpg  Barbara_tmp_2.png

#put text pus logo (xyz_tmp_2.png) onto the center of background.png, output to xyz.png
convert background.png -gravity Center Barbara_tmp_2.png -composite Barbara.png

#TODO: T1 The logo has to go upwards relative to the text.
#TODO: T2 Scaling for different text lengths, from 9 to 25 characters.
