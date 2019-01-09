# Exploring Digital Humanities

This is the course repository for [Exploring Digital Humanities: A hands-on introduction to data-driven research](http://www.thomaskoentges.io/upcoming.html). The repository will contain all sample scripts and guidelines used in the course.

## Software and Programming Languages Used in the Course

This is a list of programmes and programming languages used in the course. If you bring your own computer, please make sure that these are installed (as much as possible). That said, we will go through the installation of the [Go Progamming Language](https://golang.org) on the first day in class, since it is a bit more complicated. Some scripts can be run in the [Go Playground](https://play.golang.org), which does not require any installation.  

All tools and programmes run on all major operating systems. If you have problems with the installation of any of this, don't worry. We will have time on Day 1 to get everything up and running. We won't use programming languages before Day 2. I am aware that the whole list looks a bit daunting, but installation will be quicker than it looks.  

If you don't want to go through the whole list, please try and install at least R, RStudio, OpenRefine, and a text editor of your choice (e.g., [Visual Studio Code](https://code.visualstudio.com)).

### Programming Languages
- [Go](https://golang.org), version 1.11
- [**R**](https://www.r-project.org), version 3.5.2

### Tools
- [**RStudio**](https://www.rstudio.com); an editing environment for R Code.
- **A good text editor.** I recommend [Visual Studio Code](https://code.visualstudio.com).
- [**OpenRefine**](http://openrefine.org); used to be called GoogleRefine. It's a great GUI for working with messy data.
- [Google Fusion](https://en.wikipedia.org/wiki/Google_Fusion_Tables); you don't need to install it, but you need a Google account (i.e., gmail).
- [ToPān](https://github.com/ThomasK81/ToPan); we will install ToPān together in class on Day 4.
- [Imagemagick](https://imagemagick.org); open software that runs on all major systems (just used on Day 5).
- [Gephi](https://gephi.org); open software that runs on all major systems (just used on Day 5).

### Web Browser
While you can use a browser of your choice, I have found that [Firefox](https://www.mozilla.org/en-US/firefox/new/) and [Opera](https://www.opera.com) are particularly good for testing the user interfaces we will produce in this class.  

### Go Libraries
If you want, you can install all necessary Go libraries beforehand. However, if you have not done so, we will install them together in the course as we go along. You can install them from your Terminal / Commandline with `go get [package_name]`, e.g. `go get github.com/gorilla/mux`.  

We will use the following external Go libraries:
- `github.com/ThomasK81/gocite`
- `github.com/gorilla/mux`

### R Libraries
R has a lot of handy libraries for data analysis and restructuring. We will use several, which means we have to write fewer functions ourselves. If you want, you can install all necessary R libraries beforehand. However, if you have not done so, we will install them together in the course as we go along. You can install them from inside an R Terminal with `install.packages("[package_name]")`, e.g. `install.packages("data.table")`.  

We will use the following R libraries:
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

