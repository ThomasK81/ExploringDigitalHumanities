# Exploring Digital Humanities

This is the course repository for [Exploring Digital Humanities: A hands-on introduction to data-driven research](http://www.thomaskoentges.io/upcoming.html). The repository will contain all sample scripts and guidelines used in the course.

## Software and Programming Languages Used in the Course

This is a list of programmes and programming languages used in the course. If you bring your own computer, please make sure that they are up and running. We will go through the installation of the [Go Progamming Language](https://golang.org) on the first day in class, since it is a bit more complicated. That said, some of the scripts, you will be able to run in the [Go Playground](https://play.golang.org), which does not require any installation.  
With one exception all tools and programmes run on all major operating systems.  
If you have problems with the installation of any of this, don't worry. We will have time on Day 1 to get everything up and running. We won't use programming languages before Day 2. I am aware that the whole list looks a bit daunting, but installation will be quicker than it looks.

### Programming Languages
- [Go](https://golang.org), version 1.11
- [R](https://www.r-project.org), version 3.5.2

### Tools
- [RStudio](https://www.rstudio.com); an editing environment for R Code.
- A good text editor. I recommend [Visual Studio Code](https://code.visualstudio.com)
- [OpenRefine](http://openrefine.org); used to be GoogleRefine. It's a great GUI for working with messy data.
- [Google Fusion](https://en.wikipedia.org/wiki/Google_Fusion_Tables); you don't need to install it, but you need a Google account.
- [ToPān](https://github.com/ThomasK81/ToPan); we will install ToPān together in class.
- [Imagemagick](https://imagemagick.org); open software that runs on all major systems.
- [Gephi](https://gephi.org); open software that runs on all major systems.
- [Tableau Public](https://public.tableau.com); commercial, but free software (does not work on Linux). Is only optional for this course.

### Web Browser
While you can use a browser of your choice, I found that [Firefox](https://www.mozilla.org/en-US/firefox/new/) and [Opera](https://www.opera.com) are particularly good for testing the user interfaces we will produce in this class.  

### Go Libraries
If you want you can install all necessary Go libraries beforehand. However, if you have not done so, we will install them together in the course as we go along. You can install them from your Terminal / Commandline with `go get [package_name]`, e.g. `go get github.com/gorilla/mux`.  
We will use the following external Go libraries (there might be more):
- `github.com/ThomasK81/gocite`
- `github.com/gorilla/mux`

### R Libraries
R has a lot of handy libraries for data analysis and restructuring. We will use a good number of them, but that also means that we have to write fewer functions ourselves. If you want you can install all necessary R libraries beforehand. However, if you have not done so, we will install them together in the course as we go along. You can install them from inside an R Terminal with `install.packages("[package_name]")`, e.g. `install.packages("data.table")`. 
We will definitely use the following R libraries (there might be more):
- `tidyverse`
- `leaflet`
- `data.table`
- `stylo`
- `shiny`
- `lattice`
- `Rtsne`
- `lda`
- `LDAvis`
- `RCurl`
- `httr`
- `jsonlite`
- `XML`
- `xml2`
- `markdown`
- `philentropy`
- `ape`
- `geiger`
- `cld2`
- `reshape2`
- `spatstat`
- `gridExtra`
- `grid`
- `ggthemes`
- `ggpubr`
- `doParallel`
- `foreach`
- `ngram`
- `maps`
- `mapdata`
- `rgeos`
- `maptools`
- `mapproj`
- `PBSmapping`
- `plyr`

