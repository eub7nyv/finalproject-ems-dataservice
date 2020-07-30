package com.mckesson.consumer.service;

import org.json.JSONObject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import com.mckesson.consumer.entity.Message;
import com.mckesson.consumer.utils.ConsumerConstants;
import com.mckesson.consumer.utils.RestTemplateResponseErrorHandler;


@Service
public class ConsumerService {

	private final static Logger log = LoggerFactory.getLogger(ConsumerService.class);
	
	private RestTemplate restTeamplate;
	
	@Autowired
    public ConsumerService(RestTemplateBuilder restTemplateBuilder) {
		log.info("befor Calling the ConsumerService  REST API ERROR HANDLER START>>"  );
        RestTemplate restTemplate = restTemplateBuilder
          .errorHandler(new RestTemplateResponseErrorHandler())
          .build();
    	log.info("befor Calling the ConsumerService  REST API ERROR HANDLER END>>"  );
    }
	
	public String callDataServiceAPI(String topicName , String message) {
		log.info("befor Calling the callDataService REST API>>" + topicName + "Message" + message );
		log.info("befor Calling the callDataService REST API URI>>" + ConsumerConstants.REST_API_URL);
		System.out.println("<<<<<< Status consumerDRGPayerData From Service >>>>" + topicName + "Message" + message);
		String statusMessage = ConsumerConstants.ERROR_MESSAGE;
		try {
			    Message messageObj = new Message(topicName,message);
			    RestTemplate restTemplate = new RestTemplate();
			    String result = restTemplate.postForObject( ConsumerConstants.REST_API_URL, messageObj, String.class);
				System.out.println("<<<<<< Status REST API >>>> From Service >>>>" + result );
				statusMessage=ConsumerConstants.SUCCESS_MESSAGE;
			     log.info("After Calling the callDataServiceAPI>>" + statusMessage);
		} catch (Exception e) {
			log.error("Error Message" + e.getMessage());
		}

		return statusMessage;

	}
	
	public String CallServiceByTopic(String topic,String message) {
		
		return null;
	}

}
