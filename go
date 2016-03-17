latex paper
#latex paper
bibtex paper
bibtex paper
latex paper
latex paper

dvips -D 600 -j0 -Ppdf -G0 -t letter -o paper.ps paper.dvi

ps2pdf -dPDFSETTINGS=/printer -dCompatibilityLevel=1.7 -dMaxSubsetPct=100 -dSubsetFonts=true -dEmbedAllFonts=true -sPAPERSIZE=letter paper_final.ps


word=`detex paper_final.tex | wc -w`
if [ "$word" -le 4500 ]; then
        echo -en "\033[0;32m"
else
        echo -en "\033[0;31m"
fi
echo ===========WORDS========== ${word}
echo -en "\033[0m"
