# finalproject-ems-dataservice

## User Story/Goal
1. Build an enterprise message system to streamline data publishing and comsuming 
2. Build out an Centralize data service framework by leveraging message system to stream data and cloud and provide centralized service data 

## Issue Statement:  
### Enterprise message system to 
1. provide a centralize message structure for message queue
2. potentially replace current EMS system 

### Centralized Data Service to Resolve Below Issues
1.  Different system has to 
    1.1 build ETL load data 
    1.2 lookup data 
    1.3 Persist data in different system
2.  Replace Out date Repository  ( ie MARRSS etc.)
3.  Provide a cloud solution to centralized reference data service or other service

## Epic/User Story: 
1. Batch/Event Service to processing based on file ( or event)   
2. Producer Service	Receive message from Batch service 	(Provide a standard plumbing on streaming data into cloud or other data stored )	
3. Consumer Service	Receive message from Consumer (Provide a standard plumbing on streaming data into cloud or other data stored)
4. Hosted Data in Centralized DBStore (in Cloud)  Data Save Into Cloud DB	Post Consumer Received Data into Cloud DB	
5. Provide Centralized API to search & return a data set: Payer data search & Lookup
6. Kafka Dashboard/UI - for monitoring/end to end control/validation 		
7. Service Integration / CICD			

## Technology/Practice:  
Kafka Message Service, GCP (Postgres, NOsql, External Table?), Docker, GITHub, CICD, Golang and Java, SpringBoot
TDD, Functional Programming, DevOps, XP, Pair Programming , CPI, 

## Stretch Goal: Topic2 from event trigger to save as file

## Strategy Goal: 
Build a pattern/framework to provide centralized messaging system and provider master/reference data service. 
This message queue can be used for application event queue. 
And it can also be externed to other common data such as  NCPDP,  Medical Code , Provider search & lookup as centralized data service offering.

# doc folder 
Contain documents to the project.

# Project Planning
1. 8:00-8:30 scrum/resolution call
2. Epic/tasks are setup under this repo's project tab.