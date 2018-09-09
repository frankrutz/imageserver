#!/bin/bash

export PRINTNAME=$1
export PRINTTEXT=$2
export BACKFORMAT=$3"x200"
export TMP1PNG=$PRINTNAME"_tmp.png"
export TMP2PNG=$PRINTNAME"_tmp_2.png"
export OUTPNG=$PRINTNAME".png"

#background for the card - total canvas
convert -size 1280x800 canvas:white background.png 

#background for the name
convert -size $BACKFORMAT canvas:white backwhite.png 

#print text onto backwhite.png, output to TMP1PNG
convert backwhite.png -gravity Southeast  -font Times-Bold -pointsize 144 -draw "text 0,0 '$PRINTTEXT Life'"  $TMP1PNG

#put logo.jpg right to TMP1PNG, output to TMP2PNG
montage -background '#FFFFFF' -geometry +4+4 $TMP1PNG  logo.jpg  $TMP2PNG

#resize the image, needed for differing lengths of the names
convert $TMP2PNG -resize 800 $TMP2PNG

#put text plus logo () onto the center of background.png, output to OUTPNG
convert background.png -gravity Center $TMP2PNG -composite $OUTPNG

mv $OUTPNG ./tmp

#clean up
rm $TMP1PNG
rm $TMP2PNG
