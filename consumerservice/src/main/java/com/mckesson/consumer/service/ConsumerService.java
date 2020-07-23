package com.mckesson.consumer.service;

import org.json.JSONObject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
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


@Service
public class ConsumerService {

	private final static Logger log = LoggerFactory.getLogger(ConsumerService.class);

	public String callDataServiceAPI(String topicName , String message) {
		log.info("befor Calling the callDataServiceAPI>>" + topicName + "Message" + message);
		System.out.println("<<<<<< Status consumerDRGPayerData From Service >>>>" + topicName + "Message" + message);
		String statusMessage = ConsumerConstants.ERROR_MESSAGE;
		try {
			
			//https://mckessondataservice.app.ps.west.us.mckesson.com/greeting
			
			    final String uri = "http://dataservice:9897/mckesson/saveConsumerData";
			 
			    Message messageObj = new Message(topicName,message);
			 
			    RestTemplate restTemplate = new RestTemplate();
			    String result = restTemplate.postForObject( uri, messageObj, String.class);
			
				System.out.println("<<<<<< Status REST API >>>> From Service >>>>" + result );
			
			
			     log.info("After Calling the callDataServiceAPI>>" + statusMessage);

		} catch (Exception e) {
			log.error("Error Message" + e.getMessage());
		}

		return statusMessage;

	}

}
