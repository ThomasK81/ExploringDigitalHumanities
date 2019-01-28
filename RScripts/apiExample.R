# install.packages("jsonlite")

library(httr)
library(jsonlite)

my_api_key <- ""
url1 <- paste("http://api.digitalnz.org/v3/records.json?api_key=", my_api_key, sep = "") 
query1 <- "&text=rugby&and[collection][]=New+Zealand+Cartoon+Archive&and[year][]=2011&fields=id&per_page=1&page=1"
query2 <- "&text=rugby&and[collection][]=New+Zealand+Cartoon+Archive&and[year][]=2011&fields=id,origin_url,title,description,collection,dc_identifier,large_thumbnail_url,authorities&per_page=100&page="

### How many API calls
number_calls <- ceiling(
  as.numeric(
    sub("\\D*(\\d+).*", "\\1", 
        content(
          GET(
            paste(url1, query1, sep = "")
          ),
          "text")
    )
  ) / 100)


### To understand sub("\\D*(\\d+).*", "\\1", x) see http://stat545.com/block022_regular-expression.html

api_calls <- paste(url1, query2, as.character(c(1:number_calls)), sep = "")

### let's test

data <- fromJSON(api_calls[1])
results <- data[[1]][2]
ID <- unlist(results, recursive = FALSE)[[1]]
Title <- unlist(results, recursive = FALSE)[[2]]
Description <- unlist(results, recursive = FALSE)[[3]]

TestData <- data.frame(ID,Title,Description)

# GET DCDL identifiers
Identifier <- unlist(results, recursive = FALSE)[[5]]
DCDL_IDs <- vector()

for(i in 1:length(Identifier)) {
  DCDL_IDs[i] <- Identifier[[i]][5]
}

# Add it to the data.frame

TestData['DCDL_IDs'] <- DCDL_IDs

# Let's extract the date from the title

install.packages('stringr')
library(stringr)
Date <- str_extract(Title, "\\d{1,2}\\s\\w*\\s\\d{4}")
TestData['Date'] <- Date

### How would you add a Collection to the data.frame? 
### Add the cartoonist collection to the data.frame!

##### Your code goes here

### write a function

DataframeFromCall <- function(x) {
  data <- fromJSON(x)
  results <- data[[1]][2]
  ID <- unlist(results, recursive = FALSE)[[1]]
  Title <- unlist(results, recursive = FALSE)[[2]]
  Description <- unlist(results, recursive = FALSE)[[3]]
  Identifier <- unlist(results, recursive = FALSE)[[5]]
  DCDL_IDs <- vector()
  for(i in 1:length(Identifier)) {
    DCDL_IDs[i] <- Identifier[[i]][5]
  }
  Date <- str_extract(Title, "\\d{1,2}\\s\\w*\\s\\d{4}")
  TestData <- data.frame(ID,Title,Description,DCDL_IDs,Date)
  return(TestData)
} 

### Run it all

all.data <- data.frame(ID=character(),
                       Title=character(),
                       Description=character(),
                       DCDL_IDs=character(),
                       Date=character(),
                       stringsAsFactors=FALSE)

for (i in 1:number_calls) {
  all.data <- rbind(all.data, DataframeFromCall(api_calls[i]))
}
