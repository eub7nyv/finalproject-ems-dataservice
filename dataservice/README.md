
## Mckesson Data Service


## BUILD
```
mvn clean install
docker-compose build
docker-compose up


##Implemented Topics 

1. drgpayer
2. drgplan
3. ncpdp

## Sample Messages



{
	"topicName" : "drgpayer",
	"message": "842|Ector County Independent School District|606"
}

{
	"topicName" : "drgplan",
	"message": "140|17629|Elderplan Advantage For Nursing Home Residents |Medicare|Medicare SN|Y|N|N"
}

{
	"topicName" : "ncpdp",
	"message": "140|Metropolitan Jewish Health System Elderplan|320AAAAA"
}

# REST API CALL

http://dataservice:9897/mckesson/saveConsumerDat
Parameters:
Message object contains
 1. TopicName
 2. Message



