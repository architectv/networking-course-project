#!/usr/bin/bash -x

FILENAME[1]="../StepanenkoVA.pdf"
FULLNAME[1]="Степаненко Владимир Александрович"
ABBRNAME[1]="В.А. Степаненко"
TEACHER[1]="И.И. Иванов"
GROUP[1]="ИУ7--72Б"

FILENAME[2]="../VasukovAV.pdf"
FULLNAME[2]="Васюков Алексей Владимирович"
ABBRNAME[2]="А.В. Васюков"
TEACHER[2]="И.И. Иванов"
GROUP[2]="ИУ7--72Б"

FILENAME[3]="../VolkovEA.pdf"
FULLNAME[3]="Волков Егор Андреевич"
ABBRNAME[3]="Е.А. Волков"
TEACHER[3]="И.И. Иванов"
GROUP[3]="ИУ7--75Б"

FLAGS="-pdf -outdir=out -f -interaction=nonstopmode -shell-escape -silent"

cd src
for i in 1 2 3;
do
	FILENAME=${FILENAME[$i]}
	env \
	"FULLNAME=${FULLNAME[$i]}" \
	"ABBRNAME=${ABBRNAME[$i]}" \
	"TEACHER=${TEACHER[$i]}" \
	"GROUP=${GROUP[$i]}" \
	latexmk $FLAGS main.tex
	mv out/main.pdf $FILENAME
done
