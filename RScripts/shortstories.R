## what we did

lines <- c(0,2,5,12,16,17)
days <- c(0,1,2,3,4,5)
plot(y = lines,x = days, type = "o", xlab = "Days", ylab = "Knowledge")

## actual class

apicall <- "http://api.digitalnz.org/v3/records.json?api_key=LUq9-soDzWWhShuy3XhU&text=short+stories&and[collection][]=Papers+Past&fields=id,title,fulltext&per_page=100&page="

DataframeFromCall1 <- function(x) {
  data <- fromJSON(x)
  results <- data[[1]][2]
  ID <- unlist(results, recursive = FALSE)[[1]]
  TITLE <- unlist(results, recursive = FALSE)[[2]]
  FULLTEXT <- unlist(results, recursive = FALSE)[[3]]
  DATE <- str_extract(TITLE, "\\d{1,2}\\s\\w*\\s\\d{4}")
  TestData <- data.frame(ID,FULLTEXT,TITLE,DATE,stringsAsFactors=FALSE)
  return(TestData)
} 

all.data2 <- data.frame(ID=character(),
                        FULLTEXT=character(),
                        TITLE=character(),
                        DATE=character(),
                       stringsAsFactors=FALSE)

for (i in 1:30) {
  forcall <- paste0(apicall, as.character(i))
  all.data2 <- rbind(all.data2, DataframeFromCall1(forcall))
}

topic_model <- data.frame(Identifier = paste0("urn:cts:short:mixed01.mixed01:", as.character(all.data2$ID)), Text = all.data2$FULLTEXT, stringsAsFactors = F)
topic_model$Text <- gsub(x = topic_model$Text, pattern = '"', replacement = "'")
topic_model$Text <- trimws(topic_model$Text)
topic_model$Identifier <- trimws(topic_model$Identifier)

topic_model <- topic_model[-which(nchar(topic_model$Text) < 240),]


write.csv(topic_model, file = "topicmodeldata.csv", quote = T, row.names = F)
