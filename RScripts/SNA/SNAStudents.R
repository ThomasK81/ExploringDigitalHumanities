library(tidyverse)
library(rtweet)

create_token(
  app = "",
  consumer_key = "",
  consumer_secret = "",
  access_token = "",
  access_secret = "")

my_followers <- get_followers("ThomasKoentges")
my_friends <- get_friends("ThomasKoentges")
my_data <- lookup_users("ThomasKoentges")
my_user_id <- my_data$user_id[1]

my_friends %>% 
  select(user_id) %>% 
  mutate(i_follow = TRUE) %>% 
  full_join(my_followers %>% mutate(follows_me = TRUE), by = 'user_id') %>% 
  mutate_all(funs(replace(., which(is.na(.)), FALSE))) %>% 
  select(2:3) %>% table

my_followers_data <- lookup_users(my_followers$user_id)

my_followers_data %>% arrange(desc(followers_count)) %>% select(screen_name, followers_count)

following <- data_frame(source_id = rep(my_user_id, n = length(my_followers$user_id)), user_id = my_followers$user_id)

for (i in 1:15) {
  followers <- get_followers(my_followers$user_id[i])
  if (length(followers) == 0) next
  followers$source_id <- rep(my_followers$user_id[i], n = length(followers$user_id))
  followers <- followers[,c(2,1)]
  following <- rbind(following, followers)
}

following_short <- following[duplicated(following$user_id),]
my_user_id %in% following_short$source_id
my_user_id %in% following$source_id
my_user_id %in% following_short$user_id

names(following_short) <- c("Source", "Target")

nodes <- unique(c(following_short$Source, following_short$Target))
short_data <- lookup_users(nodes)


nodesdf <- data.frame(Id = short_data$user_id, Label = short_data$screen_name, Link = paste0("https://twitter.com/", short_data$screen_name))
  
write_csv(following_short, path = "edges.csv")
write_csv(nodesdf, path = "nodes.csv")


