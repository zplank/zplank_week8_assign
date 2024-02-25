###load data 
mydata <- data.frame(ames_housing_data)


###get memory and time before regression starts 
memory_before <- object.size(mydata)
time_before <- proc.time()

##run linear regression model 
model <- lm(SalePrice ~ ., data = mydata)

##print summary
summary(model)

###get memory usage and time after regression ends 
time_after <- proc.time()
memory_after <- object.size(step_model)

###calculate memory usage and processing time 
memory_change <- memory_after - memory_before
processing_time <- time_after - time_before

###print memory usage and processing times
cat("Memory Usage:", memory_change, "bytes\n")
cat("Processing Time:", processing_time[3], "seconds\n")
